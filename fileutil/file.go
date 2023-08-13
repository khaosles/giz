package fileutil

import (
	"bufio"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

/*
   @File: file.go
   @Author: khaosles
   @Time: 2023/8/12 11:25
   @Desc:
*/

// Exist judge whether exists filepath
func Exist(path string) bool {
	path = Format(path)
	// path stat
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

// IsFile judge whether is a file
func IsFile(path string) bool {
	path = Format(path)
	// path stat
	fileStat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fileStat.IsDir()
}

// IsDir judge whether is a dir
func IsDir(path string) bool {
	path = Format(path)
	// path stat
	fileStat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileStat.IsDir()
}

// Format the path
func Format(path string) string {
	// delete the space at the both ends
	path = strings.TrimSpace(path)
	// simplified path
	path = filepath.Clean(path)
	// \\ to /
	path = filepath.ToSlash(path)
	// / to \ or /
	path = filepath.FromSlash(path)
	return path
}

// FileSize obtain file size
func FileSize(path string) int64 {
	path = Format(path)
	if !IsFile(path) {
		return 0
	}
	fileStat, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fileStat.Size()
}

// Basename get filename
func Basename(path string) string {
	path = Format(path)
	return filepath.Base(path)
}

// Dirname get file dir name
func Dirname(path string) string {
	path = Format(path)
	return filepath.Dir(path)
}

// Join the path
func Join(elem ...string) string {
	return Format(filepath.Join(elem...))
}

// Split get file dir name
func Split(path string) (string, string) {
	path = Format(path)
	return filepath.Split(path)
}

// Suffix get file suffix
func Suffix(path string) string {
	path = Format(path)
	return filepath.Ext(path)
}

// Mkdir create a folder
func Mkdir(path string) error {
	path = Format(path)
	if !IsDir(path) {
		// create the folder
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

// MkdirP create a path's parents dir
func MkdirP(path string) error {
	return Mkdir(Dirname(path))
}

// Abs get file absolute path
func Abs(path string) (string, error) {
	path = Format(path)
	if filepath.IsAbs(path) {
		return path, nil
	}
	return filepath.Abs(path)
}

// Rmf  rm -f  file or folder
func Rmf(path string) error {
	if !Exist(path) {
		return nil
	}
	return os.RemoveAll(path)
}

// Rm remove a file
func Rm(path string) error {
	if !IsFile(path) {
		return nil
	}
	return os.Remove(path)
}

// Pwd get work path
func Pwd() string {
	rootPath, _ := os.Getwd()
	return rootPath
}

// ExecPath get the program`s directory
func ExecPath() string {
	rootPath, _ := os.Executable()
	return rootPath
}

// CurrentPath return current absolute path.
func CurrentPath() string {
	var absPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		absPath = path.Dir(filename)
	}

	return absPath
}

// Move a file from `src` to `dst`
func Move(src, dst string) error {
	return os.Rename(src, dst)
}

// Relvate get the relvative path of `targpath` to `basepath`
func Relvate(basepath, targpath string) (string, error) {
	return filepath.Rel(basepath, targpath)
}

// FilenameNoExt get the filename without subfix
func FilenameNoExt(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

// GenUniqueFilename gen a unique filename
func GenUniqueFilename(filename string, tries int, callback func(name string) string) (string, error) {
	if !Exist(filename) {
		return filename, nil
	}
	name := FilenameNoExt(filename)
	ext := filepath.Ext(filename)
	var newName string
	i := 1

	for {
		if callback != nil {
			newName = callback(name)
		} else {
			newName = fmt.Sprintf("%s(%d)", name, i)
		}
		newFilename := newName + ext
		if !Exist(newFilename) {
			return newFilename, nil
		}
		if i > tries {
			return "", errors.New("too many tries")
		}
		i++
	}
}

func CleanPath(p string) string {
	const stackBufSize = 128

	// Turn empty string into "/"
	if p == "" {
		return "/"
	}

	// Reasonably sized buffer on stack to avoid allocations in the common case.
	// If a larger buffer is required, it gets allocated dynamically.
	buf := make([]byte, 0, stackBufSize)

	n := len(p)

	// Invariants:
	//      reading from path; r is index of next byte to process.
	//      writing to buf; w is index of next byte to write.

	// path must start with '/'
	r := 1
	w := 1

	if p[0] != '/' {
		r = 0

		if n+1 > stackBufSize {
			buf = make([]byte, n+1)
		} else {
			buf = buf[:n+1]
		}
		buf[0] = '/'
	}
	trailing := n > 1 && p[n-1] == '/'
	// A bit more clunky without a 'lazybuf' like the path package, but the loop
	// gets completely inlined (bufApp calls).
	// So in contrast to the path package this loop has no expensive function
	// calls (except make, if needed).

	for r < n {
		switch {
		case p[r] == '/':
			// empty path element, trailing slash is added after the end
			r++

		case p[r] == '.' && r+1 == n:
			trailing = true
			r++

		case p[r] == '.' && p[r+1] == '/':
			// . element
			r += 2

		case p[r] == '.' && p[r+1] == '.' && (r+2 == n || p[r+2] == '/'):
			// .. element: remove to last /
			r += 3

			if w > 1 {
				// can backtrack
				w--
				if len(buf) == 0 {
					for w > 1 && p[w] != '/' {
						w--
					}
				} else {
					for w > 1 && buf[w] != '/' {
						w--
					}
				}
			}
		default:
			// Real path element.
			// Add slash if needed
			if w > 1 {
				bufApp(&buf, p, w, '/')
				w++
			}
			// Copy element
			for r < n && p[r] != '/' {
				bufApp(&buf, p, w, p[r])
				w++
				r++
			}
		}
	}

	// Re-append trailing slash
	if trailing && w > 1 {
		bufApp(&buf, p, w, '/')
		w++
	}
	// If the original string was not modified (or only shortened at the end),
	// return the respective substring of the original string.
	// Otherwise return a new string from the buffer.
	if len(buf) == 0 {
		return p[:w]
	}
	return string(buf[:w])
}

// Internal helper to lazily create a buffer if necessary.
// Calls to this function get inlined.
func bufApp(buf *[]byte, s string, w int, c byte) {
	b := *buf
	if len(b) == 0 {
		// No modification of the original string so far.
		// If the next character is the same as in the original string, we do
		// not yet have to allocate a buffer.
		if s[w] == c {
			return
		}

		// Otherwise use either the stack buffer, if it is large enough, or
		// allocate a new buffer on the heap, and copy all previous characters.
		if l := len(s); l > cap(b) {
			*buf = make([]byte, len(s))
		} else {
			*buf = (*buf)[:l]
		}
		b = *buf

		copy(b, s[:w])
	}
	b[w] = c
}

// ReadToString return string of file content.
func ReadToString(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// WriteStringToFile write string to target file.
func WriteStringToFile(filepath string, content string, append bool) error {
	flag := os.O_RDWR | os.O_CREATE
	if append {
		flag = flag | os.O_APPEND
	}

	f, err := os.OpenFile(filepath, flag, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	return err
}

// ReadByLine read file line by line.
func ReadByLine(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	result := make([]string, 0)
	buf := bufio.NewReader(f)

	for {
		line, _, err := buf.ReadLine()
		l := string(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		result = append(result, l)
	}

	return result, nil
}

// WriteBytesToFile write bytes to target file.
func WriteBytesToFile(filepath string, content []byte) error {
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(content)
	return err
}

// ListFiles return all file names in the path.
func ListFiles(path string) ([]string, error) {
	if !Exist(path) {
		return []string{}, nil
	}
	fs, err := os.ReadDir(path)
	if err != nil {
		return []string{}, err
	}
	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}
	result := make([]string, sz)
	for i := 0; i < sz; i++ {
		if !fs[i].IsDir() {
			result[i] = fs[i].Name()
		}
	}
	return result, nil
}

// CopyFile copy src file to dest file.
func CopyFile(srcPath string, dstPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	distFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer distFile.Close()

	var tmp = make([]byte, 1024*4)
	for {
		n, err := srcFile.Read(tmp)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		_, err = distFile.Write(tmp[:n])
		if err != nil {
			return err
		}
	}
}

// IsLink checks if a file is symbol link or not.
func IsLink(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeSymlink != 0
}

// FileMode return file's mode and permission.
func FileMode(path string) (fs.FileMode, error) {
	fi, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	return fi.Mode(), nil
}

// MiMeType return file mime type
// param `file` should be string(file path) or *os.File.
func MiMeType(file any) string {
	var mediatype string

	readBuffer := func(f *os.File) ([]byte, error) {
		buffer := make([]byte, 512)
		_, err := f.Read(buffer)
		if err != nil {
			return nil, err
		}
		return buffer, nil
	}

	if filePath, ok := file.(string); ok {
		f, err := os.Open(filePath)
		if err != nil {
			return mediatype
		}
		buffer, err := readBuffer(f)
		if err != nil {
			return mediatype
		}
		return http.DetectContentType(buffer)
	}

	if f, ok := file.(*os.File); ok {
		buffer, err := readBuffer(f)
		if err != nil {
			return mediatype
		}
		return http.DetectContentType(buffer)
	}
	return mediatype
}

// MTime returns file modified time.
func MTime(filepath string) (int64, error) {
	f, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	return f.ModTime().Unix(), nil
}

// Sha returns file sha value, param `shaType` should be 1, 256 or 512.
func Sha(filepath string, shaType ...int) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	h := sha1.New()
	if len(shaType) > 0 {
		if shaType[0] == 1 {
			h = sha1.New()
		} else if shaType[0] == 256 {
			h = sha256.New()
		} else if shaType[0] == 512 {
			h = sha512.New()
		} else {
			return "", errors.New("param `shaType` should be 1, 256 or 512.")
		}
	}

	_, err = io.Copy(h, file)

	if err != nil {
		return "", err
	}

	sha := fmt.Sprintf("%x", h.Sum(nil))

	return sha, nil
}

// ReadCsvFile read file content into slice.
func ReadCsvFile(filepath string) ([][]string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// WriteCsvFile write content to target csv file.
func WriteCsvFile(filepath string, records [][]string, append bool) error {
	flag := os.O_RDWR | os.O_CREATE

	if append {
		flag = flag | os.O_APPEND
	}

	f, err := os.OpenFile(filepath, flag, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	writer := csv.NewWriter(f)
	writer.Comma = ','

	return writer.WriteAll(records)
}
