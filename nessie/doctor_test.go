package nessie

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestDoctor(t *testing.T) {
	accessKeyID := viper.GetString("access_key_id")
	secretAccessKey := viper.GetString("secret_access_key")
	endPointURL := viper.GetString("endpoint_url") + "/api/v1"
	emptyConfigFileName := "empty.yaml"
	wrongFormatFileName := "wrong_format.yaml"
	os.Create(emptyConfigFileName)
	ioutil.WriteFile(wrongFormatFileName, []byte("nothing"), 0644)
	RunCmdAndVerifySuccessWithFile(t, LakectlWithParams(accessKeyID, secretAccessKey, endPointURL)+" doctor", false, "lakectl_doctor_ok", emptyVars)
	RunCmdAndVerifyFailureWithFile(t, lakectlLocation()+" doctor -c "+wrongFormatFileName, false, "lakectl_doctor_wrong_format_file", emptyVars)
	RunCmdAndVerifyFailureWithFile(t, lakectlLocation()+" doctor -c "+emptyConfigFileName, false, "lakectl_doctor_empty_file", emptyVars)
	RunCmdAndVerifyFailureWithFile(t, lakectlLocation()+" doctor -c not_exited.yaml", false, "lakectl_doctor_not_existed_file", emptyVars)
	RunCmdAndVerifyFailureWithFile(t, LakectlWithParams(accessKeyID, secretAccessKey, endPointURL+"1")+" doctor", false, "lakectl_doctor_wrong_endpoint", emptyVars)
	RunCmdAndVerifyFailureWithFile(t, LakectlWithParams(accessKeyID, secretAccessKey, "wrong_uri")+" doctor", false, "lakectl_doctor_wrong_uri_format_endpoint", emptyVars)
	RunCmdAndVerifyFailureWithFile(t, LakectlWithParams("AKIAJZZZZZZZZZZZZZZQ", secretAccessKey, endPointURL)+" doctor", false, "lakectl_doctor_wrong_access_key_id", emptyVars)
	RunCmdAndVerifyFailureWithFile(t, LakectlWithParams("AKIAJOIhCOZ5JBYHCSDQ", secretAccessKey, endPointURL)+" doctor", false, "lakectl_doctor_with_lower_case_access_key_id", emptyVars)
	RunCmdAndVerifyFailureWithFile(t, LakectlWithParams("AKIBJOIHCOZ5JBYHCSDQ", secretAccessKey, endPointURL)+" doctor", false, "lakectl_doctor_with_wrong_prefix_access_key_id", emptyVars)
	RunCmdAndVerifyFailureWithFile(t, LakectlWithParams("AKIAJOIHCOZ5JBYHCSDB", secretAccessKey, endPointURL)+" doctor", false, "lakectl_doctor_with_wrong_sufix_access_key_id", emptyVars)
	RunCmdAndVerifyFailureWithFile(t, LakectlWithParams("AKIAJOICO!Z5JBYHCSDQ", secretAccessKey, endPointURL)+" doctor", false, "lakectl_doctor_with_invalid_char_access_key_id", emptyVars)
	RunCmdAndVerifyFailureWithFile(t, LakectlWithParams("AKIAJOICOZ5JBYHCSDQ", secretAccessKey, endPointURL)+" doctor", false, "lakectl_doctor_with_short_length_access_key_id", emptyVars)
	RunCmdAndVerifyFailureWithFile(t, LakectlWithParams("AKIAJOIHHCOZ5JBYHCSDQ", secretAccessKey, endPointURL)+" doctor", false, "lakectl_doctor_with_long_length_access_key_id", emptyVars)
	RunCmdAndVerifyFailureWithFile(t, LakectlWithParams(accessKeyID, "zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz", endPointURL)+" doctor", false, "lakectl_doctor_wrong_secret_access_key", emptyVars)
	RunCmdAndVerifyFailureWithFile(t, LakectlWithParams(accessKeyID, "TQG5JcovOozCGJnIRmIKH7Flq1tLxnuByi9/WmJ", endPointURL)+" doctor", false, "lakectl_doctor_short_secret_access_key", emptyVars)
	RunCmdAndVerifyFailureWithFile(t, LakectlWithParams(accessKeyID, "TQG5JcovOozCGJnIRmIKH7Flq1tLxnuByi9/WmJzz", endPointURL)+" doctor", false, "lakectl_doctor_long_secret_access_key", emptyVars)
	RunCmdAndVerifyFailureWithFile(t, LakectlWithParams(accessKeyID, "TQG5JcovOozCGJnIRmIKH7Flq1tLxnuByi9/WmJ!", endPointURL)+" doctor", false, "lakectl_doctor_invalid_char_secret_access_key", emptyVars)
}
