package main

import (
    "fmt"
    "html/template"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "strings"
)

const (
    serverPort       = 520
    webPath          = "./520"
    mixesDir         = "mixes"
    playlistFileName = "playlist.txt"
)

var (
    indexTemplate *template.Template
    playTemplate  *template.Template
    mixesPath     = fmt.Sprintf("%s/%s", webPath, mixesDir)
)

type Mix struct {
    Name           string
    URLEscapedPath string
}

type IndexTemplate struct {
    Mixes []Mix
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    var mixes []Mix
    currentMixes, err := ioutil.ReadDir(mixesPath)
    if err != nil {
        log.Println(fmt.Sprintf("%s failed reading directory %s for mixes", err, mixesDir))
        w.Write([]byte("Failed to index mixes"))
        return
    }
    for _, mix := range currentMixes {
        mixName := mix.Name()
        // skip *nix special files
        if strings.HasPrefix(mixName, ".") {
            continue
        }
        mix := Mix{
            Name:           mixName,
            URLEscapedPath: fmt.Sprintf("%s/%s", mixesDir, mixName),
        }
        mixes = append(mixes, mix)
    }
    indexData := IndexTemplate{
        Mixes: mixes,
    }
    indexTemplate.Execute(w, indexData)
}

type PlayTemplate struct {
    VideoDirectory string
    PlaylistURI    string
}

func playHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        queryParams := r.URL.Query()
        mix := queryParams.Get("mix")
        if mix == "" {
            w.Write([]byte("No mix to play specified"))
            return
        }
        rawVideoDirectory, err := url.QueryUnescape(mix)
        if err != nil {
            w.Write([]byte(fmt.Sprintf("Invalid mix to play specified %s", mix)))
            return
        }
        var playlistURI string
        playlistFilePath := fmt.Sprintf("%s/%s/%s", webPath, rawVideoDirectory, playlistFileName)
        playlistFileBytes, err := ioutil.ReadFile(playlistFilePath)
        if err != nil {
            log.Println(fmt.Sprintf("error %s trying to read file %s", err, playlistFilePath))
        } else {
            playlistURI = strings.TrimSpace(string(playlistFileBytes))
        }
        playData := PlayTemplate{
            VideoDirectory: rawVideoDirectory,
            PlaylistURI:    playlistURI,
        }
        playTemplate.Execute(w, playData)
    }
}

func main() {
    indexTemplate = template.Must(template.ParseFiles(fmt.Sprintf("%s/index.html", webPath)))
    http.HandleFunc("/", indexHandler)
    playTemplate = template.Must(template.ParseFiles(fmt.Sprintf("%s/play.html", webPath)))
    http.HandleFunc("/play", playHandler)
    http.Handle("/mixes/", http.FileServer(http.Dir(webPath)))
    http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil)
}
