# json2csv – Parse JSON files to CSV with data qualifier

The parser can read given folder with `*.json` files, filtering and 
qualifying input data with intent & stop words dictionaries and save results 
to CSV files by given chunk size.

Minimal dependency on other Go packages, maximum performance even on large 
amounts of input data.

## ⚡️ Quick start

Install package:

```bash
go install https://github.com/koddr/json2csv@latest
```

Next, run `json2csv` parser:

```bash
json2csv \
  -json    /path/to/input/json/folder \
  -intents /path/to/intents-file.json \
  -filter  /path/to/filter-file.json \
  -output  /path/to/output/csv/folder \
  -content data \
  -min-word-len 5 \
  -chunk 1000
```

> 💡 Note: output CSV files have a default comma (`,`) separators between
> columns.

## ✨ Solving case

Imagine that you have many JSON files of the same type in a folder on your
computer. The structure of each such file is identical and can be described
in a single Go structure.

But your task is not only to turn these files into a structured CSV, but 
also to make some qualification of incoming data based on dictionaries with 
intent.

The output should be CSV files (with a specified number of lines in each),
where the first column contains the original phrases from JSON files, and
the second column will contain the associated intents after qualification.

This is what this Go package solves! ✌️

## ⚠️ License

`json2csv` is free and open-source software licensed under the
[Apache 2.0 License](LICENSE), created and supported with 🩵 for people and
robots by [Vic Shóstak](https://github.com/koddr).
