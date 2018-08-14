// compile for Windows with:
// GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -ldflags "-H windowsgui" -o ImgViewer.exe
// source: https://github.com/zserge/webview/issues/22#issuecomment-332153397

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/zserge/webview"

	"github.com/apahl/utls"

	"github.com/apahl/img_viewer/config"
)

const (
	version      = "0.2.0"
	windowWidth  = 1200
	windowHeight = 600
)

var confDir = config.GetConfigDir()

func writeConfig(slot, plateDir, plate, well string) {
	slotFn := filepath.FromSlash(confDir + fmt.Sprintf("/slot_%s.txt", slot))
	bytes := []byte(fmt.Sprintf("%s,%s,%s", plateDir, plate, well))
	err := ioutil.WriteFile(slotFn, bytes, 0644)
	utls.QuitOnError(err)
}

func loadConfig(w webview.WebView) {
	pathIsSet := false
	for slot := 1; slot <= 6; slot++ {
		slotFn := filepath.FromSlash(confDir + fmt.Sprintf("/slot_%d.txt", slot))
		dat, err := ioutil.ReadFile(slotFn)
		if err != nil {
			continue
		}
		txt := string(dat)
		comp := strings.Split(txt, ",")
		plateDir := comp[0]
		plate := comp[1]
		well := comp[2]
		evalPlate := fmt.Sprintf("document.getElementById('plate-id%d').value='%s'", slot, plate)
		w.Eval(evalPlate)
		evalWell := fmt.Sprintf("document.getElementById('well-id%d').value='%s'", slot, well)
		w.Eval(evalWell)
		if !pathIsSet {
			pathIsSet = true
			plateDir = filepath.ToSlash(plateDir)
			evalPath := fmt.Sprintf("document.getElementById('basepath').value='%s'", plateDir)
			w.Eval(evalPath)
		}
	}
}

// writeUI writes the UI to disk and returns the full file name
func writeUI() string {
	uiFn := filepath.ToSlash(os.TempDir() + "/plate_viewer_ui.html")
	bytes := []byte(indexHTML)
	err := ioutil.WriteFile(uiFn, bytes, 0644)
	utls.QuitOnError(err)
	return uiFn
}

func next(well string, offset int) string {
	result := ""
	if len(well) < 2 || len(well) > 3 {
		return result
	}
	rowNum := well[0] - 64
	if rowNum < 1 || rowNum > 24 {
		return result
	}
	colNum, err := strconv.Atoi(well[1:])
	if err != nil {
		return result
	}
	colNum += offset
	if colNum > 24 {
		colNum -= 24
		rowNum++
		if rowNum > 16 {
			rowNum -= 16
		}
	} else if colNum < 1 {
		colNum += 24
		rowNum--
		if rowNum < 1 {
			rowNum += 16
		}
	}
	var pad string
	if colNum < 10 {
		pad = "0"
	} else {
		pad = ""
	}
	result = string(rowNum+64) + pad + fmt.Sprint(colNum)
	return result
}

func loadPlate(w webview.WebView, args string) {
	comp := strings.Split(args, ",")
	slot := comp[0]
	plateDir := comp[1]
	plateDir = strings.TrimSuffix(filepath.ToSlash(plateDir), "/")
	plate := comp[2]
	well := comp[3]
	width := "100"
	imgFn := plateDir + "/" + plate + "/images/" + well
	evalWell := fmt.Sprintf("document.getElementById('well%s').innerHTML='%s'", slot, well)
	writeConfig(slot, plateDir, plate, well)
	w.Eval(evalWell)
	for ch := 1; ch <= 5; ch++ {
		imgSrc := fmt.Sprintf(`<img width="%s%%" src="file://%s_w%d.jpg" />`, width, imgFn, ch)
		// fmt.Println(imgSrc)
		evalW := fmt.Sprintf("document.getElementById('w%s_%d').innerHTML='%s'", slot, ch, imgSrc)
		w.Eval(evalW)
	}
}

func loadSync(w webview.WebView, args string) {
	comp := strings.Split(args, ",")
	plateDir := string(comp[0])
	plateDir = strings.TrimSuffix(filepath.ToSlash(plateDir), "/")
	plates := comp[1:7]
	wells := comp[7:13]
	offset, err := strconv.Atoi(comp[13])
	if err != nil {
		log.Println("loadSync: ", err)
	}
	for idx := range plates {
		slot := idx + 1
		plate := plates[idx]
		well := wells[idx]
		nextWell := next(well, offset)
		evalWell := fmt.Sprintf("document.getElementById('well-id%d').value='%s'", slot, nextWell)
		w.Eval(evalWell)
		loadArgs := fmt.Sprintf("%d,%s,%s,%s", slot, plateDir, plate, nextWell)
		loadPlate(w, loadArgs)
	}
}

func nextWell(w webview.WebView, args string) {
	comp := strings.Split(args, ",")
	slot := comp[0]
	plateDir := comp[1]
	plateDir = strings.TrimSuffix(filepath.ToSlash(plateDir), "/")
	plate := comp[2]
	well := comp[3]
	offset, err := strconv.Atoi(comp[4])
	if err != nil {
		log.Println("nextWell: ", err)
	}
	autoload := comp[5]
	nextWell := next(well, offset)
	evalWell := fmt.Sprintf("document.getElementById('well-id%s').value='%s'", slot, nextWell)
	w.Eval(evalWell)
	if autoload == "true" {
		loadArgs := fmt.Sprintf("%s,%s,%s,%s", slot, plateDir, plate, nextWell)
		loadPlate(w, loadArgs)
	}
}

func handleRPC(w webview.WebView, data string) {
	switch {
	case data == "close":
		w.Terminate()
	case strings.HasPrefix(data, "config:"):
		loadConfig(w)
	case strings.HasPrefix(data, "load:"):
		loadPlate(w, strings.TrimPrefix(data, "load:"))
	case strings.HasPrefix(data, "sync:"):
		loadSync(w, strings.TrimPrefix(data, "sync:"))
	case strings.HasPrefix(data, "next:"):
		nextWell(w, strings.TrimPrefix(data, "next:"))
	}
}

func main() {
	// fmt.Println("confDir: ", confDir)
	err := os.MkdirAll(confDir, 0755)
	if err != nil {
		log.Println(err)
	}
	uiFn := writeUI()
	url := "file://" + uiFn
	// fmt.Println(url)

	w := webview.New(webview.Settings{
		Width:     windowWidth,
		Height:    windowHeight,
		Title:     "Plate Image Viewer v" + version,
		Resizable: true,
		URL:       url,
		ExternalInvokeCallback: handleRPC,
	})
	w.SetColor(255, 255, 255, 255)
	defer w.Exit()
	// loadConfig(w)
	w.Run()
}
