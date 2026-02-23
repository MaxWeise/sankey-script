package internal

import (
	"context"
	"testing"

	_ "modernc.org/sqlite"
)

// Create two objects and check that the ids are different
func TestCreateEntryUniqueId(t *testing.T) {
	// Arrange
	underTest, err := setupTestService()
	if err != nil {
		t.Errorf("Setting up the test raised an error: %s", err)
	}

	// Run Tests
	object1, err := underTest.CreateEntry("source1", "target1", 10.0, "")
	if err != nil {
		t.Error("Error while creating first object.")
	}
	object2, err := underTest.CreateEntry("source2", "target2", 10.0, "")
	if err != nil {
		t.Error("Error while creating second object.")
	}

	// Asserts
	if object1 == object2 {
		t.Errorf("The crated ids are not unique: %d == %d", object1.ID, object2.ID)
	}
}

func TestCreateEntry(t *testing.T) {
	// Arrange
	underTest, err := setupTestService()
	if err != nil {
		t.Errorf("Setting up the test raised an error: %s", err)
	}

	// Run tests
	actual, err := underTest.CreateEntry("Gehalt", "Einnahmen", 100, "")

	// Asserts
	if err != nil {
		t.Errorf("An error occured in testing: %s", err)
	}

	if actual.Source != "Gehalt" {
		t.Errorf("The object does not the correct value for Source.\nExpected: Gehalt\nActual: %s", actual.Source)
	}

	if actual.Target != "Einnahmen" {
		t.Errorf("The object does not the correct value for Source.\nExpected: Einnahmen\nActual: %s", actual.Source)
	}

	// NOTE: Make sure this is float32
	if actual.Amount != 100.0 {
		t.Errorf("The object does not the correct value for Amount.\nExpected: 100\nActual: %d", actual.Amount)
	}

	if actual.Description.Valid == true || actual.Description.String != "" {
		t.Errorf("The object has a value which should be nil: Discription %s", actual.Description.String)
	}
}

func TestCreateEntryWithDescription(t *testing.T) {
	// Arrange
	underTest, err := setupTestService()
	if err != nil {
		t.Errorf("Setting up the test raised an error: %s", err)
	}

	// Run tests
	actual, err := underTest.CreateEntry("Gehalt", "Einnahmen", 100, "Test Description")

	// Asserts
	if err != nil {
		t.Errorf("An error occured in testing: %s", err)
	}

	if actual.Source != "Gehalt" {
		t.Errorf("The object does not the correct value for Source.\nExpected: Gehalt\nActual: %s", actual.Source)
	}

	if actual.Target != "Einnahmen" {
		t.Errorf("The object does not the correct value for Source.\nExpected: Einnahmen\nActual: %s", actual.Source)
	}

	// NOTE: Make sure this is float32
	if actual.Amount != 100.0 {
		t.Errorf("The object does not the correct value for Amount.\nExpected: 100\nActual: %d", actual.Amount)
	}

	if actual.Description.Valid == false || actual.Description.String == "" {
		t.Error("The description is missing")
	}

	if actual.Description.String != "Test Description" {
		t.Errorf("The object does not the correct value for Description.\nExpected: Test Description\nActual: %s", actual.Description.String)
	}
}

func setupTestService() (CoreService, error) {
	// Arrange
	ctx := context.Background()
	q, err := GetQueries(ctx, ":memory:")

	if err != nil {
		return CoreService{}, err
	}

	return CoreService{q, ctx}, nil
}
