# Go_markdown-directory-snapshot

Description: Snapshot a directory & save non-excluded results of snapshot in an output.md file

# How to use

First, ensure you have [installed Go](https://go.dev/dl/).

In your terminal, navigate to the directory where you want to apply this project, and type the following commands:

```bash
git clone https://github.com/gooddavvy/Go_markdown-directory-snapshot
go mod init [your_module_name]
```

Be sure to replace `your-module-name` with your actual module name.

Also, change the name of `test_directory` to the name that fits your needs, and replace its contents to the contents you actually need. But, be sure to also adjust the `main` function in `main.go`, for example:

```go
func main() {
	// Define the root path and ignore list
	rootPath := "test_directory" // < replace this if needed
	ignoreList := []string{
		"ignore_this_file.txt",
		"ignore_this_directory",
		"accept_this_directory/ignore_this_thing.txt",
	} // ^ Replace those if needed

    // v Leave the rest as it is unless you want to add extra logic.
	// Call the function to generate the Markdown snapshot
	err := generateMarkdownSnapshot(rootPath, ignoreList)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
```

Run it using the following command:

```bash
go run main.go
```

An `output.md` file will be created at the root level of your project, containing a snapshot of non-ignored files and their contents.

Please let me know (in the [Issues Section](https://github.com/gooddavvy/Go_markdown-directory-snapshot/issues)) if you encounter any issues during setup.

# What's coming soon

An extremely amazing application, with a PythonStreamlit-powered frontend and GolangFiber-powered backend, based on this program, is coming soon, so stay tuned! [Follow me](https://github.com/gooddavvy) to get a notification immediately after I release it to the public.
