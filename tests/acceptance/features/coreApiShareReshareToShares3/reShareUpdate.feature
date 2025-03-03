@api @issue-1328
Feature: sharing
  As a user
  I want to update share permissions
  So that I can decide what resources can be shared with which permission

  Background:
    Given these users have been created with default attributes and without skeleton files:
      | username |
      | Alice    |
      | Brian    |
      | Carol    |
    And user "Alice" has created folder "/TMP"


  Scenario Outline: update of reshare can reduce permissions
    Given using OCS API version "<ocs_api_version>"
    And user "Alice" has shared folder "/TMP" with user "Brian" with permissions "share,create,update,read"
    And user "Brian" has accepted share "/TMP" offered by user "Alice"
    And user "Brian" has shared folder "Shares/TMP" with user "Carol" with permissions "share,create,update,read"
    And user "Carol" has accepted share "/TMP" offered by user "Brian"
    When user "Brian" updates the last share using the sharing API with
      | permissions | share,read |
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And user "Carol" should not be able to upload file "filesForUpload/textfile.txt" to "Shares/TMP/textfile.txt"
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |


  Scenario Outline: update of reshare can increase permissions to the maximum allowed
    Given using OCS API version "<ocs_api_version>"
    And user "Alice" has shared folder "/TMP" with user "Brian" with permissions "share,create,update,read"
    And user "Brian" has accepted share "/TMP" offered by user "Alice"
    And user "Brian" has shared folder "Shares/TMP" with user "Carol" with permissions "share,read"
    And user "Carol" has accepted share "/TMP" offered by user "Brian"
    When user "Brian" updates the last share using the sharing API with
      | permissions | share,create,update,read |
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And user "Carol" should be able to upload file "filesForUpload/textfile.txt" to "Shares/TMP/textfile.txt"
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |


  Scenario Outline: do not allow update of reshare to exceed permissions
    Given using OCS API version "<ocs_api_version>"
    And user "Alice" has shared folder "/TMP" with user "Brian" with permissions "share,read"
    And user "Brian" has accepted share "/TMP" offered by user "Alice"
    And user "Brian" has shared folder "Shares/TMP" with user "Carol" with permissions "share,read"
    And user "Carol" has accepted share "/TMP" offered by user "Brian"
    When user "Brian" updates the last share using the sharing API with
      | permissions | all |
    Then the OCS status code should be "404"
    And the HTTP status code should be "<http_status_code>"
    And user "Carol" should not be able to upload file "filesForUpload/textfile.txt" to "Shares/TMP/textfile.txt"
    Examples:
      | ocs_api_version | http_status_code |
      | 1               | 200              |
      | 2               | 404              |


  Scenario Outline: update of user reshare by the original share owner can increase permissions up to the permissions of the top-level share
    Given using OCS API version "<ocs_api_version>"
    And user "Alice" has shared folder "/TMP" with user "Brian" with permissions "share,create,update,read"
    And user "Brian" has accepted share "/TMP" offered by user "Alice"
    And user "Brian" has shared folder "Shares/TMP" with user "Carol" with permissions "share,read"
    And user "Carol" has accepted share "/TMP" offered by user "Brian"
    When user "Alice" updates the last share using the sharing API with
      | permissions | share,create,update,read |
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And user "Carol" should be able to upload file "filesForUpload/textfile.txt" to "Shares/TMP/textfile.txt"
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |


  Scenario Outline: update of user reshare by the original share owner can increase permissions to more than the permissions of the top-level share
    Given using OCS API version "<ocs_api_version>"
    And user "Alice" has shared folder "/TMP" with user "Brian" with permissions "share,update,read"
    And user "Brian" has accepted share "/TMP" offered by user "Alice"
    And user "Brian" has shared folder "Shares/TMP" with user "Carol" with permissions "share,read"
    And user "Carol" has accepted share "/TMP" offered by user "Brian"
    When user "Alice" updates the last share using the sharing API with
      | permissions | share,create,update,read |
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And user "Carol" should be able to upload file "filesForUpload/textfile.txt" to "Shares/TMP/textfile.txt"
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |


  Scenario Outline: update of group reshare by the original share owner can increase permissions up to permissions of the top-level share
    Given using OCS API version "<ocs_api_version>"
    And group "grp1" has been created
    And user "Carol" has been added to group "grp1"
    And user "Alice" has shared folder "/TMP" with user "Brian" with permissions "share,create,update,read"
    And user "Brian" has accepted share "/TMP" offered by user "Alice"
    And user "Brian" has shared folder "Shares/TMP" with group "grp1" with permissions "share,read"
    And user "Carol" has accepted share "/TMP" offered by user "Brian"
    When user "Alice" updates the last share using the sharing API with
      | permissions | share,create,update,read |
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And user "Carol" should be able to upload file "filesForUpload/textfile.txt" to "Shares/TMP/textfile.txt"
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |


  Scenario Outline: update of group reshare by the original share owner can increase permissions to more than the permissions of the top-level share
    Given using OCS API version "<ocs_api_version>"
    And group "grp1" has been created
    And user "Carol" has been added to group "grp1"
    And user "Alice" has shared folder "/TMP" with user "Brian" with permissions "share,update,read"
    And user "Brian" has accepted share "/TMP" offered by user "Alice"
    And user "Brian" has shared folder "Shares/TMP" with group "grp1" with permissions "share,read"
    And user "Carol" has accepted share "/TMP" offered by user "Brian"
    When user "Alice" updates the last share using the sharing API with
      | permissions | share,create,update,read |
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And user "Carol" should be able to upload file "filesForUpload/textfile.txt" to "Shares/TMP/textfile.txt"
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |


  Scenario Outline: after losing share permissions user can still delete a previously reshared share
    Given using OCS API version "<ocs_api_version>"
    And user "Alice" has shared folder "/TMP" with user "Brian" with permissions "share,create,update,read"
    And user "Brian" has accepted share "/TMP" offered by user "Alice"
    And user "Brian" has shared folder "Shares/TMP" with user "Carol" with permissions "share,create,update,read"
    And user "Carol" has accepted share "/TMP" offered by user "Brian"
    And user "Alice" has updated the last share of "Alice" with
      | permissions | create,update,read |
    When user "Brian" deletes the last share using the sharing API
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And user "Carol" should not have any received shares
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |
