package root

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/cli/flag"
	log "github.com/sirupsen/logrus"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// MergeFiles is the Cobra command definition.
var MergeFiles = &cobra.Command{
	Use:   "merge-files",
	Short: "Merge all file content from a source directory into a single file, with optional exclusions",
	Run:   mergeFilesHandler,
}

func init() {
	// Add flags to the command.
	// 1. --source-dir: The directory to walk.
	// 2. --output-file: The file where the merged content should be written.
	// 3. --exclude-file-extension: Optional repeated flag to exclude certain file extensions.
	MergeFiles.PersistentFlags().String(string(flag.SourceDir), "", "Path to the source directory containing files")
	MergeFiles.PersistentFlags().String(string(flag.OutFile), "", "Path (including filename) of the merged file")
	MergeFiles.PersistentFlags().StringArray(string(flag.ExcludeFileExtension), []string{}, "File extensions to exclude (e.g. .txt, .pdf)")
}

// mergeFilesHandler is called when the 'merge-files' command is invoked.
func mergeFilesHandler(cmd *cobra.Command, args []string) {
	// Extract flag values
	sourceDir, err := cmd.Flags().GetString(string(flag.SourceDir))
	flag.HandleFlagErrAndVal(err, flag.SourceDir, sourceDir)
	outputFile, err := cmd.Flags().GetString(string(flag.OutFile))
	flag.HandleFlagErrAndVal(err, flag.OutFile, outputFile)
	excludes, err := cmd.Flags().GetStringArray(string(flag.ExcludeFileExtension))
	flag.HandleFlagErr(err, flag.ExcludeFileExtension)

	// Perform the merge
	if err := mergeAllFiles(sourceDir, outputFile, excludes); err != nil {
		log.Fatalf("Failed to merge files: %v", err)
	}

	log.Printf("All files from %q have been merged into %q\n", sourceDir, outputFile)
}

// mergeAllFiles walks through 'sourceDir', finds all files, and merges them into 'outputFile',
// skipping any file whose extension is in 'excludes' or if it's the output file itself.
// mergeAllFiles walks through 'sourceDir', finds all files, and merges them into 'outputFile'.
// It excludes any file whose name starts with a dot (i.e., "hidden files").
func mergeAllFiles(sourceDir, outputFile string, excludes []string) error {
	// Remove the output file if it exists, so we start fresh.
	if err := os.Remove(outputFile); err != nil && !os.IsNotExist(err) {
		return errors.Wrapf(err, "could not remove existing output file %q", outputFile)
	}

	// Create the output file
	out, err := os.Create(outputFile)
	if err != nil {
		return errors.Wrapf(err, "could not create output file %q", outputFile)
	}
	defer out.Close()

	// Normalize the excluded extensions for case-insensitive matching (if that's what you want).
	normalizedExcludes := make([]string, len(excludes))
	for i, ext := range excludes {
		normalizedExcludes[i] = strings.ToLower(ext)
	}

	// Walk through sourceDir, processing each file
	err = filepath.WalkDir(sourceDir, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return errors.Wrap(walkErr, "failed to walk directory")
		}

		// Skip hidden directories or files (their name starts with a dot).
		// e.g., .DS_Store, .git, .vscode, .env, etc.
		baseName := filepath.Base(path)
		if strings.HasPrefix(baseName, ".") {
			// If it's a directory starting with ".", skip it entirely.
			if d.IsDir() {
				return filepath.SkipDir
			}
			// If it's a file starting with ".", skip it and move on.
			return nil
		}

		// Skip directories (just keep walking deeper) but don't process them directly.
		if d.IsDir() {
			return nil
		}

		// Skip if the path is the output file itself (avoid infinite loops).
		if path == outputFile {
			return nil
		}

		// Check if the file’s extension matches any "excludes"
		fileExt := strings.ToLower(filepath.Ext(path))
		for _, excludedExt := range normalizedExcludes {
			if fileExt == excludedExt {
				return nil
			}
		}

		// Write a marker in the output file for this file
		dir := filepath.Dir(path)
		fileName := filepath.Base(path)
		if _, err := fmt.Fprintf(out, "## %s/%s\n", dir, fileName); err != nil {
			return errors.Wrapf(err, "failed to write marker for file %q", path)
		}

		// Read the file
		in, err := os.Open(path)
		if err != nil {
			return errors.Wrapf(err, "failed to open file %q", path)
		}

		// Copy its contents into the output
		if _, err := io.Copy(out, in); err != nil {
			_ = in.Close()
			return errors.Wrapf(err, "failed to copy content for file %q", path)
		}
		// Close the input file
		if err := in.Close(); err != nil {
			return errors.Wrapf(err, "failed to close file %q", path)
		}

		// Add a blank line for spacing between files
		if _, err := fmt.Fprintln(out); err != nil {
			return errors.Wrapf(err, "failed to add newline after file %q", path)
		}

		return nil
	})
	if err != nil {
		return errors.Wrapf(err, "failed to walk directory %q", sourceDir)
	}

	return nil
}
