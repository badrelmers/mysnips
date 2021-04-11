package util

import(
    "os"
    "os/exec"
    "log"
    
    "io"
    "io/ioutil"
    "encoding/json"
    "bytes"

    "fmt"
    "math/rand"
    "time"

)

// https://www.digitalocean.com/community/tutorials/how-to-write-packages-in-go
// Exported Code
// You may have noticed that all of the declarations in the greet.go file you called were capitalized. Go does not have the concept of public, private, or protected modifiers like other languages do. External visibility is controlled by capitalization. Types, variables, functions, and so on, that start with a capital letter are available, publicly, outside the current package. A symbol that is visible outside its package is considered to be exported.

// If you add a new method to Octopus called reset, you can call it from within the greet package, but not from your main.go file, which is outside the greet package:

var ( INFOC = "\033[32m"; WARNC = "\033[33m"; ERRORC = "\033[0;1;31m"; HIDEC = "\033[37m"; INFO2C = "\033[36m"; INFO3C = "\033[0;1;35m"; INFO4C = "\033[0;1;34m";        INFOCB = "\033[0;30m\033[42m"; WARNCB = "\033[0;1;33;40;7m"; ERRORCB = "\033[0;1;37m\033[41m"; HIDECB = "\033[0;1;30m\033[47m"; INFO2CB = "\033[0;30m\033[46m"; INFO3CB = "\033[0;1;37m\033[45m"; INFO4CB = "\033[0;1;37m\033[44m";        ENDC = "\033[0m" 
)

func Test() {
    text := "text %s text\n"
    intt := 1234
    
    byte := []byte("byte\n")

    mapp := map[int]string{}
    mapp[1] = "map 1 %s test"
    mapp[2] = "map 2\n"

    jsonn,_ := json.Marshal(mapp)

    slicee := []string{"one", "two", "three"}
    slicee2 := []int{10, 200, 3000}


    os.Remove("met2.t")
    Append("met2.t",text)
    Append("met2.t",intt)
    Append("met2.t","\n")
    Append("met2.t",byte)
    Append("met2.t",mapp)
    Append("met2.t","\n")
    Append("met2.t",jsonn)
    Append("met2.t","\n")
    Append("met2.t",slicee)
    Append("met2.t",slicee2)

}



// ####################################################################
// https://github.com/go-rod/rod/blob/master/lib/utils/utils.go
// OutputFile auto creates file if not exists, it will try to detect the data type and
// auto output binary, string or json
func Write(p string, data interface{}) error {
    var bin []byte

    switch t := data.(type) {
    case []byte:
        bin = t
    case string:
        bin = []byte(t)
    case io.Reader:
        f, _ := os.OpenFile(p, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
        _, err := io.Copy(f, t)
        defer f.Close()
        return err
    default:
        // this will be used by map and json +++++
        bin = MustToJSONBytes(data)
    }

    return ioutil.WriteFile(p, bin, 0664)
}

func Append(p string, data interface{}) error {
    var bin []byte

    switch t := data.(type) {
    case []byte:
        bin = t
    case string:
        bin = []byte(t)
    case io.Reader:
        f, _ := os.OpenFile(p, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
        _, err := io.Copy(f, t)
        defer f.Close()
        return err
    default:
        // this will be used by map and json +++++
        bin = MustToJSONBytes(data)
    }

    // return ioutil.WriteFile(p, bin, 0664)


    // If the file doesn't exist, create it, or append to the file
    file, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)

    if err != nil {
        log.Fatalf("error while opening the file. %v", err)
    }

    // close the file once program execution complete
    defer file.Close()

    if _, err := file.Write([]byte(bin)); err != nil {
        log.Fatalf("error while writing the file. %v", err)
    }

    return err
}

// Mkdir makes dir recursively
func Mkdir(path string) error {
    return os.MkdirAll(path, 0775)
}

// MustToJSONBytes encode data to json bytes
func MustToJSONBytes(data interface{}) []byte {
    buf := bytes.NewBuffer(nil)
    enc := json.NewEncoder(buf)
    enc.SetEscapeHTML(false)
    E(enc.Encode(data))
    b := buf.Bytes()
    return b[:len(b)-1]
}

// E if the last arg is error, panic it
func E(args ...interface{}) []interface{} {
    err, ok := args[len(args)-1].(error)
    if ok {
        panic(err)
    }
    return args
}

// ########################################################################
func Rename(oldName, newName string) {
    err := os.Rename(oldName, newName)
    Err(err)
}


func Mv(oldLocation, newLocation string) {
    // os.Rename() can also move file from one location to another at same time renaming file name.
    err := os.Rename(oldLocation, newLocation)
    Err(err)
}

func Cp(src,dst string) {
    sourceFile, err := os.Open(src)
    Err(err)
    defer sourceFile.Close()
 
    // Create new file
    newFile, err := os.Create(dst)
    Err(err)
    defer newFile.Close()
 
    bytesCopied, err := io.Copy(newFile, sourceFile)
    Err(err)
    log.Printf("Copied %d bytes.", bytesCopied)
}

// TODO:
// https://stackoverflow.com/questions/51779243/copy-a-folder-in-go
// https://github.com/moby/moby/blob/master/daemon/graphdriver/copy/copy.go
// https://github.com/plus3it/gorecurcopy
// https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04
// https://github.com/TryStreambits/coreutils/blob/master/io.go
func cpdir(){
// TODO
}


func Rm(file string) {
    err := os.Remove(file)
    Err(err)
}

// delete directory and its contents with os.RemoveAll
// The RemoveAll removes the directory and its contents recursively.
func Rmrf(file string) {
    err := os.RemoveAll(file)
    Err(err)
}


// Temporary files and direct­ories
// create temp dir
func TempDir(dir, prefix string) string{
    dir, err := ioutil.TempDir(dir, prefix)
    Err(err)
    fmt.Println("Temp dir created:", dir)
    return dir
}

// create temp file
func TempFile(dir, prefix string) string{
    f,err := ioutil.TempFile(dir, prefix)
    Err(err)
    fmt.Println("Temp file created:", f.Name())
    return f.Name()
}


func RandInt(min int, max int) int {
    // https://stackoverflow.com/questions/44659343/how-to-choose-a-random-number-from-the-range
    // Seed should be set once, better spot is func init()
    // TODO: osea funcionara si la pongo arriba tb? 
    rand.Seed(time.Now().UTC().UnixNano())

    return min + rand.Intn(max-min)
}


// #####################################################################
// https://github.com/TryStreambits/coreutils/blob/master/exec.go
// ExecCommand executes a command with args and returning the stringified output
func ExecCommand(command string, args []string, redirect bool) string {
    if ExecutableExists(command) { // If the executable exists
        var output []byte
        runner := exec.Command(command, args...)

        if redirect { // If we should redirect output to var
            output, _ = runner.CombinedOutput() // Combine the output of stderr and stdout
        } else {
            runner.Stdout = os.Stdout
            runner.Stderr = os.Stderr
            runner.Run()
        }

        return string(output[:])
    } else { // If the executable doesn't exist
        return command + " is not an executable."
    }
}

// ExecutableExists checks if an executable exists
func ExecutableExists(executableName string) bool {
    _, existsErr := exec.LookPath(executableName)
    return (existsErr == nil)
}

// #####################################################################
func Pause() {
    // time.Sleep(time.Hour)
    fmt.Println("")
    fmt.Println("")
    fmt.Println("press enter to continue... no uses ctrl+c TODO add ctrl+c to close all too")
    fmt.Scanln() // pause and wait for Enter Key
}

func Pauseexit() {
    // time.Sleep(time.Hour)
    fmt.Println("")
    fmt.Println("")
    fmt.Println("press enter to exit... no uses ctrl+c TODO add ctrl+c to close all too")
    fmt.Scanln() // pause and wait for Enter Key
    os.Exit(0)
}

func Err(err error) {
    if err != nil {
        fmt.Println(ERRORC, err, ENDC)

        // os.Exit(1)
        panic(err) // panic es mejor ke exit pk ejecuta defer functions ke son como trap in bash
    }
}

func GetDateNow() string {
    t := time.Now()
    return fmt.Sprintf("%d%02d%02d_%02d%02d%02d",t.Year(), t.Month(), t.Day(),t.Hour(), t.Minute(), t.Second()) 
}
