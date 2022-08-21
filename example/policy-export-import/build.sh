POLICY_MIGRATION=/home/mj/Projects/mjlab/gobdgz/example/policy-export-import/main.go
DIST=/home/mj/Projects/mjlab/gobdgz/dist

GOOS=linux GOARCH=386 go build -o $DIST/linux/x86/policy-migration $POLICY_MIGRATION
GOOS=linux GOARCH=amd64 go build -o $DIST/linux/x64/policy-migration $POLICY_MIGRATION

GOOS=windows GOARCH=386 go build -o $DIST/windows/x86/policy-migration.exe $POLICY_MIGRATION
GOOS=windows GOARCH=amd64 go build -o $DIST/windows/x64/policy-migration.exe $POLICY_MIGRATION

# GOOS=darwin GOARCH=386 go build -o $DIST/macOS/x86/policy-migration $POLICY_MIGRATION
GOOS=darwin GOARCH=amd64 go build -o $DIST/macOS/x64/policy-migration $POLICY_MIGRATION

