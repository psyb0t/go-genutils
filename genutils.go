package genutils

import (
    "bytes"
    "io/ioutil"
    "compress/flate"
)

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
    compressed := new(bytes.Buffer)
    compressed.Write(to_decompress)

    d := flate.NewReader(compressed)
    decompressed, err := ioutil.ReadAll(d)
    d.Close()

    if err != nil {
        return nil, err
    }

    return decompressed, nil
}
