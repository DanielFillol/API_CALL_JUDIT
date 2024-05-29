package csv

import (
	"API_CALL_JUDIT/models"
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// WriteLawsuits writes a CSV file with the given file name and folder name, and the data from the responses.
func WriteLawsuits(fileName string, folderName string, responses []models.ResponseToCSV) error {
	// Create a slice to hold all the rows for the CSV file
	var rows [][]string

	// Add the headers to the slice
	rows = append(rows, generateHeaders())

	// Add the data rows to the slice
	for _, response := range responses {
		rows = append(rows, generateRow(response)...)
	}

	// Create the CSV file
	cf, err := createFile(folderName + "/" + fileName + ".csv")
	if err != nil {
		log.Println(err)
		return err
	}

	// Close the file when the function completes
	defer cf.Close()

	// Create a new CSV writer
	w := csv.NewWriter(cf)

	// WriteLawsuits all the rows to the CSV file
	err = w.WriteAll(rows)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// createFile function takes in a file path and creates a file in the specified directory. It returns a pointer to the created file and an error if there is any.
func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		log.Println(err)
		return nil, err
	}
	return os.Create(p)
}

// generateHeaders function returns a slice of strings containing the header values for the CSV file.
func generateHeaders() []string {
	return []string{
		"Document",
		"Total Time",
		"Code",
		"Justice",
		"Tribunal",
		"Instance",
		"DistributionDate",
		"TribunalAcronym",
		"SecrecyLevel",
		"label_IsFallbackSource",
		"label_CrawlId",
		"label_DictionaryUpdatedAt",
		"label_Criminal",
		"subjects_Code",
		"subjects_Name",
		"subjects_Date",
		"all_subjects",
		"classifications_Code",
		"classifications_Name",
		"classifications_Date",
		"all_classifications",
		"courts_Code",
		"courts_Name",
		"courts_Date",
		"all_courts",
		"parties_Name",
		"parties_Side",
		"parties_PersonType",
		"parties_MainDocument",
		"parties_EntityType",
		"all_parties",
		"steps_LawsuitCnj",
		"steps_LawsuitInstance",
		"steps_StepId",
		"steps_StepDate",
		"steps_Content",
		"steps_Private",
		"steps_Tags_CrawlId",
		"steps_StepsCount",
		"all_steps",
		"Attachments",
		"RelatedLawsuits",
		"crawler_SourceName",
		"crawler_CrawlId",
		"crawler_Weight",
		"crawler_UpdatedAt",
		"Amount",
		"lastStep_LawsuitCnj",
		"lastStep_LawsuitInstance",
		"lastStep_StepId",
		"lastStep_StepDate",
		"lastStep_Private",
		"lastStep_StepsCount",
		"lastStep_Content",
		"lastStep_Tags_CrawlId",
		"Phase",
		"Status",
		"Name",
		"Judge",
		"FreeJustice",
	}
}

// generateRow function takes in a single models.WriteStruct argument and returns a slice of strings containing the values to be written in a row of the CSV file.
func generateRow(response models.ResponseToCSV) [][]string {
	var rows [][]string
	if len(response.R.PageData) == 0 {
		// Append Identification
		row := []string{
			response.Document,
		}
		rows = append(rows, row)
	} else {
		for _, pageData := range response.R.PageData {
			row := []string{
				response.Document,
				response.TotalTime.String(),
			}
			row = append(row)
			row = append(row, pageData.ResponseData.Code)
			row = append(row, pageData.ResponseData.Justice)
			row = append(row, pageData.ResponseData.Tribunal)
			row = append(row, strconv.Itoa(pageData.ResponseData.Instance))
			row = append(row, pageData.ResponseData.DistributionDate.String())
			row = append(row, pageData.ResponseData.TribunalAcronym)
			row = append(row, strconv.Itoa(pageData.ResponseData.SecrecyLevel))
			row = append(row, strconv.FormatBool(pageData.ResponseData.Labels.IsFallbackSource))
			row = append(row, pageData.ResponseData.Labels.CrawlId)
			row = append(row, pageData.ResponseData.Labels.DictionaryUpdatedAt.String())
			row = append(row, strconv.FormatBool(pageData.ResponseData.Labels.Criminal))

			// Append specific subject
			if len(pageData.ResponseData.Subjects) != 0 {
				row = append(row, pageData.ResponseData.Subjects[0].Code)
				row = append(row, pageData.ResponseData.Subjects[0].Name)
				row = append(row, pageData.ResponseData.Subjects[0].Date.String())
				if len(pageData.ResponseData.Subjects) > 1 {
					var subjects string
					for _, subject := range pageData.ResponseData.Subjects {
						subjects += "|code: " + subject.Code + " name: " + subject.Name + " data: " + subject.Date.String() + "|"
					}
					row = append(row, subjects)
				} else {
					row = append(row, "1")
				}
			} else {
				row = append(row, "")
				row = append(row, "")
				row = append(row, "")
				row = append(row, "")
			}

			// Append specific Classification
			if len(pageData.ResponseData.Classifications) != 0 {
				row = append(row, pageData.ResponseData.Classifications[0].Code)
				row = append(row, pageData.ResponseData.Classifications[0].Name)
				row = append(row, pageData.ResponseData.Classifications[0].Date.String())
				if len(pageData.ResponseData.Classifications) > 1 {
					var classifications string
					for _, classification := range pageData.ResponseData.Classifications {
						classifications += "|code: " + classification.Code + " name: " + classification.Name + " data: " + classification.Date.String() + "|"
					}
					row = append(row, classifications)
				} else {
					row = append(row, "1")
				}
			} else {
				row = append(row, "")
				row = append(row, "")
				row = append(row, "")
				row = append(row, "")
			}

			// Append specific Court
			if len(pageData.ResponseData.Courts) != 0 {
				row = append(row, pageData.ResponseData.Courts[0].Code)
				row = append(row, pageData.ResponseData.Courts[0].Name)
				row = append(row, pageData.ResponseData.Courts[0].Date.String())
				if len(pageData.ResponseData.Courts) > 1 {
					var courts string
					for _, court := range pageData.ResponseData.Courts {
						courts += "|code: " + court.Code + " name: " + court.Name + " data: " + court.Date.String() + "|"
					}
					row = append(row, courts)
				} else {
					row = append(row, "1")
				}
			} else {
				row = append(row, "")
				row = append(row, "")
				row = append(row, "")
				row = append(row, "")
			}

			// Append specific Part
			if len(pageData.ResponseData.Parties) != 0 {
				row = append(row, pageData.ResponseData.Parties[0].Name)
				row = append(row, pageData.ResponseData.Parties[0].Side)
				row = append(row, pageData.ResponseData.Parties[0].PersonType)
				row = append(row, pageData.ResponseData.Parties[0].MainDocument)
				row = append(row, pageData.ResponseData.Parties[0].EntityType)
				if len(pageData.ResponseData.Parties) > 1 {
					var parties string
					for _, partie := range pageData.ResponseData.Parties {
						parties += "|nomeParte: " + partie.Name + " ladoParte: " + partie.Side + " tipoParte: " + partie.PersonType + " documentoPrincipal: " + partie.MainDocument + " entidadeTipo: " + partie.EntityType + "|"
					}
					row = append(row, parties)
				} else {
					row = append(row, "1")
				}
			} else {
				row = append(row, "")
				row = append(row, "")
				row = append(row, "")
				row = append(row, "")
				row = append(row, "")
				row = append(row, "")
			}

			// Append specific step
			if len(pageData.ResponseData.Steps) != 0 {
				row = append(row, pageData.ResponseData.Steps[0].LawsuitCnj)
				row = append(row, strconv.Itoa(pageData.ResponseData.Steps[0].LawsuitInstance))
				row = append(row, pageData.ResponseData.Steps[0].StepId)
				row = append(row, pageData.ResponseData.Steps[0].StepDate.String())
				row = append(row, pageData.ResponseData.Steps[0].Content)
				row = append(row, strconv.FormatBool(pageData.ResponseData.Steps[0].Private))
				row = append(row, pageData.ResponseData.Steps[0].Tags.CrawlId)
				row = append(row, strconv.Itoa(pageData.ResponseData.Steps[0].StepsCount))
				if len(pageData.ResponseData.Steps) > 1 {
					var steps string
					for _, step := range pageData.ResponseData.Steps {
						steps += "|LawsuitCnj: " + step.LawsuitCnj + " LawsuitInstance: " + strconv.Itoa(step.LawsuitInstance) + " StepId: " + step.StepId + " StepDate: " + step.StepDate.String() + " Content: " + step.Content + " Private: " + strconv.FormatBool(pageData.ResponseData.Steps[0].Private) + " CrawledId: " + pageData.ResponseData.Steps[0].Tags.CrawlId + " StepsCount: " + strconv.Itoa(pageData.ResponseData.Steps[0].StepsCount) + "|"
					}
					row = append(row, steps)
				} else {
					row = append(row, "1")
				}
			} else {
				row = append(row, "")
				row = append(row, "")
				row = append(row, "")
				row = append(row, "")
				row = append(row, "")
				row = append(row, "")
				row = append(row, "")
				row = append(row, "")
				row = append(row, "")
			}
			row = append(row, "")
			row = append(row, "")
			row = append(row, pageData.ResponseData.Crawler.SourceName)
			row = append(row, pageData.ResponseData.Crawler.CrawlId)
			row = append(row, strconv.Itoa(pageData.ResponseData.Crawler.Weight))
			row = append(row, pageData.ResponseData.Crawler.UpdatedAt.String())
			row = append(row, strconv.FormatFloat(pageData.ResponseData.Amount, 'f', 6, 64)) //*
			row = append(row, pageData.ResponseData.LastStep.LawsuitCnj)
			row = append(row, strconv.Itoa(pageData.ResponseData.LastStep.LawsuitInstance))
			row = append(row, pageData.ResponseData.LastStep.StepId)
			row = append(row, pageData.ResponseData.LastStep.StepDate.String())
			row = append(row, strconv.FormatBool(pageData.ResponseData.LastStep.Private))
			row = append(row, strconv.Itoa(pageData.ResponseData.LastStep.StepsCount))
			row = append(row, pageData.ResponseData.LastStep.Content)
			row = append(row, pageData.ResponseData.LastStep.Tags.CrawlId)
			row = append(row, pageData.ResponseData.Phase)
			row = append(row, pageData.ResponseData.Status)
			row = append(row, pageData.ResponseData.Name)
			row = append(row, "")
			row = append(row, strconv.FormatBool(pageData.ResponseData.FreeJustice))
			rows = append(rows, row)
		}
	}
	return rows
}
