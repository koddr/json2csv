# json2csv – Parse JSON files to CSV with data qualifier

![Go version][go_version_img]
[![Go report][go_report_img]][go_report_url]
![Code coverage][code_coverage_img]
[![Wiki][wiki_img]][wiki_url]
[![License][license_img]][license_url]

The parser can read given folder with `*.json` files, filtering and 
qualifying input data with intent & stop words dictionaries and save results 
to CSV files by given chunk size.

**Minimal** dependency on other Go packages, but **maximum** performance 
even on large amounts of the input JSON data.

## ⚡️ Quick start

First, [download][go_download] and install **Go**. Version `1.20` or higher 
is required. 

Installation is done by using the [`go install`][go_install] command:

```bash
go install https://github.com/koddr/json2csv@latest
```

Prepare folder with your data source (format `*.json`) and create JSON files 
with:

- intents (for ex., `intents-file.json`); 
- filter (for ex., `filter-file.json`);

> 💡 Note: see the [`Wiki`][wiki_url] page to understand structures of 
> JSON files and get general recommendations for preparing the input data. 

Next, run `json2csv` parser with (or without) options:

```bash
json2csv \
  -json    /path/to/input/json/folder \
  -intents /path/to/intents-file.json \
  -filter  /path/to/filter-file.json \
  -output  /path/to/output/csv/folder \
  -content-field message \
  -min-word-len  5 \
  -chunk 1000
```

> 💡 Note: output CSV file has a default comma (`,`) separators between columns.

## 🧩 Options

- `-json [path]` is an option to set a path to the folder with JSON source
  file(s);
- `-intents [path]` is an option to set a path to the file with intents (see 
  [Wiki][wiki_intents_url]);
- `-filter [path]` is an option to set a path to the file with a filter (see 
  [Wiki][wiki_filter_url]);
- `-content-field [string]` is an option to set a name of the content field
  (attribute that contains string to qualify and filter) in JSON source file(s);
- `-min-word-len [int]` is an option to set a min word length count to filter
  input words (if a word is smaller than this option, it will be skipped);
- `-chunk [int]` is an option to set a chunk size for one CSV file;

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

[`json2csv`][json2csv_url] is free and open-source software licensed under the
[Apache 2.0 License][license_url], created and supported with 🩵 for people and
robots by [Vic Shóstak][author].

[go_download]: https://golang.org/dl/
[go_install]: https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies
[go_version_img]: https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go
[go_report_img]: https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none
[go_report_url]: https://goreportcard.com/report/github.com/koddr/json2csv
[code_coverage_img]: https://img.shields.io/badge/code_coverage-98%25-success?style=for-the-badge&logo=none
[wiki_img]: https://img.shields.io/badge/-wiki_page-blue?style=for-the-badge&logo=wikipedia
[wiki_url]: https://github.com/koddr/json2csv/wiki
[wiki_intents_url]: https://github.com/koddr/json2csv/wiki#intents
[wiki_filter_url]: https://github.com/koddr/json2csv/wiki#filter
[license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none
[license_url]: https://github.com/koddr/json2csv/blob/main/LICENSE
[json2csv_url]: https://github.com/koddr/json2csv
[author]: https://github.com/koddr
