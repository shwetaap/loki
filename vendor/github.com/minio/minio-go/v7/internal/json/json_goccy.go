//go:build !stdlibjson

/*
 * MinIO Go Library for Amazon S3 Compatible Cloud Storage
 * Copyright 2025 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package json

import "github.com/goccy/go-json"

// This file defines the JSON functions used internally and forwards them
// to goccy/go-json. Alternatively, the standard library can be used by setting
// the build tag stdlibjson. This can be useful for testing purposes or if
// goccy/go-json causes issues.
//
// This file does not contain all definitions from goccy/go-json; if needed, more
// can be added, but keep in mind that json_stdlib.go will also need to be
// updated.

var (
	// Unmarshal is a wrapper around goccy/go-json Unmarshal function.
	Unmarshal = json.Unmarshal
	// Marshal is a wrapper around goccy/go-json Marshal function.
	Marshal = json.Marshal
	// NewEncoder is a wrapper around goccy/go-json NewEncoder function.
	NewEncoder = json.NewEncoder
	// NewDecoder is a wrapper around goccy/go-json NewDecoder function.
	NewDecoder = json.NewDecoder
)

type (
	// Encoder is an alias for goccy/go-json Encoder.
	Encoder = json.Encoder
	// Decoder is an alias for goccy/go-json Decoder.
	Decoder = json.Decoder
)
