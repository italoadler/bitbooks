default:
	@tmux_send "make pdf"

pandoc := pandoc \
	chapter_01.md \
	chapter_02.md \
	chapter_03.md \
	chapter_04.md \
	chapter_05.md \
	chapter_06.md \
	--ascii \
	--epub-cover-image=images/cover.png \
	--toc \
	--pdf-engine=xelatex \
	--metadata title="Coding Trigonometry" \
	-V margin-left=1in \
	-V margin-right=1in \
	-V margin-top=1in \
	-V margin-bottom=1in \
	-V mainfont="DejaVu Serif" \
	-V monofont="DejaVu Sans Mono" \
	-V mainfontoptions="ItalicFont=DejaVu Serif Italic" \
	-o 

all: ../out/coding-trigonometry.pdf ../out/coding-trigonometry.epub ../out/coding-trigonometry.mobi ../out/coding-trigonometry.azw3
	@echo "\ncomplete. files in out folder:"
	@ls -la ../out/coding-trigonometry.*

epub: clean ../out/coding-trigonometry.epub
	ebook-viewer ../out/coding-trigonometry.epub

pdf: clean ../out/coding-trigonometry.pdf
	xdg-open ../out/coding-trigonometry.pdf

mobi: clean ../out/coding-trigonometry.mobi
	ebook-viewer ../out/coding-trigonometry.mobi

azw3: clean ../out/coding-trigonometry.azw3
	ebook-viewer ../out/coding-trigonometry.azw3

clean:
	@echo "cleaning..."
	@rm -rf ../out
	@mkdir ../out

../out/coding-trigonometry.epub:
	@echo "generating epub..."
	@$(pandoc) ../out/coding-trigonometry.epub

../out/coding-trigonometry.pdf:
	@echo "generating pdf..."
	@$(pandoc) ../out/coding-trigonometry.pdf

../out/coding-trigonometry.mobi: ../out/coding-trigonometry.epub
	@echo "generating mobi..."
	@../tools/kindlegen ../out/coding-trigonometry.epub 

../out/coding-trigonometry.azw3: ../out/coding-trigonometry.epub
	@echo "generating azw3..."
	@ebook-convert ../out/coding-trigonometry.epub ../out/coding-trigonometry.azw3

