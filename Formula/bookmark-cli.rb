class BookmarkCli < Formula
  desc "Command Line tool to manage all of your bookmarks at one place"
  homepage "https://jogendra.github.io/building-command-line-tools-in-go"
  url "https://github.com/jogendra/bookmark-cli/archive/2.0.0.tar.gz"
  sha256 "78d9434e921ffdece53dbb53197a3b63c8afc4b709dd502b29fffaf820b0ff0c"
  license "Apache-2.0"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    bookmark_path = buildpath/"src/github.com/jogendra/bookmark-cli"
    bookmark_path.install buildpath.children

    cd bookmark_path do
      system "go", "build"
      bin.install "bookmark"
    end
  end

  test do
    assert_match "2.0.0", shell_output("#{bin}/bookmark --version ")
  end
end
