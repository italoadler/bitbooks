default:
	@tmux_send "make build"

pandoc := pandoc \
	00_intro.md \
	01_triangles.md \
	02_pythag.md \
	03_sine_cosine.md \
	-V margin-left=1in \
	-V margin-right=1in \
	-V margin-top=1in \
	-V margin-bottom=1in \
	--smart \
	--ascii \
	-fmarkdown-implicit_figures \
	--toc \
	-o 

build: clean out/coding-trigonometry.pdf out/coding-trigonometry.epub out/coding-trigonometry.mobi out/coding-trigonometry.azw3
	@echo "\ncomplete. files in out folder:"
	@ls -la out/coding-trigonometry.*

clean:
	@echo "cleaning..."
	@rm -rf out
	@mkdir out

out/coding-trigonometry.epub:
	@echo "\ngenerating epub..."
	@$(pandoc) out/coding-trigonometry.epub

out/coding-trigonometry.pdf:
	@echo "\ngenerating pdf..."
	@$(pandoc) out/coding-trigonometry.pdf

out/coding-trigonometry.mobi: out/coding-trigonometry.epub
	@echo "\ngenerating mobi..."
	@ebook-convert out/coding-trigonometry.epub out/coding-trigonometry.mobi

out/coding-trigonometry.azw3: out/coding-trigonometry.epub
	@echo "\ngenerating azw3..."
	@ebook-convert out/coding-trigonometry.epub out/coding-trigonometry.azw3

