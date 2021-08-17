/*
 * This file is part of the libvirt-go-module project
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 * Copyright (C) 2022 Red Hat, Inc.
 *
 */

package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

type API struct {
	XMLName   xml.Name      `xml:"api"`
	Name      string        `xml:"name,attr"`
	Files     []APIFile     `xml:"files>file"`
	Macros    []APIMacro    `xml:"symbols>macro"`
	Typedefs  []APITypedef  `xml:"symbols>typedef"`
	Enums     []APIEnum     `xml:"symbols>enum"`
	Structs   []APIStruct   `xml:"symbols>struct"`
	Functions []APIFunction `xml:"symbols>function"`
	Functypes []APIFunctype `xml:"symbols>functype"`
	Variables []APIVariable `xml:"symbols>variable"`
}

type APIFile struct {
	Name        string `xml:"name,attr"`
	Summary     string `xml:"summary"`
	Description string `xml:"description"`

	Exports []APIExport `xml:"exports"`
}

type APIExport struct {
	Symbol string `xml:"symbol,attr"`
	Type   string `xml:"type,attr"`
}

type APIMacro struct {
	Name    string        `xml:"name,attr"`
	File    string        `xml:"file,attr"`
	Params  string        `xml:"params,attr"`
	String  string        `xml:"string,attr"`
	Raw     string        `xml:"raw,attr"`
	Version string        `xml:"version,attr"`
	Info    string        `xml:"info"`
	Args    []APIMacroArg `xml:"arg"`
}

type APIMacroArg struct {
	Name string `xml:"name,attr"`
	Info string `xml:"info,attr"`
}

type APIEnum struct {
	Name          string `xml:"name,attr"`
	File          string `xml:"file,attr"`
	Info          string `xml:"info,attr"`
	Value         string `xml:"value,attr"`
	ValueHex      string `xml:"value_hex,attr"`
	ValueBitshift string `xml:"value_bitshift,attr"`
	Type          string `xml:"type,attr"`
	Version       string `xml:"version,attr"`
	ValueRaw      int64
}

type APIStruct struct {
	Name    string `xml:"name,attr"`
	File    string `xml:"file,attr"`
	Type    string `xml:"type,attr"`
	Version string `xml:"version,attr"`

	Fields   []APIStructField `xml:"field"`
	Exported bool
}

type APIStructField struct {
	Name  string    `xml:"name,attr"`
	Type  string    `xml:"type,attr"`
	Info  string    `xml:"info,attr"`
	Union *APIUnion `xml:"union"`
}

type APIUnion struct {
	Fields []APIStructField `xml:"field"`
}

type APITypedef struct {
	Name    string `xml:"name,attr"`
	File    string `xml:"file,attr"`
	Type    string `xml:"type,attr"`
	Info    string `xml:"info"`
	Version string `xml:"version,attr"`
}

type APIFunction struct {
	Name    string            `xml:"name,attr"`
	File    string            `xml:"file,attr"`
	Module  string            `xml:"module,attr"`
	Version string            `xml:"version,attr"`
	Info    string            `xml:"info"`
	Return  APIFunctionReturn `xml:"return"`
	Args    []APIFunctionArg  `xml:"arg"`
}

type APIFunctype struct {
	Name    string            `xml:"name,attr"`
	File    string            `xml:"file,attr"`
	Module  string            `xml:"module,attr"`
	Info    string            `xml:"info"`
	Return  APIFunctionReturn `xml:"return"`
	Args    []APIFunctionArg  `xml:"arg"`
	Version string            `xml:"version,attr"`
}

type APIFunctionArg struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
	Info string `xml:"info,attr"`
}

type APIFunctionReturn struct {
	Type string `xml:"type,attr"`
	Info string `xml:"info,attr"`
}

type APIVariable struct {
	Name    string `xml:"name,attr"`
	File    string `xml:"file,attr"`
	Type    string `xml:"type,attr"`
	Version string `xml:"version,attr"`
}

// Type which implements the interface for ordering the array of APIEnum
type EnumByType []APIEnum

func (a EnumByType) Len() int      { return len(a) }
func (a EnumByType) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a EnumByType) Less(i, j int) bool {
	// First sort level is by type
	if a[i].Type != a[j].Type {
		return a[i].Type < a[j].Type
	}
	// Second sort level is hex value
	if a[i].ValueRaw != a[j].ValueRaw {
		return a[i].ValueRaw < a[j].ValueRaw
	}
	// Last is by name
	return a[i].Name < a[j].Name
}

// Type which implements the interface for ordering the array of APIStruct
type StructByType []APIStruct

func (a StructByType) Len() int      { return len(a) }
func (a StructByType) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a StructByType) Less(i, j int) bool {
	// First sort level is by type
	if a[i].Type != a[j].Type {
		return a[i].Type < a[j].Type
	}
	// Last is by name
	return a[i].Name < a[j].Name
}

// Calculate raw values and sort the Enums first by Type and then by its raw value.
func prepareEnums(a, coreAPI *API) {
	enumValues := make(map[string]*APIEnum)
	/* Some of the secondary API module enums are defined
	 * using constants from the core API */
	if coreAPI != nil {
		for idx, _ := range coreAPI.Enums {
			enum := &coreAPI.Enums[idx]
			enumValues[enum.Name] = enum
		}
	}
	for idx, _ := range a.Enums {
		enum := &a.Enums[idx]
		val, err := strconv.ParseInt(enum.Value, 10, 64)
		if err == nil {
			enum.ValueRaw = val
			enumValues[enum.Name] = enum
		}
	}
	for idx, _ := range a.Enums {
		enum := &a.Enums[idx]
		_, ok := enumValues[enum.Name]
		if !ok {
			enumref, ok2 := enumValues[enum.Value]
			if !ok2 {
				log.Fatalf("Resolving %s -> %s still empty", enum.Name, enum.Value)
			} else {
				enum.ValueRaw = enumref.ValueRaw
			}
		}
	}

	sort.Sort(EnumByType(a.Enums))
}

// Removes duplicated by type
func prepareStructs(a *API) {
	sort.Sort(StructByType(a.Structs))

	exported := make(map[string]bool)
	for _, f := range a.Files {
		for _, e := range f.Exports {
			if e.Type == "struct" {
				exported[e.Symbol] = true
			}
		}
	}

	for idx, _ := range a.Structs {
		str := &a.Structs[idx]

		if !strings.HasPrefix(str.Type, "struct ") {
			log.Fatalf("Struct name %s", str.Type)
		}
		typ := str.Type[7:]

		_, ok := exported[typ]
		if ok {
			str.Exported = true
		}
	}
}

func (a *API) prepare(coreAPI *API) {
	if coreAPI != nil {
		prepareEnums(coreAPI, nil)
	}
	prepareEnums(a, coreAPI)
	prepareStructs(a)
}

func versionToNumber(version string) (int, int, int) {
	strv := strings.Split(version, ".")
	if len(strv) != 3 {
		panic(fmt.Sprintf("bad version: %v (%s)", strv, version))
	}
	major, err1 := strconv.Atoi(strv[0])
	minor, err2 := strconv.Atoi(strv[1])
	micro, err3 := strconv.Atoi(strv[2])
	if err1 != nil || err2 != nil || err3 != nil {
		panic(fmt.Sprintf("bad version: %v", strv))
	}
	return major, minor, micro
}

func getVersionMajor(version string) int {
	major, _, _ := versionToNumber(version)
	return major
}

func getVersionMinor(version string) int {
	_, minor, _ := versionToNumber(version)
	return minor
}

func getVersionMicro(version string) int {
	_, _, micro := versionToNumber(version)
	return micro
}

// Break line at all ',' and indent based on @prefix's length.
func indentArgs(input, prefix string) string {
	indent := fmt.Sprintf(",\n%s", strings.Repeat(" ", len(prefix)))
	return strings.ReplaceAll(input, ",", indent)
}

func getIncludeName(module string) string {
	return strings.Replace(module, "-", "_", -1)
}

func getEnumString(enum APIEnum) string {
	if enum.ValueBitshift != "" {
		mod := ""
		val, err := strconv.ParseInt(enum.ValueBitshift, 10, 64)
		if err == nil && val >= 31 {
			mod = "U"
		}
		return "(1" + mod + " << " + enum.ValueBitshift + ")"
	} else {
		return enum.Value
	}
}

func getAPIPathPkgConfig(varname, modname string) (string, error) {
	cmd := exec.Command("pkg-config", "--variable="+varname, modname)

	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("Failed to run pkg-config --variable=%s %s",
			varname, modname)
	}
	apixml := strings.TrimSpace(string(cmdOutput.Bytes()))
	if apixml == "" {
		return "", fmt.Errorf("Missing API XML from pkg-config --variable=%s %s",
			varname, modname)
	}
	return apixml, nil
}

func runTemplate(pathTemplate, pathOutput string, fnMap template.FuncMap, api *API) {
	templateFile, err := os.ReadFile(pathTemplate)
	if err != nil {
		log.Fatalf("Input: %s", err)
	}

	tmpl, err := template.New(pathOutput).Funcs(fnMap).Parse(string(templateFile))
	if err != nil {
		log.Fatalf("Parsing: %s", err)
	}

	output, err := os.OpenFile(pathOutput, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Output: %s", err)
	}

	err = tmpl.Execute(output, api)
	if err != nil {
		log.Fatalf("Execution: %s", err)
	}
}

func generate(apixml string, coreAPI *API) (*API, error) {
	var api API
	xmldata, err := ioutil.ReadFile(apixml)
	if err != nil {
		return nil, fmt.Errorf("Cannot read %s: %s", apixml, err)
	}

	err = xml.Unmarshal(xmldata, &api)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse %s: %s", apixml, err)
	}

	api.prepare(coreAPI)

	// The suffix of the output file name, maps
	// to similarly named template input file
	outputFiles := []string{
		"generated.h",
		"generated_macros.h",
		"generated_enums.h",
		"generated_typedefs.h",
		"generated_callbacks.h",
		"generated_structs.h",
		"generated_variables.h",
	}

	fnMap := template.FuncMap{
		"contains":        strings.Contains,
		"indent":          indentArgs,
		"getVersionMajor": getVersionMajor,
		"getVersionMinor": getVersionMinor,
		"getVersionMicro": getVersionMicro,
		"getIncludeName":  getIncludeName,
		"getEnumString":   getEnumString,
	}

	for _, outputSuffix := range outputFiles {
		output := strings.Replace(api.Name, "-", "_", -1) + "_" + outputSuffix
		input := path.Join("gen", "api_"+outputSuffix+".tmpl")
		runTemplate(input, output, fnMap, &api)
	}

	return &api, nil
}

type APIModule struct {
	APIVar        string
	PkgConfigFile string
}

func main() {
	apimodules := []APIModule{
		APIModule{"libvirt_api", "libvirt"},
		APIModule{"libvirt_lxc_api", "libvirt-lxc"},
		APIModule{"libvirt_qemu_api", "libvirt-qemu"},
	}

	var coreAPI *API
	for _, apimodule := range apimodules {
		apixml, err := getAPIPathPkgConfig(apimodule.APIVar, apimodule.PkgConfigFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}

		api, err := generate(apixml, coreAPI)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}

		if coreAPI == nil {
			coreAPI = api
		}
	}
	os.Exit(0)
}
