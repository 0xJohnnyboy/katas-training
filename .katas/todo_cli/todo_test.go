package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestAddAndList(t *testing.T) {
	t.Parallel()
	csvPath := filepath.Join(t.TempDir(), "todos.csv")

	_, err := run([]string{"add", "2026-03-01", "PayRent", "send bank transfer"}, csvPath)
	if err != nil {
		t.Fatalf("add failed: %v", err)
	}

	out, err := run([]string{"list"}, csvPath)
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}
	if !strings.Contains(out, "1 | [ ] | 2026-03-01 | PayRent | send bank transfer") {
		t.Fatalf("unexpected list output: %s", out)
	}
}

func TestUpdateAndDelete(t *testing.T) {
	t.Parallel()
	csvPath := filepath.Join(t.TempDir(), "todos.csv")

	_, _ = run([]string{"add", "2026-03-01", "Task1", "desc1"}, csvPath)
	_, _ = run([]string{"add", "2026-03-02", "Task2", "desc2"}, csvPath)

	_, err := run([]string{"update", "2", "2026-04-10", "Task2b", "new desc", "true"}, csvPath)
	if err != nil {
		t.Fatalf("update failed: %v", err)
	}

	out, err := run([]string{"list"}, csvPath)
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}
	if !strings.Contains(out, "2 | [x] | 2026-04-10 | Task2b | new desc") {
		t.Fatalf("updated todo not found, got: %s", out)
	}

	_, err = run([]string{"delete", "1"}, csvPath)
	if err != nil {
		t.Fatalf("delete failed: %v", err)
	}

	out, err = run([]string{"list"}, csvPath)
	if err != nil {
		t.Fatalf("list after delete failed: %v", err)
	}
	if strings.Contains(out, "Task1") {
		t.Fatalf("deleted todo still present, got: %s", out)
	}
}

func TestAddRejectsInvalidDate(t *testing.T) {
	t.Parallel()
	csvPath := filepath.Join(t.TempDir(), "todos.csv")

	_, err := run([]string{"add", "03-01-2026", "Task1", "desc1"}, csvPath)
	if err == nil {
		t.Fatal("expected error for invalid date")
	}
}

func TestListShowsOnlyLastTwenty(t *testing.T) {
	t.Parallel()
	csvPath := filepath.Join(t.TempDir(), "todos.csv")

	for i := 1; i <= 25; i++ {
		_, err := run([]string{"add", "2026-03-01", "Task", "item"}, csvPath)
		if err != nil {
			t.Fatalf("add failed at %d: %v", i, err)
		}
	}

	out, err := run([]string{"list"}, csvPath)
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 20 {
		t.Fatalf("expected 20 lines, got %d: %s", len(lines), out)
	}
	if !strings.HasPrefix(lines[0], "6 | ") {
		t.Fatalf("expected first visible id to be 6, got: %s", lines[0])
	}
	if !strings.HasPrefix(lines[len(lines)-1], "25 | ") {
		t.Fatalf("expected ids 6..25 to be shown, got: %s", out)
	}
}
