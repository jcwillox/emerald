package emerald

import (
	"os"
	"path/filepath"
	"strings"
)

type FileColorsT struct {
	Base        string
	NonExistent string

	Temp       string
	Immediate  string
	Image      string
	Video      string
	Music      string
	Lossless   string
	Crypto     string
	Document   string
	Compressed string
	Compiled   string

	Directory  string
	Symlink    string
	Pipe       string
	Device     string
	Socket     string
	Irregular  string
	Executable string
	File       string
}

var FileColors = FileColorsT{
	Base:        Cyan,
	NonExistent: ColorCode("red+u"),

	Temp:       ColorIndexFg(244),
	Immediate:  ColorCode("yellow+bu"),
	Image:      ColorIndexFg(133),
	Video:      ColorIndexFg(135),
	Music:      ColorIndexFg(92),
	Lossless:   ColorIndexFg(93),
	Crypto:     ColorIndexFg(109),
	Document:   ColorIndexFg(105),
	Compressed: Red,
	Compiled:   ColorIndexFg(137),

	Directory:  ColorCode("blue+b"),
	Symlink:    ColorCode("cyan+b"),
	Pipe:       Yellow,
	Device:     ColorCode("yellow+b"),
	Socket:     ColorCode("red+b"),
	Irregular:  Yellow,
	Executable: ColorCode("green+b"),
	File:       White,
}

var (
	ImmediateNames = []string{
		"Makefile", "Cargo.toml", "SConstruct", "CMakeLists.txt",
		"build.gradle", "pom.xml", "Rakefile", "package.json", "Gruntfile.js",
		"Gruntfile.coffee", "BUILD", "BUILD.bazel", "WORKSPACE", "build.xml", "Podfile",
		"webpack.config.js", "meson.build", "composer.json", "RoboFile.php", "PKGBUILD",
		"Justfile", "Procfile", "Dockerfile", "Containerfile", "Vagrantfile", "Brewfile",
		"Gemfile", "Pipfile", "build.sbt", "mix.exs", "bsconfig.json", "tsconfig.json",
	}
	ImageExt = []string{
		"png", "jfi", "jfif", "jif", "jpe", "jpeg", "jpg", "gif", "bmp",
		"tiff", "tif", "ppm", "pgm", "pbm", "pnm", "webp", "raw", "arw",
		"svg", "stl", "eps", "dvi", "ps", "cbr", "jpf", "cbz", "xpm",
		"ico", "cr2", "orf", "nef", "heif",
	}
	VideoExt = []string{
		"avi", "flv", "m2v", "m4v", "mkv", "mov", "mp4", "mpeg",
		"mpg", "ogm", "ogv", "vob", "wmv", "webm", "m2ts", "heic",
	}
	MusicExt = []string{
		"aac", "m4a", "mp3", "ogg", "wma", "mka", "opus",
	}
	LosslessExt = []string{
		"alac", "ape", "flac", "wav",
	}
	CryptoExt = []string{
		"asc", "enc", "gpg", "pgp", "sig", "signature", "pfx", "p12",
	}
	DocumentExt = []string{
		"djvu", "doc", "docx", "dvi", "eml", "eps", "fotd", "key",
		"keynote", "numbers", "odp", "odt", "pages", "pdf", "ppt",
		"pptx", "rtf", "xls", "xlsx",
	}
	CompressedExt = []string{
		"zip", "tar", "Z", "z", "gz", "bz2", "a", "ar", "7z",
		"iso", "dmg", "tc", "rar", "par", "tgz", "xz", "txz",
		"lz", "tlz", "lzma", "deb", "rpm", "zst",
	}
	TempExt = []string{
		"tmp", "swp", "swo", "swn", "bak", "bkp", "bk",
	}
	CompiledExt = []string{
		"class", "elc", "hi", "o", "pyc", "zwc", "ko",
	}
)

func ExtensionIsOneOf(ext string, extensions []string) bool {
	for _, extension := range extensions {
		if extension == ext {
			return true
		}
	}
	return false
}

func GetExtOnly(filename string) string {
	return strings.TrimPrefix(filepath.Ext(filename), ".")
}

func IsTemp(filename string) bool {
	if strings.HasSuffix(filename, "~") {
		return true
	}
	if strings.HasPrefix(filename, "#") && strings.HasSuffix(filename, "#") {
		return true
	}
	return ExtensionIsOneOf(GetExtOnly(filename), TempExt)
}

func IsImmediate(filename string) bool {
	if strings.HasPrefix(strings.ToLower(filename), "readme") {
		return true
	} else if strings.HasSuffix(filename, ".ninja") {
		return true
	}
	for _, name := range ImmediateNames {
		if filename == name {
			return true
		}
	}
	return false
}

func IsImage(ext string) bool {
	return ExtensionIsOneOf(ext, ImageExt)
}

func IsVideo(ext string) bool {
	return ExtensionIsOneOf(ext, VideoExt)
}

func IsMusic(ext string) bool {
	return ExtensionIsOneOf(ext, MusicExt)
}

func IsLossless(ext string) bool {
	return ExtensionIsOneOf(ext, LosslessExt)
}

func IsCrypto(ext string) bool {
	return ExtensionIsOneOf(ext, CryptoExt)
}

func IsDocument(ext string) bool {
	return ExtensionIsOneOf(ext, DocumentExt)
}

func IsCompressed(ext string) bool {
	return ExtensionIsOneOf(ext, CompressedExt)
}

func IsCompiled(ext string) bool {
	return ExtensionIsOneOf(ext, CompiledExt)
}

func HighlightFile(filename string, mode ...os.FileMode) string {
	return GetFileColor(filename, mode...) + filename + Reset
}

func HighlightFileStat(filename string, stat ...os.FileInfo) string {
	if len(stat) == 0 {
		info, _ := os.Lstat(filename)
		return GetFileColorStat(info) + filepath.Base(filename) + Reset
	}
	return GetFileColorStat(stat[0]) + filename + Reset
}

func HighlightPath(path string, mode ...os.FileMode) string {
	if plain {
		return path
	}
	dir, file := filepath.Split(path)
	return FileColors.Base + dir + HighlightFile(file, mode...)
}

func HighlightPathStat(path string, stat ...os.FileInfo) string {
	if len(stat) == 0 {
		info, _ := os.Lstat(path)
		stat = []os.FileInfo{info}
	}
	if plain {
		return path
	}
	if stat[0] == nil {
		return FileColors.NonExistent + path + Reset
	}
	dir, file := filepath.Split(path)
	return FileColors.Base + dir + HighlightFile(file, stat[0].Mode())
}

func GetFileColor(filename string, mode ...os.FileMode) string {
	if len(mode) > 0 {
		color := GetFileModeColor(mode[0])
		if color != "" {
			return color
		}
	}
	color := GetFileTypeColor(filename)
	if color != "" {
		return color
	}
	return Reset
}

func GetFileColorStat(stat os.FileInfo) string {
	if plain {
		return ""
	}
	if stat == nil {
		return FileColors.NonExistent
	}
	color := GetFileModeColor(stat.Mode())
	if color != "" {
		return color
	}
	color = GetFileTypeColor(stat.Name())
	if color != "" {
		return color
	}
	return Reset
}

func GetFileTypeColor(filename string) string {
	if plain {
		return ""
	}
	ext := GetExtOnly(filename)
	switch {
	case IsTemp(filename):
		return FileColors.Temp
	case IsImmediate(filename):
		return FileColors.Immediate
	case IsImage(ext):
		return FileColors.Image
	case IsVideo(ext):
		return FileColors.Video
	case IsMusic(ext):
		return FileColors.Music
	case IsLossless(ext):
		return FileColors.Lossless
	case IsCrypto(ext):
		return FileColors.Crypto
	case IsDocument(ext):
		return FileColors.Document
	case IsCompressed(ext):
		return FileColors.Compressed
	case IsCompiled(ext):
		return FileColors.Compiled
	}
	return ""
}

func GetFileModeColor(mode os.FileMode) string {
	if plain {
		return ""
	}
	if mode&os.ModeDir != 0 {
		return FileColors.Directory
	} else if mode&os.ModeSymlink != 0 {
		return FileColors.Symlink
	} else if mode&os.ModeNamedPipe != 0 {
		return FileColors.Pipe
	} else if mode&os.ModeDevice != 0 {
		return FileColors.Device
	} else if mode&os.ModeCharDevice != 0 {
		return FileColors.Device
	} else if mode&os.ModeSocket != 0 {
		return FileColors.Socket
	} else if mode&os.ModeIrregular != 0 {
		return FileColors.Irregular
	} else if mode&0111 != 0 {
		return FileColors.Executable
	}
	return ""
}

func boldAnsi(ansi string) string {
	if strings.HasPrefix(ansi, "\x1b[0;") && len(ansi) > 4 && ansi[4] != '1' {
		return "\x1b[0;1;" + ansi[4:]
	}
	return ansi
}

func HighlightFileMode(mode os.FileMode) string {
	perms := mode.String()
	sb := strings.Builder{}
	sb.Grow(128)

	color := GetFileModeColor(mode)
	if color == FileColors.Executable {
		color = ""
	}
	sb.WriteString(boldAnsi(color))
	if perms[0] == '-' {
		sb.WriteString(".")
	} else {
		sb.WriteByte(perms[0])
	}
	for i := 1; i < len(perms); i++ {
		if perms[i] == '-' {
			if i == 1 || perms[i-1] != '-' {
				sb.WriteString(LightBlack)
			}
			sb.WriteByte(perms[i])
		} else {
			if i < 4 {
				sb.WriteString(Bold)
			}
			switch perms[i] {
			case 'r':
				sb.WriteString(Yellow)
			case 'w':
				sb.WriteString(Red)
			case 'x':
				if i == 3 && mode.IsRegular() && !plain {
					sb.WriteString(start + "4;32m")
				}
				sb.WriteString(Green)
			}
			sb.WriteByte(perms[i])
			sb.WriteString(Reset)
		}
	}
	sb.WriteString(Reset)
	return sb.String()
}
