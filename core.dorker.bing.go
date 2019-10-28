/*
	Copyright © 2019 The Goca.io team

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package main

import (
	"fmt"
	"net/url"

	"github.com/gocaio/goca/dork"
)

// core.dorker.bing.go implements the bing engine specific dorker structure and
// functionality.

const bingEngineURL = "https://www.bing.com/search?q=%s&first=%d0"

// Bing execute dorks against Bing engine
func (d *Dorker) Bing(plugDork *dork.Dork) (results []string) {
	re := `<h2><a href="(.*?)" h=`
	escapedDork := url.QueryEscape(plugDork.Query)

	for i := 0; i < d.pages; i++ {
		u := fmt.Sprintf(bingEngineURL, escapedDork, i)

		wbuf, err := d.get(u)
		if err != nil {
			logError(fmt.Sprintf("[!] (%s): %s", u, err.Error()))
			continue
		}

		parsedData := parser(wbuf, re)
		for j := range parsedData {
			urlUnescaped, err := url.QueryUnescape(parsedData[j][1])
			if err != nil {
				logError(fmt.Sprintf("[!] %s", err.Error()))
				continue
			}

			results = append(results, urlUnescaped)
		}
	}

	return
}