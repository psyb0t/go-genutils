package genutils

import (
    "os"
    "bytes"
    "io/ioutil"
    "compress/flate"
)

func RevStrSlice(slc []string) []string {
    for i, j := 0, len(slc)-1; i < j; i, j = i+1, j-1 {
        slc[i], slc[j] = slc[j], slc[i]
    }

    return slc
}

func StringInSlice(str string, slc []string) bool {
    for _, x := range(slc) {
        if x == str {
            return true
        }
    }

    return false
}

func MkDirAll(path string) error {
    _, err := os.Stat(path)

    if os.IsNotExist(err) {
        err = os.MkdirAll(path, 0644)

        if err != nil {
            return err
        }
    }

    if err != nil {
        return err
    }

    return nil
}

func Compress(to_compress []byte) ([]byte, error) {
    compressed := new(bytes.Buffer)
    c, err := flate.NewWriter(compressed, flate.BestCompression)

    if err != nil {
        return compressed.Bytes(), err
    }

    c.Write(to_compress)
    c.Close()

    return compressed.Bytes(), nil
}

func Decompress(to_decompress []byte) ([]byte, error) {
    c := new(bytes.Buffer)
    c.Write(to_decompress)

    d := flate.NewReader(c)
    decompressed, err := ioutil.ReadAll(d)
    d.Close()

    if err != nil {
        return nil, err
    }

    return decompressed, nil
}
