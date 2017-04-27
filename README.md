# gomin
Merge multiple Go files into one

Running the compiled program will read all productive *.go-files (not ending on _test.go) recursively from the directory you're currently in and merges them into a single sourcecode string.
The merged sourcecode is copied to clipboard.

* Tested under Windows
* Scans all files from Execution path (current dir of your cmd/shell)
  * only processes *.go-files which are not *_test.go-files
  * scans package names
* Merges multiple self-written packages into one file, replacing import declarations if suitable

Currently known (and accepted) bug: Using an import prefix which is hidden by a Variable name may cause unwanted replacements and merging will result into a non-compilable sourceode.
