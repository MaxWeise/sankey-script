package internal

import (
	"context"
	"database/sql"
	"maxweise/sankey-script/data_acess"
	"testing"

	_ "modernc.org/sqlite"
)

// Create two objects and check that the ids are different
func TestCreateEntryUniqueId(t *testing.T) {
	// Arrange
	underTest, err := setupTestService()
	if err != nil {
		t.Errorf("Setting up the test raised an error: %s", err)
		return
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
		return
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
		return
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

// Test that all entries get retrieved from the database.
func TestReadEntries(t *testing.T) {
	// arrange
	underTest, err := setupTestService()
	if err != nil {
		t.Errorf("Setting up the test raised an error: %s", err)
		return
	}
	_, _ = underTest.CreateEntry("Gehalt", "Einnahmen", 100, "Test Description")

	// Run test
	actual, err := underTest.ReadAllEntries()
	if err != nil {
		t.Errorf("Someting went wrong: %s", err)
	}

	// Assert
	if len(actual) == 0 {
		t.Error("The read function did not return any values.")
	}
}

// All attributes should handle change values correctly.
func TestChangeEntry(t *testing.T) {
	// arrange
	underTest, err := setupTestService()
	if err != nil {
		t.Errorf("Setting up the test raised an error: %s", err)
		return
	}
	o, _ := underTest.CreateEntry("Gehalt", "Einnahmen", 100, "Test Description")

	// Run test
	actual, err := underTest.ChangeEntry(
		o.ID,
		"TestString",
		"",
		sql.NullFloat64{Valid: false},
		sql.NullString{Valid: true, String: "Eine neue Beschreibung"},
	)

	// Assert
	if err != nil {
		t.Errorf("There has been an error while testing: %s", err)
	}

	expected := data_acess.Entry{
		ID:     o.ID,
		Source: "TestString",
		Target: "Einnahmen",
		Amount: float64(100),
		Description: sql.NullString{
			String: "Eine neue Beschreibung",
			Valid:  true,
		},
	}

	if actual != expected {
		t.Errorf("\nExpected: \t%#v\nGot: \t\t%#v", expected, actual)
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
