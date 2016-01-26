/*
This file is part of GO gallery.

GO gallery is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GO gallery is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with GO gallery.  If not, see <http://www.gnu.org/licenses/>.
*/

package file

import (
	"io"
	"os"
	"path"
	"strings"
	"time"
)

// Copies file source to destination dest.
func CopyFile(source string, dest string) (err error) {
	sf, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sf.Close()
	df, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer df.Close()

	_, err = io.Copy(df, sf)
	if err != nil {
		return err
	}

	err = df.Close()
	if err != nil {
		return err
	}
	si, err := os.Stat(source)
	if err != nil {
		return err
	}
	err = os.Chmod(dest, si.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(dest, time.Now(), si.ModTime())
	return
}

func PathJoin(ps ...interface{}) string {
	pps := make([]string, len(ps))
	for i, p := range ps {
		pps[i] = strings.Replace(p.(string), "..", ".", -1)
	}
	return path.Join(pps...)
}
