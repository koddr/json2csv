# json2csv – Parse JSON files to CSV with data qualifier

[![Go version][go_version_img]][go_dev_url]
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

First, [download][go_download] and install **Go**. Version `1.20` (or higher)
is required.

Installation is done by using the [`go install`][go_install] command:

```console
go install github.com/koddr/json2csv@latest
```

> 💡 Note: See the repository's [Release page][repo_release_url], if you want
> to download a ready-made `deb`, `rpm`, `apk` or `Arch Linux` package.

GNU/Linux and macOS users available way to install via [Homebrew][brew_url]:

```console
# Tap a new formula:
brew tap koddr/tap

# Installation:
brew install koddr/tap/json2csv
```

Prepare folder with your data input in the `*.json` format, create JSON files
with intents and filter.

> 💡 Note: See the repository's [Wiki page][wiki_url] to understand
> structures of JSON files and get general recommendations for preparing the
> input data.

Next, run `json2csv` parser with options:

```console
json2csv \
  -json    /path/to/input/json/folder \
  -intents /path/to/intents-file.json \
  -filter  /path/to/filter-file.json \
  -output  /path/to/output/csv/folder \
  -content-field message \
  -min-word-len  5 \
  -chunk 1000
```

Done! 🎉 Output CSV file(s) has a default comma (`,`) separators between columns.

### 🐳 Docker-way to quick start

If you don't want to physically install `json2csv` to your system, you feel
free to using our [official Docker image][docker_image_url] and run it from
isolated container:

```console
docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} koddr/json2csv:latest [COMMANDS]
```

## 🧩 Options

| Option | Description | Is required? | Default value | Docs link |
| --- | --- | --- | --- | --- |
| `-json` | set a **path** to the folder with JSON source file(s) | no | `./json_files` | [Wiki][wiki_json_folder_url] |
| `-intents` | set a **path** to the file with intents | no | `./intents.json` | [Wiki][wiki_intents_url] |
| `-filter` | set a **path** to the file with a filter | no | `./filter.json` | [Wiki][wiki_filter_url] |
| `-output` | set a **path** to the output folder where the prepared CSV file(s) will be placed | no | `./output_files` | [Wiki][wiki_output_folder_url] |
| `-content-field` | set a **name of the content field** (attribute that contains string to qualify and filter) in JSON source file(s) | no | `content` | |
| `-min-word-len` | set a **min word length count** to filter input words (if a word is smaller than this option, it will be skipped) | no | `0` | |
| `-chunk` | set a **chunk size** for one CSV file | no | `5000` | |

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

[`json2csv`][repo_url] is free and open-source software licensed under the
[Apache 2.0 License][license_url], created and supported with 🩵 for people and
robots by [Vic Shóstak][author].

[go_download]: https://golang.org/dl/
[go_install]: https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies
[go_version_img]: https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go
[go_report_img]: https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none
[go_report_url]: https://goreportcard.com/report/github.com/koddr/json2csv
[go_dev_url]: https://pkg.go.dev/github.com/koddr/json2csv
[code_coverage_img]: https://img.shields.io/badge/code_coverage-in_progress-success?style=for-the-badge&logo=none
[brew_url]: https://brew.sh
[docker_image_url]: https://hub.docker.com/repository/docker/koddr/json2csv
[wiki_img]: https://img.shields.io/badge/docs-wiki_page-blue?style=for-the-badge&logo=none
[wiki_url]: https://github.com/koddr/json2csv/wiki
[wiki_intents_url]: https://github.com/koddr/json2csv/wiki#intents
[wiki_filter_url]: https://github.com/koddr/json2csv/wiki#filter
[wiki_json_folder_url]: https://github.com/koddr/json2csv/wiki#folder-with-json-files
[wiki_output_folder_url]: https://github.com/koddr/json2csv/wiki#folder-with-output-csv-files
[license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none
[license_url]: https://github.com/koddr/json2csv/blob/main/LICENSE
[repo_url]: https://github.com/koddr/json2csv
[repo_release_url]: https://github.com/koddr/json2csv/releases
[author]: https://github.com/koddr
