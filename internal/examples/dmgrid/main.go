package main

import (
	"fmt"
	"os"
	"time"

	"github.com/prasangmisra/maroto/pkg/consts"
	"github.com/prasangmisra/maroto/pkg/pdf"
	"github.com/prasangmisra/maroto/pkg/props"
)

func main() {
	begin := time.Now()
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)
	m.SetBorder(true)

	m.Row(40, func() {
		m.Col(2, func() {
			m.DataMatrixCode("https://github.com/prasangmisra/maroto", props.Rect{
				Percent: 50,
			})
		})
		m.Col(4, func() {
			m.DataMatrixCode("https://github.com/prasangmisra/maroto", props.Rect{
				Percent: 75,
			})
		})
		m.Col(6, func() {
			m.DataMatrixCode("https://github.com/prasangmisra/maroto", props.Rect{
				Percent: 100,
			})
		})
	})

	m.Row(40, func() {
		m.Col(2, func() {
			m.DataMatrixCode("https://github.com/prasangmisra/maroto", props.Rect{
				Center:  true,
				Percent: 50,
			})
		})
		m.Col(4, func() {
			m.DataMatrixCode("https://github.com/prasangmisra/maroto", props.Rect{
				Center:  true,
				Percent: 75,
			})
		})
		m.Col(6, func() {
			m.DataMatrixCode("https://github.com/prasangmisra/maroto", props.Rect{
				Center:  true,
				Percent: 100,
			})
		})
	})

	m.Row(40, func() {
		m.Col(6, func() {
			m.DataMatrixCode("https://github.com/prasangmisra/maroto", props.Rect{
				Percent: 50,
			})
		})
		m.Col(4, func() {
			m.DataMatrixCode("https://github.com/prasangmisra/maroto", props.Rect{
				Percent: 75,
			})
		})
		m.Col(2, func() {
			m.DataMatrixCode("https://github.com/prasangmisra/maroto", props.Rect{
				Percent: 100,
			})
		})
	})

	m.Row(40, func() {
		m.Col(6, func() {
			m.DataMatrixCode("https://github.com/prasangmisra/maroto", props.Rect{
				Center:  true,
				Percent: 50,
			})
		})
		m.Col(4, func() {
			m.DataMatrixCode("https://github.com/prasangmisra/maroto", props.Rect{
				Center:  true,
				Percent: 75,
			})
		})
		m.Col(2, func() {
			m.DataMatrixCode("https://github.com/prasangmisra/maroto", props.Rect{
				Center:  true,
				Percent: 100,
			})
		})
	})

	err := m.OutputFileAndClose("internal/examples/pdfs/dmgrid.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
