package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/kr/pretty"
	"gitlab.b2w/team/alpha/omega.git/aggregation/sitemanager"
)

type AnywayReader struct {
	reader io.Reader
}

func (r *AnywayReader) FetchContent(context.Context) (content []byte, err error) {
	return ioutil.ReadAll(r.reader)
}

func NewAnywayReader(r io.Reader) *AnywayReader {
	return &AnywayReader{r}
}

func main() {
	// Open our xmlFile
	xmlFile, err := os.Open("/home/msantanaz/Downloads/menu.xml")
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	anywayreader := NewAnywayReader(xmlFile)

	_ = anywayreader

	sm, err := sitemanager.NewClient(
		sitemanager.SetObjectFetcher(anywayreader),
	)
	if err != nil {
		panic(err)
	}

	cc := sitemanager.NewClientCategory(sm)

	aggs, err := cc.FetchByCategoryId(context.Background(), "443456")
	if err != nil {
		panic(err)
	}

	fmt.Println(aggs)
	pretty.Println(aggs)
}
