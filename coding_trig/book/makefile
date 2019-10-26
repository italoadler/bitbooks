default:
	@tmux_send "make build"

pandoc := pandoc \
	index.md \
	chapter_01/chapter_01.md \
	chapter_02/chapter_02.md \
	chapter_03/chapter_03.md \
	chapter_04/chapter_04.md \
	chapter_05/chapter_05.md \
	chapter_06/chapter_06.md \
	-V margin-left=1in \
	-V margin-right=1in \
	-V margin-top=1in \
	-V margin-bottom=1in \
	--ascii \
	-fmarkdown-implicit_figures \
	--toc \
	--pdf-engine=xelatex \
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

