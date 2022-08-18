package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"

	g "github.com/AllenDang/giu"
	"github.com/fsnotify/fsnotify"
	"github.com/getlantern/systray"
	"github.com/tidwall/buntdb"
)

const (
	root_folder string = "/Users/doseo/Downloads"
)

var (
	db   *buntdb.DB
	dirs []string
)

func main() {
	//db, err := buntdb.Open(":memory:")
	//db, err = buntdb.Open("data.db")
	//if err != nil {
	//	log.Fatal(err)
	//}
	defer db.Close()
	FindDirs(root_folder)
	systray.Run(onReady, onExit)
}

func DirWatch(dirname string) {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	done := make(chan bool)
	go func() {
		for {
			select {
			case ev := <-watcher.Events:
				//log.Println("event:", ev)
				//fmt.Println(1, ev.Name)
				//fmt.Println(2, ev.Op)
				//fmt.Println(3, ev.String())
				//fmt.Println(4, ev.Op.String())
				if ev.Op.String() == "REMOVE" {
					fmt.Println(ev.Name + " is removed!!!!!")
				} else if ev.Op.String() == "RENAME" {
					fmt.Println(ev.Name + " is renamed!!!!!")
				} else if ev.Op.String() == "CREATE" {
					fmt.Println(ev.Name + " is created!!!!!")
				} else if ev.Op.String() == "WRITE" {
					fmt.Println(ev.Name + " is written!!!!!")
				} else if ev.Op.String() == "CHMOD" {
					fmt.Println(ev.Name + " is chmoded!!!!!")
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(dirname)
	if err != nil {
		log.Fatal(err)
	}

	<-done

}

func FindDirs(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
		return
	}
	go DirWatch(path)
	for _, f := range files {
		if f.IsDir() {
			fpath := path + "/" + f.Name()
			dirs = append(dirs, fpath)
			fmt.Println(fpath)
			go DirWatch(fpath)
			//FindDirs(fpath)
		}
	}
}
func walkitem(path string, info fs.FileInfo, err error) error {
	/*
		err := db.Update(func(tx *buntdb.Tx) error {
			_, _, err := tx.Set("mykey", "myvalue", nil)
			return err
		})
		path := "./path/to/fileOrDir"
		fileInfo, err := os.Stat(path)
		if err != nil {
			// error handling
		}

		if fileInfo.IsDir() {
			// is a directory
		} else {
			// is not a directory
		}

		err := db.View(func(tx *buntdb.Tx) error {
			err := tx.Ascend("", func(key, value string) bool {
				fmt.Printf("key: %s, value: %s\n", key, value)
				return true // continue iteration
			})
			return err
		})

		var delkeys []string
		tx.AscendKeys("object:*", func(k, v string) bool {
			if someCondition(k) == true {
				delkeys = append(delkeys, k)
			}
			return true // continue
		})
		for _, k := range delkeys {
			if _, err = tx.Delete(k); err != nil {
				return err
			}
		}
	*/
	return nil
}

func onReady() {
	systray.SetIcon(IconFwatch)
	systray.SetTitle("Fwatch App")
	systray.SetTooltip("Fwatch configuration")
	mControl := systray.AddMenuItem("Setting", "Setting the app")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	// Sets the icon of a menu item. Only available on Mac and Windows.
	mQuit.SetIcon(IconQuit)
	mControl.SetIcon(IconControl)
	go func() {
		for {
			select {
			case <-mControl.ClickedCh:
				go EditConfig()
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func onExit() {
	// clean up here
	systray.Quit()
}

func EditConfig() {
	wnd := g.NewMasterWindow("Hello world", 400, 200, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}

func loop() {
	g.SingleWindow().Layout(
		g.Label("Hello world from giu"),
		g.Row(
			g.Button("Click Me").OnClick(onClickMe),
			g.Button("I'm so cute").OnClick(onImSoCute),
		),
	)
}

func onClickMe() {
	fmt.Println("Hello world!")
}

func onImSoCute() {
	fmt.Println("Im sooooooo cute!!")
}
