.PHONY: help init-kata

help:
	@echo "Katas Training - Available commands:"
	@echo ""
	@echo "  make init-kata KATA=<name>  Initialize a new kata version"
	@echo ""
	@echo "Available katas:"
	@{ \
	ls -1 .katas/*.tar.gz 2>/dev/null | xargs -n1 basename | sed 's/.tar.gz$$//'; \
	ls -1d .katas/*/ 2>/dev/null | xargs -n1 basename; \
	} | sort -u | sed 's/^/  - /' || echo "  (none)"
	@echo ""
	@echo "Example:"
	@echo "  make init-kata KATA=movie_rental"

init-kata:
	@./.katas/init-kata.sh $(KATA)
