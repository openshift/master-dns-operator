package oos

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Execution is a nested struct in oos response
type Execution struct {
	TemplateId        string        `json:"TemplateId" xml:"TemplateId"`
	ExecutedBy        string        `json:"ExecutedBy" xml:"ExecutedBy"`
	EndDate           string        `json:"EndDate" xml:"EndDate"`
	TemplateName      string        `json:"TemplateName" xml:"TemplateName"`
	RamRole           string        `json:"RamRole" xml:"RamRole"`
	IsParent          bool          `json:"IsParent" xml:"IsParent"`
	StatusMessage     string        `json:"StatusMessage" xml:"StatusMessage"`
	Mode              string        `json:"Mode" xml:"Mode"`
	StartDate         string        `json:"StartDate" xml:"StartDate"`
	SafetyCheck       string        `json:"SafetyCheck" xml:"SafetyCheck"`
	TemplateVersion   string        `json:"TemplateVersion" xml:"TemplateVersion"`
	CreateDate        string        `json:"CreateDate" xml:"CreateDate"`
	Parameters        string        `json:"Parameters" xml:"Parameters"`
	Counters          string        `json:"Counters" xml:"Counters"`
	UpdateDate        string        `json:"UpdateDate" xml:"UpdateDate"`
	ParentExecutionId string        `json:"ParentExecutionId" xml:"ParentExecutionId"`
	Outputs           string        `json:"Outputs" xml:"Outputs"`
	ExecutionId       string        `json:"ExecutionId" xml:"ExecutionId"`
	Status            string        `json:"Status" xml:"Status"`
	CurrentTasks      []CurrentTask `json:"CurrentTasks" xml:"CurrentTasks"`
}