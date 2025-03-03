@api @antivirus @skipOnReva
Feature: antivirus
  As a system administrator and user
  I want to protect myself and others from known viruses
  So that I can prevent files with viruses from being uploaded

  Background:
    Given user "Alice" has been created with default attributes and without skeleton files


  Scenario Outline: upload a normal file without virus
    Given using <dav-path-version> DAV path
    When user "Alice" uploads file "filesForUpload/textfile.txt" to "/normalfile.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Alice" file "/normalfile.txt" should exist
    And the content of file "/normalfile.txt" for user "Alice" should be:
    """
    This is a testfile.

    Cheers.
    """
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |


  Scenario Outline: upload a file with virus
    Given using <dav-path-version> DAV path
    When user "Alice" uploads file "filesForUpload/filesWithVirus/eicar.com" to "/aFileWithVirus.txt" using the WebDAV API
    # antivirus service can scan files during post-processing. on demand scanning is currently not available
    Then the HTTP status code should be "201"
    And user "Alice" should get a notification with subject "Virus found" and message:
      | message                                                                             |
      | Virus found in aFileWithVirus.txt. Upload not possible. Virus: Win.Test.EICAR_HDB-1 |
    And as "Alice" file "/aFileWithVirus.txt" should not exist
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |


  Scenario Outline: upload a file with virus and a file without virus
    Given using <dav-path-version> DAV path
    When user "Alice" uploads file "filesForUpload/filesWithVirus/eicar.com" to "/aFileWithVirus.txt" using the WebDAV API
    # antivirus service can scan files during post-processing. on demand scanning is currently not available
    Then the HTTP status code should be "201"
    And user "Alice" uploads file "filesForUpload/textfile.txt" to "/normalfile.txt" using the WebDAV API
    And the HTTP status code should be "201"
    And user "Alice" should get a notification with subject "Virus found" and message:
      | message                                                                             |
      | Virus found in aFileWithVirus.txt. Upload not possible. Virus: Win.Test.EICAR_HDB-1 |
    And as "Alice" file "/aFileWithVirus.txt" should not exist
    But as "Alice" file "/normalfile.txt" should exist
    And the content of file "/normalfile.txt" for user "Alice" should be:
    """
    This is a testfile.

    Cheers.
    """
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |


  Scenario Outline: upload a file with virus in chunks
    Given using <dav-path-version> DAV path
    When user "Alice" uploads the following chunks to "/myChunkedFile.txt" with old chunking and using the WebDAV API
      | number | content                 |
      | 1      | X5O!P%@AP[4\PZX54(P^)7C |
      | 2      | C)7}$EICAR-STANDARD-ANT |
      | 3      | IVIRUS-TEST-FILE!$H+H*  |
    # antivirus service can scan files during post-processing. on demand scanning is currently not available
    Then the HTTP status code should be "201"
    And user "Alice" should get a notification with subject "Virus found" and message:
      | message                                                                            |
      | Virus found in myChunkedFile.txt. Upload not possible. Virus: Win.Test.EICAR_HDB-1 |
    And as "Alice" file "/myChunkedFile.txt" should not exist
    Examples:
      | dav-path-version |
      | old              |
      | spaces           |


  Scenario Outline: upload a file with the virus to a public share
    Given using <dav-path-version> DAV path
    And user "Alice" has created folder "/uploadFolder"
    And user "Alice" has created a public link share with settings
      | path        | /uploadFolder            |
      | name        | sharedlink               |
      | permissions | change                   |
      | expireDate  | 2040-01-01T23:59:59+0100 |
    When user "Alice" uploads file "filesForUpload/filesWithVirus/<filename>" to "/<newfilename>" using the WebDAV API
    Then the HTTP status code should be "201"
    And user "Alice" should get a notification with subject "Virus found" and message:
      | message                                                                        |
      | Virus found in <newfilename>. Upload not possible. Virus: Win.Test.EICAR_HDB-1 |
    And as "Alice" file "/uploadFolder/<newfilename>" should not exist
    Examples:
      | dav-path-version | filename      | newfilename    |
      | old              | eicar.com     | virusFile1.txt |
      | old              | eicar_com.zip | virusFile2.zip |
      | new              | eicar.com     | virusFile1.txt |
      | new              | eicar_com.zip | virusFile2.zip |
      | spaces           | eicar.com     | virusFile1.txt |
      | spaces           | eicar_com.zip | virusFile2.zip |


  Scenario Outline: upload a file with the virus to a password-protected public share
    Given using <dav-path-version> DAV path
    And user "Alice" has created folder "/uploadFolder"
    And user "Alice" has created a public link share with settings
      | path        | /uploadFolder            |
      | name        | sharedlink               |
      | permissions | change                   |
      | password    | newpasswd                |
      | expireDate  | 2040-01-01T23:59:59+0100 |
    When user "Alice" uploads file "filesForUpload/filesWithVirus/<filename>" to "/<newfilename>" using the WebDAV API
    Then the HTTP status code should be "201"
    And user "Alice" should get a notification with subject "Virus found" and message:
      | message                                                                        |
      | Virus found in <newfilename>. Upload not possible. Virus: Win.Test.EICAR_HDB-1 |
    And as "Alice" file "/uploadFolder/<newfilename>" should not exist
    Examples:
      | dav-path-version | filename      | newfilename    |
      | old              | eicar.com     | virusFile1.txt |
      | old              | eicar_com.zip | virusFile2.zip |
      | new              | eicar.com     | virusFile1.txt |
      | new              | eicar_com.zip | virusFile2.zip |
      | spaces           | eicar.com     | virusFile1.txt |
      | spaces           | eicar_com.zip | virusFile2.zip |


  Scenario Outline: upload a file with virus to a user share
    Given using <dav-path-version> DAV path
    And user "Brian" has been created with default attributes and without skeleton files
    And user "Alice" has created folder "uploadFolder"
    And user "Alice" has shared folder "uploadFolder" with user "Brian" with permissions "all"
    And user "Brian" has accepted share "/uploadFolder" offered by user "Alice"
    When user "Brian" uploads file "filesForUpload/filesWithVirus/<filename>" to "/Shares/uploadFolder/<newfilename>" using the WebDAV API
    Then the HTTP status code should be "201"
    And user "Brian" should get a notification with subject "Virus found" and message:
      | message                                                                        |
      | Virus found in <newfilename>. Upload not possible. Virus: Win.Test.EICAR_HDB-1 |
    And as "Brian" file "/Shares/uploadFolder/<newfilename>" should not exist
    And as "Alice" file "/uploadFolder/<newfilename>" should not exist
    Examples:
      | dav-path-version | filename      | newfilename    |
      | old              | eicar.com     | virusFile1.txt |
      | old              | eicar_com.zip | virusFile2.zip |
      | new              | eicar.com     | virusFile1.txt |
      | new              | eicar_com.zip | virusFile2.zip |


  Scenario Outline: upload a file with virus to a user share using spaces dav endpoint
    Given using spaces DAV path
    And user "Brian" has been created with default attributes and without skeleton files
    And user "Alice" has created folder "uploadFolder"
    And user "Alice" has shared folder "uploadFolder" with user "Brian" with permissions "all"
    And user "Brian" has accepted share "/uploadFolder" offered by user "Alice"
    When user "Brian" uploads a file "filesForUpload/filesWithVirus/<filename>" to "/uploadFolder/<newfilename>" in space "Shares" using the WebDAV API
    Then the HTTP status code should be "201"
    And user "Brian" should get a notification with subject "Virus found" and message:
      | message                                                                        |
      | Virus found in <newfilename>. Upload not possible. Virus: Win.Test.EICAR_HDB-1 |
    And as "Brian" file "/Shares/uploadFolder/<newfilename>" should not exist
    And as "Alice" file "/uploadFolder/<newfilename>" should not exist
    Examples:
      | filename      | newfilename    |
      | eicar.com     | virusFile1.txt |
      | eicar_com.zip | virusFile2.zip |


  Scenario Outline: upload a file with virus to a group share
    Given using <dav-path-version> DAV path
    And user "Brian" has been created with default attributes and without skeleton files
    And group "group1" has been created
    And user "Brian" has been added to group "group1"
    And user "Alice" has created folder "uploadFolder"
    And user "Alice" has shared folder "uploadFolder" with group "group1"
    And user "Brian" has accepted share "/uploadFolder" offered by user "Alice"
    When user "Brian" uploads file "filesForUpload/filesWithVirus/<filename>" to "/Shares/uploadFolder/<newfilename>" using the WebDAV API
    Then the HTTP status code should be "201"
    And user "Brian" should get a notification with subject "Virus found" and message:
      | message                                                                        |
      | Virus found in <newfilename>. Upload not possible. Virus: Win.Test.EICAR_HDB-1 |
    And as "Brian" file "/Shares/uploadFolder/<newfilename>" should not exist
    And as "Alice" file "/uploadFolder/<newfilename>" should not exist
    Examples:
      | dav-path-version | filename      | newfilename    |
      | old              | eicar.com     | virusFile1.txt |
      | old              | eicar_com.zip | virusFile2.zip |
      | new              | eicar.com     | virusFile1.txt |
      | new              | eicar_com.zip | virusFile2.zip |


  Scenario Outline: upload a file with virus to a group share using spaces dav endpoint
    Given using spaces DAV path
    And user "Brian" has been created with default attributes and without skeleton files
    And group "group1" has been created
    And user "Brian" has been added to group "group1"
    And user "Alice" has created folder "uploadFolder"
    And user "Alice" has shared folder "uploadFolder" with group "group1"
    And user "Brian" has accepted share "/uploadFolder" offered by user "Alice"
    When user "Brian" uploads a file "filesForUpload/filesWithVirus/<filename>" to "/uploadFolder/<newfilename>" in space "Shares" using the WebDAV API
    Then the HTTP status code should be "201"
    And user "Brian" should get a notification with subject "Virus found" and message:
      | message                                                                        |
      | Virus found in <newfilename>. Upload not possible. Virus: Win.Test.EICAR_HDB-1 |
    And as "Brian" file "/Shares/uploadFolder/<newfilename>" should not exist
    And as "Alice" file "/uploadFolder/<newfilename>" should not exist
    Examples:
      | filename      | newfilename    |
      | eicar.com     | virusFile1.txt |
      | eicar_com.zip | virusFile2.zip |


  Scenario Outline: upload a file with virus to a project space
    Given using spaces DAV path
    And the administrator has assigned the role "Space Admin" to user "Alice" using the Graph API
    And user "Alice" has created a space "new-space" with the default quota using the GraphApi
    And user "Alice" has created a folder "uploadFolder" in space "new-space"
    When user "Alice" uploads a file "filesForUpload/filesWithVirus/<filename>" to "/uploadFolder/<newfilename>" in space "new-space" using the WebDAV API
    Then the HTTP status code should be "201"
    And user "Alice" should get a notification with subject "Virus found" and message:
      | message                                                                        |
      | Virus found in <newfilename>. Upload not possible. Virus: Win.Test.EICAR_HDB-1 |
    And for user "Alice" the space "new-space" should not contain these entries:
      | /uploadFolder/<newfilename> |
    When user "Alice" uploads a file "filesForUpload/filesWithVirus/<filename>" to "/<newfilename>" in space "new-space" using the WebDAV API
    Then the HTTP status code should be "201"
    And user "Alice" should get a notification with subject "Virus found" and message:
      | message                                                                        |
      | Virus found in <newfilename>. Upload not possible. Virus: Win.Test.EICAR_HDB-1 |
    And for user "Alice" the space "new-space" should not contain these entries:
      | /<newfilename> |
    Examples:
      | filename      | newfilename    |
      | eicar.com     | virusFile1.txt |
      | eicar_com.zip | virusFile2.zip |


  Scenario Outline: upload a file with virus to a shared project space
    Given using spaces DAV path
    And user "Brian" has been created with default attributes and without skeleton files
    And the administrator has assigned the role "Space Admin" to user "Alice" using the Graph API
    And user "Alice" has created a space "new-space" with the default quota using the GraphApi
    And user "Alice" has shared a space "new-space" with settings:
      | shareWith | Brian  |
      | role      | editor |
    When user "Brian" uploads a file "/filesForUpload/filesWithVirus/<filename>" to "/<newfilename>" in space "new-space" using the WebDAV API
    Then the HTTP status code should be "201"
    And user "Brian" should get a notification with subject "Virus found" and message:
      | message                                                                        |
      | Virus found in <newfilename>. Upload not possible. Virus: Win.Test.EICAR_HDB-1 |
    And for user "Brian" the space "new-space" should not contain these entries:
      | /<newfilename> |
    And for user "Alice" the space "new-space" should not contain these entries:
      | /<newfilename> |
    Examples:
      | filename      | newfilename    |
      | eicar.com     | virusFile1.txt |
      | eicar_com.zip | virusFile2.zip |

  @env-config  @issue-6494
  Scenario Outline: upload a file with virus by setting antivirus infected file handling config to continue
    Given the config "ANTIVIRUS_INFECTED_FILE_HANDLING" has been set to "continue"
    And using <dav-path-version> DAV path
    When user "Alice" uploads file "filesForUpload/filesWithVirus/eicar.com" to "/aFileWithVirus.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Alice" file "/aFileWithVirus.txt" should exist
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |

  @env-config
  Scenario Outline: upload a file with virus smaller than the upload threshold
    Given the config "ANTIVIRUS_MAX_SCAN_SIZE" has been set to "100"
    And using <dav-path-version> DAV path
    When user "Alice" uploads file "filesForUpload/filesWithVirus/eicar.com" to "/aFileWithVirus.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And user "Alice" should get a notification with subject "Virus found" and message:
      | message                                                                             |
      | Virus found in aFileWithVirus.txt. Upload not possible. Virus: Win.Test.EICAR_HDB-1 |
    And as "Alice" file "/aFileWithVirus.txt" should not exist
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |

  @env-config
  Scenario Outline: upload a file with virus larger than the upload threshold
    Given the config "ANTIVIRUS_MAX_SCAN_SIZE" has been set to "100"
    And using <dav-path-version> DAV path
    When user "Alice" uploads file "filesForUpload/filesWithVirus/eicar_com.zip" to "/aFileWithVirus.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Alice" file "/aFileWithVirus.txt" should exist
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |

  @issue-enterprise-5706
  Scenario Outline: upload a file with virus and get notification in different languages
    Given user "Alice" has switched the system language to "<language>"
    And using <dav-path-version> DAV path
    When user "Alice" uploads file "filesForUpload/filesWithVirus/eicar.com" to "/aFileWithVirus.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And user "Alice" should get a notification with subject "<subject>" and message:
      | message   |
      | <message> |
    And as "Alice" file "/aFileWithVirus.txt" should not exist
    Examples:
      | dav-path-version | language | subject          | message                                                                                                                        |
      | old              | es       | Virus encontrado | Virus encontrado en aFileWithVirus.txt. La subida no ha sido posible. Virus: Win.Test.EICAR_HDB-1                              |
      | new              | es       | Virus encontrado | Virus encontrado en aFileWithVirus.txt. La subida no ha sido posible. Virus: Win.Test.EICAR_HDB-1                              |
      | spaces           | es       | Virus encontrado | Virus encontrado en aFileWithVirus.txt. La subida no ha sido posible. Virus: Win.Test.EICAR_HDB-1                              |
      | old              | de       | Virus gefunden   | In aFileWithVirus.txt wurde potenziell schädlicher Code gefunden. Das Hochladen wurde abgebrochen. Grund: Win.Test.EICAR_HDB-1 |
      | new              | de       | Virus gefunden   | In aFileWithVirus.txt wurde potenziell schädlicher Code gefunden. Das Hochladen wurde abgebrochen. Grund: Win.Test.EICAR_HDB-1 |
      | spaces           | de       | Virus gefunden   | In aFileWithVirus.txt wurde potenziell schädlicher Code gefunden. Das Hochladen wurde abgebrochen. Grund: Win.Test.EICAR_HDB-1 |

  @issue-enterprise-5709
  Scenario Outline: try to create a version of file by uploading virus content
    Given using <dav-path-version> DAV path
    And user "Alice" has uploaded file with content "hello world" to "test.txt"
    And user "Alice" has uploaded file with content "hello nepal" to "test.txt"
    When user "Alice" uploads file with content "X5O!P%@AP[4\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*" to "test.txt" using the WebDAV API
    Then the HTTP status code should be "204"
    And user "Alice" should get a notification with subject "Virus found" and message:
      | message                                                                   |
      | Virus found in test.txt. Upload not possible. Virus: Win.Test.EICAR_HDB-1 |
    And as "Alice" file "/test.txt" should exist
    And the version folder of file "/test.txt" for user "Alice" should contain "1" element
    And the content of file "/test.txt" for user "Alice" should be "hello nepal"
    When user "Alice" restores version index "1" of file "/test.txt" using the WebDAV API
    Then the HTTP status code should be "204"
    And the content of file "/test.txt" for user "Alice" should be "hello world"
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |


  Scenario Outline: try to overwrite a file with the virus content in a public link share
    Given using <dav-path-version> DAV path
    And user "Alice" has uploaded file with content "hello" to "test.txt"
    And user "Alice" has created a public link share with settings
      | path        | /test.txt  |
      | name        | sharedlink |
      | permissions | change     |
    When the public overwrites file "test.txt" with content "X5O!P%@AP[4\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*" using the new WebDAV API
    Then the HTTP status code should be "204"
    And user "Alice" should get a notification with subject "Virus found" and message:
      | message                                                                   |
      | Virus found in test.txt. Upload not possible. Virus: Win.Test.EICAR_HDB-1 |
    And the content of file "/test.txt" for user "Alice" should be "hello"
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |


  Scenario Outline: try to overwrite a file with the virus content in group share
    Given using <dav-path-version> DAV path
    And user "Brian" has been created with default attributes and without skeleton files
    And group "group1" has been created
    And user "Brian" has been added to group "group1"
    And user "Alice" has been added to group "group1"
    And user "Alice" has uploaded file with content "hello" to "/test.txt"
    And user "Alice" has shared file "test.txt" with group "group1"
    And user "Brian" has accepted share "/test.txt" offered by user "Alice"
    When user "Brian" uploads file with content "X5O!P%@AP[4\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*" to "test.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And user "Brian" should get a notification with subject "Virus found" and message:
      | message                                                                   |
      | Virus found in test.txt. Upload not possible. Virus: Win.Test.EICAR_HDB-1 |
    And the content of file "/test.txt" for user "Alice" should be "hello"
    And the content of file "Shares/test.txt" for user "Brian" should be "hello"
    Examples:
      | dav-path-version |
      | old              |
      | new              |


  Scenario: try to overwrite a file with the virus content in group share using spaces dav endpoint
    Given using spaces DAV path
    And user "Brian" has been created with default attributes and without skeleton files
    And group "group1" has been created
    And user "Brian" has been added to group "group1"
    And user "Alice" has been added to group "group1"
    And user "Alice" has uploaded file with content "hello" to "/test.txt"
    And user "Alice" has shared file "test.txt" with group "group1"
    And user "Brian" has accepted share "/test.txt" offered by user "Alice"
    When user "Brian" uploads a file "filesForUpload/filesWithVirus/eicar.com" to "/test.txt" in space "Shares" using the WebDAV API
    Then the HTTP status code should be "204"
    And user "Brian" should get a notification with subject "Virus found" and message:
      | message                                                                   |
      | Virus found in test.txt. Upload not possible. Virus: Win.Test.EICAR_HDB-1 |
    And for user "Brian" the content of the file "/test.txt" of the space "Shares" should be "hello"
    And the content of file "/test.txt" for user "Alice" should be "hello"
