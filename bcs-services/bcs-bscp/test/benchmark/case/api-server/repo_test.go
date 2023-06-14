/*
Tencent is pleased to support the open source community by making Basic Service Configuration Platform available.
Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except
in compliance with the License. You may obtain a copy of the License at
http://opensource.org/licenses/MIT
Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpload(t *testing.T) {
	filename := os.Getenv("filename")
	host := os.Getenv("host")
	if host == "" {
		host = "localhost:8080"
	}
	bizID := os.Getenv("biz_id")
	appID := os.Getenv("app_id")

	resp, err := upload(context.Background(), filename, host, bizID, appID, nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestDownload(t *testing.T) {

}
