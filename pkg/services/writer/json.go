/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package writer

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/gsdenys/pdgen/pkg/models"
)

type JSON struct {
	Out io.Writer
}

func (p *JSON) SetWriter(path string) error {
	file, err := CreateFile(path)

	if err != nil {
		return err
	}

	p.Out = file
	return nil
}

func (p *JSON) Init(desc models.Describe) {
	// Do nothing because have nothing to initialise
}

func (p *JSON) Title(title string) {
	//Do nothing because the unique action of this writer is Done
}

func (p *JSON) Subtitle(subtitle string) {
	//Do nothing because the unique action of this writer is Done
}

func (p *JSON) SubSubtitle(subSubtitle string) {
	//Do nothing because the unique action of this writer is Done
}

func (p *JSON) LineBreak() {
	//Do nothing because the unique action of this writer is Done
}

func (p *JSON) Body(desc string) {
	//Do nothing because the unique action of this writer is Done
}

func (p *JSON) Columns(columns []models.Columns) {
	//Do nothing because the unique action of this writer is Done
}

func (p *JSON) Table(t models.Table) {
	//Do nothing because the unique action of this writer is Done
}

func (p *JSON) Done(desc models.Describe) {
	b, _ := json.MarshalIndent(desc, "", "    ")

	fmt.Fprintf(p.Out, "%s", string(b))

	_ = p.Out.(*os.File).Close()
}
