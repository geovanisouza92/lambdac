// LambdaC API client
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/geovanisouza92/lambdac/types"
)

type client struct {
	e string
	c *http.Client
}

func New(endpoint string, hc *http.Client) API {
	return &client{
		e: endpoint + "/api/v1",
		c: hc,
	}
}

func (c *client) FunctionList() (functions types.Functions, err error) {
	err = c.send("GET", "/functions", http.StatusOK, nil, &functions)
	return
}

func (c *client) FunctionCreate(function types.Function) (out types.Function, err error) {
	err = c.send("POST", "/functions", http.StatusOK, function, &out)
	return
}

func (c *client) FunctionInfo(id string) (function types.Function, err error) {
	err = c.send("GET", escape("/functions/%s", id), http.StatusOK, nil, &function)
	return
}

func (c *client) FunctionConfig(id string, function types.Function) (err error) {
	err = c.send("PUT", escape("/functions/%s", id), http.StatusOK, function, nil)
	return
}

func (c *client) FunctionDestroy(id string, force bool) (err error) {
	data := struct {
		force bool `json:"force"`
	}{force}
	err = c.send("DELETE", escape("/functions/%s", id), http.StatusGone, data, nil)
	return
}

func (c *client) FunctionEnv(id string) (vars []string, err error) {
	err = c.send("GET", escape("/functions/%s/env", id), http.StatusOK, nil, &vars)
	return
}

func (c *client) FunctionEnvSet(id string, vars []string) (err error) {
	err = c.send("PUT", escape("/functions/%s/env", id), http.StatusAccepted, vars, nil)
	return
}

func (c *client) FunctionEnvUnset(id string, vars []string) (err error) {
	err = c.send("DELETE", escape("/functions/%s/env", id), http.StatusGone, vars, nil)
	return
}

func (c *client) FunctionPull(id string) (code string, err error) {
	// TODO Download code ...
	err = c.send("GET", escape("/functions/%s/code", id), http.StatusOK, nil, &code)
	return
}

func (c *client) FunctionPush(id, code string) (err error) {
	// TODO Upload code ...
	data := struct {
		code string `json:"code"`
	}{code}
	err = c.send("PUT", escape("/functions/%s/code", id), http.StatusAccepted, data, nil)
	return
}

func (c *client) FunctionPs(id string) (out []string, err error) {
	err = c.send("GET", escape("/functions/%s/ps", id), http.StatusOK, nil, &out)
	return
}

func (c *client) FunctionLogs(id string) (out []string, err error) {
	err = c.send("GET", escape("/functions/%s/logs", id), http.StatusOK, nil, &out)
	return
}

func (c *client) FunctionStats(id string) (out []string, err error) {
	err = c.send("GET", escape("/functions/%s/stats", id), http.StatusOK, nil, &out)
	return
}

func (c *client) FunctionInvoke(id string) (err error) {
	// TODO reqdata
	// TODO output?
	err = c.send("POST", escape("/functions/%s/invoke", id), http.StatusAccepted, nil, nil)
	return
}

func (c *client) RuntimeList() (runtimes types.Runtimes, err error) {
	err = c.send("GET", "/runtimes", http.StatusOK, nil, &runtimes)
	return
}

func (c *client) RuntimeCreate(runtime types.Runtime) (out types.Runtime, err error) {
	err = c.send("POST", "/runtimes", http.StatusCreated, runtime, &out)
	return
}

func (c *client) RuntimeInfo(id string) (runtime types.Runtime, err error) {
	err = c.send("GET", escape("/runtimes/%s", id), http.StatusOK, nil, &runtime)
	return
}

func (c *client) RuntimeDestroy(id string, force bool) (err error) {
	data := struct {
		force bool `json:"force"`
	}{force}
	err = c.send("DELETE", escape("/runtimes/%s", id), http.StatusGone, data, nil)
	return
}

func (c *client) send(method, path string, expectedStatus int, reqdata, respdata interface{}) error {
	var br io.ReadWriter

	// Serialize request body
	if reqdata != nil {
		br = &bytes.Buffer{}
		e := json.NewEncoder(br)

		if err := e.Encode(reqdata); err != nil {
			return err
		}
	}

	// Create request
	req, err := http.NewRequest(method, c.e+path, br)
	if err != nil {
		return err
	}

	// Send request
	res, err := c.c.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return err
	}

	// Process unexpected error
	if res.StatusCode != expectedStatus {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("api response with status: %d\n%s", res.StatusCode, string(b))
	}

	// Ignored response data
	if respdata == nil {
		return nil
	}

	// Parse response
	d := json.NewDecoder(res.Body)
	if err = d.Decode(respdata); err != nil {
		return err
	}

	return nil
}

func escape(tpl, s string) string {
	return fmt.Sprintf(tpl, url.QueryEscape(s))
}
