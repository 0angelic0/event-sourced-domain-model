SHELL := /usr/bin/env bash
ROOT := ${CURDIR}

.PHONY: help
help:
	@echo 'Usage:'
	@echo '    make umlgen               	Generate the uml class diagram from code.'
	@echo '    make run0                 	Run the legacy program.'
	@echo '    make run1                 	Run the transactionscript program.'
	@echo '    make run2                 	Run the activerecord program.'
	@echo '    make run3                 	Run the domainmodel program.'
	@echo '    make run4                 	Run the eventsourceddomainmodel program.'
	@echo

.PHONY: umlgen
umlgen:
	goplantuml -recursive -title legacy ./services/legacy > diagrams/diagram_legacy.puml
	goplantuml -recursive -title transactionscript ./services/transactionscript > diagrams/diagram_transactionscript.puml
	goplantuml -recursive -title activerecord ./services/activerecord > diagrams/diagram_activerecord.puml
	goplantuml -recursive -title domainmodel ./services/domainmodel > diagrams/diagram_domainmodel.puml
	goplantuml -recursive -title eventsourceddomainmodel ./services/eventsourceddomainmodel > diagrams/diagram_eventsourceddomainmodel.puml
# add this line after @startuml
# !pragma layout smetana

.PHONY: run0
run0:
	go run ./services/legacy/.

.PHONY: run1
run1:
	go run ./services/transactionscript/.

.PHONY: run2
run2:
	go run ./services/activerecord/.

.PHONY: run3
run3:
	go run ./services/domainmodel/.

.PHONY: run4
run4:
	go run ./services/eventsourceddomainmodel/.