// This file is part of GenId, a tool for generating UUIDs on the command line.
// Copyright (C) 2020 Jordan Ocokoljic.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"flag"
	"fmt"
	"github.com/JordanOcokoljic/go-uuid"
	"strings"
)

func main() {
	var plain bool
	flag.BoolVar(&plain, "plain", false, "the output will not contain dashes between segments")
	flag.Parse()

	id := uuid.Must().String()
	if plain {
		id = strings.Replace(id, "-", "", -1)
	}

	fmt.Print(id)
}
