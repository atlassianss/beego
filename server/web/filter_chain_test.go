// Copyright 2020
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/astaxie/beego/server/web/context"
)

func TestControllerRegister_InsertFilterChain(t *testing.T) {

	InsertFilterChain("/*", func(next FilterFunc) FilterFunc {
		return func(ctx *context.Context) {
			ctx.Output.Header("filter", "filter-chain")
			next(ctx)
		}
	})

	ns := NewNamespace("/chain")

	ns.Get("/*", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello"))
	})

	r, _ := http.NewRequest("GET", "/chain/user", nil)
	w := httptest.NewRecorder()

	BeeApp.Handlers.ServeHTTP(w, r)

	assert.Equal(t, "filter-chain", w.Header().Get("filter"))
}
