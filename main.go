package main

import (
    "strings"
    "bufio"
    "sort"
	"fmt"
    "os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type GraphData struct {
    lbls []string
    cnts []opts.BarData
}

type SortGraphDataByCount GraphData
func (a SortGraphDataByCount) Len() int {
    return len(a.lbls)
}
func (a SortGraphDataByCount) Swap(i, j int) {
    a.lbls[i], a.lbls[j] = a.lbls[j], a.lbls[i]
    a.cnts[i], a.cnts[j] = a.cnts[j], a.cnts[i]
}
func (a SortGraphDataByCount) Less(i, j int) bool {
    return a.cnts[i].Value.(uint) > a.cnts[j].Value.(uint)
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    input, _ := reader.ReadString('\n')
    input = strings.Replace(input, "\n", "", -1)

    freq := map[rune]uint{}
    for _, c := range input {
        if _, exists := freq[c]; exists {
            freq[c]++
        } else {
            freq[c] = 1
        }
    }

    var i uint = 0
    var lbls = make([]string, len(freq))
    var cnts = make([]opts.BarData, len(freq))
    for k, v := range freq { 
        lbls[i] = string(k)
        cnts[i] = opts.BarData{Value: v}
        i++
    }

    srtd_data := GraphData{ lbls: lbls, cnts: cnts }
    sort.Sort(SortGraphDataByCount(srtd_data))

	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Character Frequency",
		Subtitle: fmt.Sprintf("Input: %s", input),
	}))
    bar.SetXAxis(lbls)
    bar.AddSeries("", cnts)

	f, _ := os.Create("bar-graph.html")
	bar.Render(f)
    println("Bar graph rendered!")
}