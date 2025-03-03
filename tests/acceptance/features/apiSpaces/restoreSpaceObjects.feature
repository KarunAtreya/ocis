@api 
Feature: Restore files, folder
  As a user with manager and editor role
  I want to be able to restore files, folders
  So that I can get the resources that were accidentally deleted

  Note - this feature is run in CI with ACCOUNTS_HASH_DIFFICULTY set to the default for production
  See https://github.com/owncloud/ocis/issues/1542 and https://github.com/owncloud/ocis/pull/839

  Background:
    Given these users have been created with default attributes and without skeleton files:
      | username |
      | Alice    |
      | Brian    |
      | Bob      |
      | Carol    |
    And using spaces DAV path
    And the administrator has assigned the role "Space Admin" to user "Alice" using the Graph API
    And user "Alice" has created a space "restore objects" with the default quota using the GraphApi
    And user "Alice" has created a folder "newFolder" in space "restore objects"
    And user "Alice" has uploaded a file inside space "restore objects" with content "test" to "newFolder/file.txt"


  Scenario Outline: user with different role can see deleted objects in trash bin of the space via the webDav API
    Given user "Alice" has shared a space "restore objects" with settings:
      | shareWith | Brian  |
      | role      | <role> |
    And user "Alice" has removed the file "newFolder/file.txt" from space "restore objects"
    And user "Alice" has removed the folder "newFolder" from space "restore objects"
    When user "<user>" lists all deleted files in the trash bin of the space "restore objects"
    Then the HTTP status code should be "207"
    And as "<user>" folder "newFolder" should exist in the trashbin of the space "restore objects"
    And as "<user>" file "file.txt" should exist in the trashbin of the space "restore objects"
    Examples:
      | user  | role    |
      | Brian | manager |
      | Brian | editor  |
      | Brian | viewer  |


  Scenario Outline: user can restore a folder with some objects from the trash via the webDav API
    Given user "Alice" has shared a space "restore objects" with settings:
      | shareWith | Brian  |
      | role      | <role> |
    And user "Alice" has removed the folder "newFolder" from space "restore objects"
    When user "<user>" restores the folder "newFolder" from the trash of the space "restore objects" to "/newFolder"
    Then the HTTP status code should be "<code>"
    And for user "<user>" the space "restore objects" <shouldOrNotBeInSpace> contain these entries:
      | newFolder |
    And as "<user>" folder "newFolder" <shouldOrNotBeInTrash> exist in the trashbin of the space "restore objects"
    Examples:
      | user  | role    | code | shouldOrNotBeInSpace | shouldOrNotBeInTrash |
      | Alice | manager | 201  | should               | should not           |
      | Brian | manager | 201  | should               | should not           |
      | Brian | editor  | 201  | should               | should not           |
      | Brian | viewer  | 403  | should not           | should               |


  Scenario Outline: user can restore a file from the trash via the webDav API
    Given user "Alice" has shared a space "restore objects" with settings:
      | shareWith | Brian  |
      | role      | <role> |
    And user "Alice" has removed the file "newFolder/file.txt" from space "restore objects"
    When user "<user>" restores the file "file.txt" from the trash of the space "restore objects" to "newFolder/file.txt"
    Then the HTTP status code should be "<code>"
    And for user "<user>" folder "newFolder" of the space "restore objects" <shouldOrNotBeInSpace> contain these files:
      | file.txt |
    And as "<user>" file "file.txt" <shouldOrNotBeInTrash> exist in the trashbin of the space "restore objects"
    Examples:
      | user  | role    | code | shouldOrNotBeInSpace | shouldOrNotBeInTrash |
      | Alice | manager | 201  | should               | should not           |
      | Brian | manager | 201  | should               | should not           |
      | Brian | editor  | 201  | should               | should not           |
      | Brian | viewer  | 403  | should not           | should               |


  Scenario: user can restore a file even if there is not enough quota to do so
    Given user "Admin" has changed the quota of the "Brian Murphy" space to "30"
    And user "Brian" has uploaded file with content "file is less than 30 bytes" to "/file.txt"
    And user "Brian" has uploaded file with content "reduceContent" to "/file.txt"
    And user "Brian" has uploaded file with content "some content" to "newFile.txt"
    When user "Brian" restores version index "1" of file "/file.txt" using the WebDAV API
    Then the HTTP status code should be "204"
    And the content of file "/file.txt" for user "Brian" should be "file is less than 30 bytes"
