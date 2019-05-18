package main

import (
	"bytes"
	"image"
	"image/png"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
)

func mkimage(Roles map[string]int, Cards map[string]string) []*bytes.Reader {
	ToSend := []*bytes.Reader{}
	Total := 0
	for _, _ = range Roles {
		Total++
	}

	dc := &gg.Context{}

	if Total >= 5 {
		dc = gg.NewContext(320, 64*5)
	} else {
		dc = gg.NewContext(320, 64*Total)
	}
	dc.SetRGB255(255, 255, 255)

	i := 0

	for key, amount := range Roles {
		i++
		Total--

		AchDat, err := box.FindString("compo/base.png")
		check(err)
		AchImg, _, err := image.Decode(strings.NewReader(AchDat))
		check(err)
		RoleDat, err := box.FindString("40x40/" + Cards[key])
		check(err)
		RoleImg, _, err := image.Decode(strings.NewReader(RoleDat))
		check(err)

		dc.DrawImage(AchImg, 0, 64*(i-1))
		dc.DrawImage(RoleImg, 15, 10+64*(i-1))

		dc.DrawString("x"+strconv.Itoa(amount)+" "+strings.Replace(strings.Replace(key, "é", "", 3), "è", "", 3), 70, float64(35+64*(i-1)))

		if i%5 == 0 {
			i = 0
			dc.Clip()
			buff := new(bytes.Buffer)
			err := png.Encode(buff, dc.Image())
			check(err)
			reader := bytes.NewReader(buff.Bytes())
			ToSend = append(ToSend, reader)
			if Total >= 5 {
				dc = gg.NewContext(320, 64*5)
			} else {
				dc = gg.NewContext(320, 64*Total)
			}
			dc.SetRGB255(255, 255, 255)
		}
	}

	dc.Clip()
	buff := new(bytes.Buffer)
	err := png.Encode(buff, dc.Image())
	check(err)
	reader := bytes.NewReader(buff.Bytes())
	ToSend = append(ToSend, reader)

	return ToSend

}
