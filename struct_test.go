// © 2014 Steve McCoy under the MIT license. See LICENSE for details.

package tmx

import (
	"strings"
	"testing"

	"github.com/eaburns/eq"
)

func TestDecode(t *testing.T) {
	var maps []*Map

	for i, s := range []string{ testXml, testCsv, testBase64, testGzip, testZlib } {
		m, err := Decode(strings.NewReader(s))
		if err != nil {
			t.Fatalf("unexpected decode error for %d: %v", i, err)
		}
		maps = append(maps, m)
	}

	for i := 1; i < len(maps); i++ {
		if !eq.Deep(maps[i-1], maps[i]) {
			t.Fatalf("unequal %d vs. %d:\n%v\n------\n%v", i-1, i, maps[i-1], maps[i])
		}
	}
}

func TestBadDecode(t *testing.T) {
	for i, s := range []string{ testBadXml, testBadCsv, testBadBase64 } {
		_, err := Decode(strings.NewReader(s))
		if err == nil {
			t.Fatalf("expected decode error for %d, got: %v", i, err)
		}
	}
}

var testXml = `
<?xml version="1.0" encoding="UTF-8"?>
<map version="1.0" orientation="orthogonal" width="10" height="10" tilewidth="16" tileheight="16">
 <tileset firstgid="1" name="land" tilewidth="16" tileheight="16">
  <image source="tiles.png" width="48" height="16"/>
 </tileset>
 <layer name="Foreground" width="10" height="10">
  <data>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="2"/>
   <tile gid="2"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="2"/>
   <tile gid="2"/>
   <tile gid="2"/>
   <tile gid="2"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="2"/>
   <tile gid="2"/>
   <tile gid="2"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="2"/>
   <tile gid="2"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
  </data>
 </layer>
 <objectgroup name="Mountains" width="10" height="10">
  <object type="mountain" x="48" y="48" width="16" height="16"/>
  <object type="mountain" x="64" y="32" width="16" height="16"/>
  <object type="mountain" x="80" y="32" width="16" height="16"/>
  <object type="mountain" x="64" y="48" width="16" height="16"/>
  <object type="mountain" x="96" y="48" width="16" height="16"/>
  <object type="mountain" x="80" y="48" width="16" height="16"/>
  <object type="mountain" x="64" y="64" width="16" height="16"/>
  <object type="mountain" x="80" y="64" width="16" height="16"/>
  <object type="mountain" x="96" y="64" width="16" height="16"/>
  <object type="mountain" x="80" y="80" width="16" height="16"/>
  <object type="mountain" x="96" y="80" width="16" height="16"/>
 </objectgroup>
</map>
`

var testBadXml = `
<?xml version="1.0" encoding="UTF-8"?>
<map version="1.0" orientation="orthogonal" width="10" height="10" tilewidth="16" tileheight="16">
 <tileset firstgid="1" name="land" tilewidth="16" tileheight="16">
  <image source="tiles.png" width="48" height="16"/>
 </tileset>
 <layer name="Foreground" width="10" height="10">
  <data>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <derp gid="3">
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="2"/>
   <tile gid="2"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="2"/>
   <tile gid="2"/>
   <tile gid="2"/>
   <tile gid="2"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="2"/>
   <tile gid="2"/>
   <tile gid="2"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="2"/>
   <tile gid="2"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
   <tile gid="3"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="1"/>
   <tile gid="3"/>
  </data>
 </layer>
 <objectgroup name="Mountains" width="10" height="10">
  <object type="mountain" x="48" y="48" width="16" height="16"/>
  <object type="mountain" x="64" y="32" width="16" height="16"/>
  <object type="mountain" x="80" y="32" width="16" height="16"/>
  <object type="mountain" x="64" y="48" width="16" height="16"/>
  <object type="mountain" x="96" y="48" width="16" height="16"/>
  <object type="mountain" x="80" y="48" width="16" height="16"/>
  <object type="mountain" x="64" y="64" width="16" height="16"/>
  <object type="mountain" x="80" y="64" width="16" height="16"/>
  <object type="mountain" x="96" y="64" width="16" height="16"/>
  <object type="mountain" x="80" y="80" width="16" height="16"/>
  <object type="mountain" x="96" y="80" width="16" height="16"/>
 </objectgroup>
</map>
`

var testCsv = `
<?xml version="1.0" encoding="UTF-8"?>
<map version="1.0" orientation="orthogonal" width="10" height="10" tilewidth="16" tileheight="16">
 <tileset firstgid="1" name="land" tilewidth="16" tileheight="16">
  <image source="tiles.png" width="48" height="16"/>
 </tileset>
 <layer name="Foreground" width="10" height="10">
  <data encoding="csv">
3,3,1,1,1,1,1,1,1,3,
3,1,1,3,3,3,3,1,1,1,
1,1,3,3,2,2,3,3,1,1,
1,1,3,2,2,2,2,3,1,1,
1,1,3,3,2,2,2,3,1,1,
1,1,3,3,3,2,2,3,1,1,
1,1,3,3,3,3,3,3,3,1,
3,1,1,3,3,3,3,3,3,1,
3,1,1,1,1,1,3,3,1,1,
3,3,1,1,1,1,1,1,1,3
</data>
 </layer>
 <objectgroup name="Mountains" width="10" height="10">
  <object type="mountain" x="48" y="48" width="16" height="16"/>
  <object type="mountain" x="64" y="32" width="16" height="16"/>
  <object type="mountain" x="80" y="32" width="16" height="16"/>
  <object type="mountain" x="64" y="48" width="16" height="16"/>
  <object type="mountain" x="96" y="48" width="16" height="16"/>
  <object type="mountain" x="80" y="48" width="16" height="16"/>
  <object type="mountain" x="64" y="64" width="16" height="16"/>
  <object type="mountain" x="80" y="64" width="16" height="16"/>
  <object type="mountain" x="96" y="64" width="16" height="16"/>
  <object type="mountain" x="80" y="80" width="16" height="16"/>
  <object type="mountain" x="96" y="80" width="16" height="16"/>
 </objectgroup>
</map>
`

var testBadCsv = `
<?xml version="1.0" encoding="UTF-8"?>
<map version="1.0" orientation="orthogonal" width="10" height="10" tilewidth="16" tileheight="16">
 <tileset firstgid="1" name="land" tilewidth="16" tileheight="16">
  <image source="tiles.png" width="48" height="16"/>
 </tileset>
 <layer name="Foreground" width="10" height="10">
  <data encoding="csv">
3,3,1,1,1,1,1,1,1,3,
3,1,1,3,3,3,3,1,1,1,
1,1,3,3,2,2,3,3,1,1,
1,1,3,2,2,2,2,3,1,1,
1,1,3,3,2,x,2,3,1,1,
1,1,3,3,3,2,2,3,1,1,
1,1,3,3,3,3,3,3,3,1,
3,1,1,3,3,3,3,3,3,1,
3,1,1,1,1,1,3,3,1,1,
3,3,1,1,1,1,1,1,1,3
</data>
 </layer>
 <objectgroup name="Mountains" width="10" height="10">
  <object type="mountain" x="48" y="48" width="16" height="16"/>
  <object type="mountain" x="64" y="32" width="16" height="16"/>
  <object type="mountain" x="80" y="32" width="16" height="16"/>
  <object type="mountain" x="64" y="48" width="16" height="16"/>
  <object type="mountain" x="96" y="48" width="16" height="16"/>
  <object type="mountain" x="80" y="48" width="16" height="16"/>
  <object type="mountain" x="64" y="64" width="16" height="16"/>
  <object type="mountain" x="80" y="64" width="16" height="16"/>
  <object type="mountain" x="96" y="64" width="16" height="16"/>
  <object type="mountain" x="80" y="80" width="16" height="16"/>
  <object type="mountain" x="96" y="80" width="16" height="16"/>
 </objectgroup>
</map>
`

var testBase64 = `
<?xml version="1.0" encoding="UTF-8"?>
<map version="1.0" orientation="orthogonal" width="10" height="10" tilewidth="16" tileheight="16">
 <tileset firstgid="1" name="land" tilewidth="16" tileheight="16">
  <image source="tiles.png" width="48" height="16"/>
 </tileset>
 <layer name="Foreground" width="10" height="10">
  <data encoding="base64">
   AwAAAAMAAAABAAAAAQAAAAEAAAABAAAAAQAAAAEAAAABAAAAAwAAAAMAAAABAAAAAQAAAAMAAAADAAAAAwAAAAMAAAABAAAAAQAAAAEAAAABAAAAAQAAAAMAAAADAAAAAgAAAAIAAAADAAAAAwAAAAEAAAABAAAAAQAAAAEAAAADAAAAAgAAAAIAAAACAAAAAgAAAAMAAAABAAAAAQAAAAEAAAABAAAAAwAAAAMAAAACAAAAAgAAAAIAAAADAAAAAQAAAAEAAAABAAAAAQAAAAMAAAADAAAAAwAAAAIAAAACAAAAAwAAAAEAAAABAAAAAQAAAAEAAAADAAAAAwAAAAMAAAADAAAAAwAAAAMAAAADAAAAAQAAAAMAAAABAAAAAQAAAAMAAAADAAAAAwAAAAMAAAADAAAAAwAAAAEAAAADAAAAAQAAAAEAAAABAAAAAQAAAAEAAAADAAAAAwAAAAEAAAABAAAAAwAAAAMAAAABAAAAAQAAAAEAAAABAAAAAQAAAAEAAAABAAAAAwAAAA==
  </data>
 </layer>
 <objectgroup name="Mountains" width="10" height="10">
  <object type="mountain" x="48" y="48" width="16" height="16"/>
  <object type="mountain" x="64" y="32" width="16" height="16"/>
  <object type="mountain" x="80" y="32" width="16" height="16"/>
  <object type="mountain" x="64" y="48" width="16" height="16"/>
  <object type="mountain" x="96" y="48" width="16" height="16"/>
  <object type="mountain" x="80" y="48" width="16" height="16"/>
  <object type="mountain" x="64" y="64" width="16" height="16"/>
  <object type="mountain" x="80" y="64" width="16" height="16"/>
  <object type="mountain" x="96" y="64" width="16" height="16"/>
  <object type="mountain" x="80" y="80" width="16" height="16"/>
  <object type="mountain" x="96" y="80" width="16" height="16"/>
 </objectgroup>
</map>
`

var testBadBase64 = `
<?xml version="1.0" encoding="UTF-8"?>
<map version="1.0" orientation="orthogonal" width="10" height="10" tilewidth="16" tileheight="16">
 <tileset firstgid="1" name="land" tilewidth="16" tileheight="16">
  <image source="tiles.png" width="48" height="16"/>
 </tileset>
 <layer name="Foreground" width="10" height="10">
  <data encoding="base64">
   AwAAAAMAAAABAAAAAQAAAAEAAAABAAAAAQAAAAEAAAABAAAAAwAAAAMAAAABAAAAAQAAAAMAAAADAAAAAwAAAAMAAAABAAAAAQAAAAEAAAABAAAAAQAAAAMAAAADAAAAAgAAAAIAAAADAAAAAwAAAAEAAAABAAAAAQAAAAEAAAADAAAAAgAAAAIAAAAC~AAAAgAAAAMAAAABAAAAAQAAAAEAAAABAAAAAwAAAAMAAAACAAAAAgAAAAIAAAADAAAAAQAAAAEAAAABAAAAAQAAAAMAAAADAAAAAwAAAAIAAAACAAAAAwAAAAEAAAABAAAAAQAAAAEAAAADAAAAAwAAAAMAAAADAAAAAwAAAAMAAAADAAAAAQAAAAMAAAABAAAAAQAAAAMAAAADAAAAAwAAAAMAAAADAAAAAwAAAAEAAAADAAAAAQAAAAEAAAABAAAAAQAAAAEAAAADAAAAAwAAAAEAAAABAAAAAwAAAAMAAAABAAAAAQAAAAEAAAABAAAAAQAAAAEAAAABAAAAAwAAAA==
  </data>
 </layer>
 <objectgroup name="Mountains" width="10" height="10">
  <object type="mountain" x="48" y="48" width="16" height="16"/>
  <object type="mountain" x="64" y="32" width="16" height="16"/>
  <object type="mountain" x="80" y="32" width="16" height="16"/>
  <object type="mountain" x="64" y="48" width="16" height="16"/>
  <object type="mountain" x="96" y="48" width="16" height="16"/>
  <object type="mountain" x="80" y="48" width="16" height="16"/>
  <object type="mountain" x="64" y="64" width="16" height="16"/>
  <object type="mountain" x="80" y="64" width="16" height="16"/>
  <object type="mountain" x="96" y="64" width="16" height="16"/>
  <object type="mountain" x="80" y="80" width="16" height="16"/>
  <object type="mountain" x="96" y="80" width="16" height="16"/>
 </objectgroup>
</map>
`

var testGzip = `
<?xml version="1.0" encoding="UTF-8"?>
<map version="1.0" orientation="orthogonal" width="10" height="10" tilewidth="16" tileheight="16">
 <tileset firstgid="1" name="land" tilewidth="16" tileheight="16">
  <image source="tiles.png" width="48" height="16"/>
 </tileset>
 <layer name="Foreground" width="10" height="10">
  <data encoding="base64" compression="gzip">
   H4sIAAAAAAAAA2NmYGBgBmJGPJgZTQ0zGsanhwmKcalFVoOslpB5hNSh241PHTb/4PMvLnXEhBmhcAYABg1fRZABAAA=
  </data>
 </layer>
 <objectgroup name="Mountains" width="10" height="10">
  <object type="mountain" x="48" y="48" width="16" height="16"/>
  <object type="mountain" x="64" y="32" width="16" height="16"/>
  <object type="mountain" x="80" y="32" width="16" height="16"/>
  <object type="mountain" x="64" y="48" width="16" height="16"/>
  <object type="mountain" x="96" y="48" width="16" height="16"/>
  <object type="mountain" x="80" y="48" width="16" height="16"/>
  <object type="mountain" x="64" y="64" width="16" height="16"/>
  <object type="mountain" x="80" y="64" width="16" height="16"/>
  <object type="mountain" x="96" y="64" width="16" height="16"/>
  <object type="mountain" x="80" y="80" width="16" height="16"/>
  <object type="mountain" x="96" y="80" width="16" height="16"/>
 </objectgroup>
</map>
`

var testZlib = `
<?xml version="1.0" encoding="UTF-8"?>
<map version="1.0" orientation="orthogonal" width="10" height="10" tilewidth="16" tileheight="16">
 <tileset firstgid="1" name="land" tilewidth="16" tileheight="16">
  <image source="tiles.png" width="48" height="16"/>
 </tileset>
 <layer name="Foreground" width="10" height="10">
  <data encoding="base64" compression="zlib">
   eJxjZmBgYAZiRjyYGU0NMxrGp4cJinGpRVaDrJaQeYTUoduNTx02/+DzLy51xIQZoXAGAJlEAMI=
  </data>
 </layer>
 <objectgroup name="Mountains" width="10" height="10">
  <object type="mountain" x="48" y="48" width="16" height="16"/>
  <object type="mountain" x="64" y="32" width="16" height="16"/>
  <object type="mountain" x="80" y="32" width="16" height="16"/>
  <object type="mountain" x="64" y="48" width="16" height="16"/>
  <object type="mountain" x="96" y="48" width="16" height="16"/>
  <object type="mountain" x="80" y="48" width="16" height="16"/>
  <object type="mountain" x="64" y="64" width="16" height="16"/>
  <object type="mountain" x="80" y="64" width="16" height="16"/>
  <object type="mountain" x="96" y="64" width="16" height="16"/>
  <object type="mountain" x="80" y="80" width="16" height="16"/>
  <object type="mountain" x="96" y="80" width="16" height="16"/>
 </objectgroup>
</map>
`
