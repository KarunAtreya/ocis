@api
Feature: propagation of etags when moving files or folders
  As a client app
  I want metadata (etags) of parent folder(s) to change when a file or folder is moved
  So that the client app can know to re-scan and sync the content of the folder(s)

  Background:
    Given user "Alice" has been created with default attributes and without skeleton files


  Scenario Outline: renaming a file inside a folder changes its etag
    Given using <dav_version> DAV path
    And user "Alice" has created folder "/upload"
    And user "Alice" has uploaded file with content "uploaded content" to "/upload/file.txt"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    When user "Alice" moves file "/upload/file.txt" to "/upload/renamed.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path    |
      | Alice | /       |
      | Alice | /upload |
    Examples:
      | dav_version |
      | old         |
      | new         |

    @skipOnRevaMaster
    Examples:
      | dav_version |
      | spaces      |

  @issue-product-280
  Scenario Outline: moving a file from one folder to an other changes the etags of both folders
    Given using <dav_version> DAV path
    And user "Alice" has created folder "/src"
    And user "Alice" has created folder "/dst"
    And user "Alice" has uploaded file with content "uploaded content" to "/src/file.txt"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/src"
    And user "Alice" has stored etag of element "/dst"
    When user "Alice" moves file "/src/file.txt" to "/dst/file.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path |
      | Alice | /    |
      | Alice | /src |
      | Alice | /dst |
    Examples:
      | dav_version |
      | old         |
      | new         |

    @skipOnRevaMaster
    Examples:
      | dav_version |
      | spaces      |

  @issue-product-280
  Scenario Outline: moving a file into a subfolder changes the etags of all parents
    Given using <dav_version> DAV path
    And user "Alice" has created folder "/upload"
    And user "Alice" has created folder "/upload/sub"
    And user "Alice" has uploaded file with content "uploaded content" to "/upload/file.txt"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    And user "Alice" has stored etag of element "/upload/sub"
    When user "Alice" moves file "/upload/file.txt" to "/upload/sub/file.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path        |
      | Alice | /           |
      | Alice | /upload     |
      | Alice | /upload/sub |
    Examples:
      | dav_version |
      | old         |
      | new         |

    @skipOnRevaMaster
    Examples:
      | dav_version |
      | spaces      |


  Scenario Outline: renaming a folder inside a folder changes its etag
    Given using <dav_version> DAV path
    And user "Alice" has created folder "/upload"
    And user "Alice" has created folder "/upload/src"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    When user "Alice" moves folder "/upload/src" to "/upload/dst" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path    |
      | Alice | /       |
      | Alice | /upload |
    Examples:
      | dav_version |
      | old         |
      | new         |

    @skipOnRevaMaster
    Examples:
      | dav_version |
      | spaces      |

  @issue-product-280
  Scenario Outline: moving a folder from one folder to an other changes the etags of both folders
    Given using <dav_version> DAV path
    And user "Alice" has created folder "/src"
    And user "Alice" has created folder "/src/folder"
    And user "Alice" has created folder "/dst"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/src"
    And user "Alice" has stored etag of element "/dst"
    When user "Alice" moves folder "/src/folder" to "/dst/folder" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path |
      | Alice | /    |
      | Alice | /src |
      | Alice | /dst |
    Examples:
      | dav_version |
      | old         |
      | new         |

    @skipOnRevaMaster
    Examples:
      | dav_version |
      | spaces      |

  @issue-product-280
  Scenario Outline: moving a folder into a subfolder changes the etags of all parents
    Given using <dav_version> DAV path
    And user "Alice" has created folder "/upload"
    And user "Alice" has created folder "/upload/folder"
    And user "Alice" has created folder "/upload/sub"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    And user "Alice" has stored etag of element "/upload/sub"
    When user "Alice" moves folder "/upload/folder" to "/upload/sub/folder" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path        |
      | Alice | /           |
      | Alice | /upload     |
      | Alice | /upload/sub |
    Examples:
      | dav_version |
      | old         |
      | new         |

    @skipOnRevaMaster
    Examples:
      | dav_version |
      | spaces      |


  Scenario Outline: sharee renaming a file inside a folder changes its etag for all collaborators
    Given user "Brian" has been created with default attributes and without skeleton files
    And using <dav_version> DAV path
    And user "Alice" has created folder "/upload"
    And user "Alice" has uploaded file with content "uploaded content" to "/upload/file.txt"
    And user "Alice" has shared folder "/upload" with user "Brian"
    And user "Brian" has accepted share "/upload" offered by user "Alice"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    And user "Brian" has stored etag of element "/"
    And user "Brian" has stored etag of element "/Shares"
    And user "Brian" has stored etag of element "/Shares/upload"
    When user "Brian" moves file "/Shares/upload/file.txt" to "/Shares/upload/renamed.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path           |
      | Alice | /              |
      | Alice | /upload        |
      | Brian | /              |
      | Brian | /Shares        |
      | Brian | /Shares/upload |
    Examples:
      | dav_version |
      | old         |
      | new         |


  Scenario Outline: sharer renaming a file inside a folder changes its etag for all collaborators
    Given user "Brian" has been created with default attributes and without skeleton files
    And using <dav_version> DAV path
    And user "Alice" has created folder "/upload"
    And user "Alice" has uploaded file with content "uploaded content" to "/upload/file.txt"
    And user "Alice" has shared folder "/upload" with user "Brian"
    And user "Brian" has accepted share "/upload" offered by user "Alice"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    And user "Brian" has stored etag of element "/"
    And user "Brian" has stored etag of element "/Shares"
    And user "Brian" has stored etag of element "/Shares/upload"
    When user "Alice" moves file "/upload/file.txt" to "/upload/renamed.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path           |
      | Alice | /              |
      | Alice | /upload        |
      | Brian | /              |
      | Brian | /Shares        |
      | Brian | /Shares/upload |
    Examples:
      | dav_version |
      | old         |
      | new         |

  @issue-product-280
  Scenario Outline: sharer moving a file from one folder to an other changes the etags of both folders for all collaborators
    Given user "Brian" has been created with default attributes and without skeleton files
    And using <dav_version> DAV path
    And user "Alice" has created folder "/src"
    And user "Alice" has created folder "/dst"
    And user "Alice" has uploaded file with content "uploaded content" to "/src/file.txt"
    And user "Alice" has shared folder "/src" with user "Brian"
    And user "Brian" has accepted share "/src" offered by user "Alice"
    And user "Alice" has shared folder "/dst" with user "Brian"
    And user "Brian" has accepted share "/dst" offered by user "Alice"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/src"
    And user "Alice" has stored etag of element "/dst"
    And user "Brian" has stored etag of element "/"
    And user "Brian" has stored etag of element "/Shares"
    And user "Brian" has stored etag of element "/Shares/src"
    And user "Brian" has stored etag of element "/Shares/dst"
    When user "Alice" moves file "/src/file.txt" to "/dst/file.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path        |
      | Alice | /           |
      | Alice | /src        |
      | Alice | /dst        |
      | Brian | /           |
      | Brian | /Shares     |
      | Brian | /Shares/src |
      | Brian | /Shares/dst |
    Examples:
      | dav_version |
      | old         |
      | new         |


  Scenario Outline: sharee moving a file from one folder to an other changes the etags of both folders for all collaborators
    Given user "Brian" has been created with default attributes and without skeleton files
    And using <dav_version> DAV path
    And user "Alice" has created folder "/src"
    And user "Alice" has created folder "/dst"
    And user "Alice" has uploaded file with content "uploaded content" to "/src/file.txt"
    And user "Alice" has shared folder "/src" with user "Brian"
    And user "Brian" has accepted share "/src" offered by user "Alice"
    And user "Alice" has shared folder "/dst" with user "Brian"
    And user "Brian" has accepted share "/dst" offered by user "Alice"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/src"
    And user "Alice" has stored etag of element "/dst"
    And user "Brian" has stored etag of element "/"
    And user "Brian" has stored etag of element "/Shares"
    And user "Brian" has stored etag of element "/Shares/src"
    And user "Brian" has stored etag of element "/Shares/dst"
    When user "Brian" moves file "/Shares/src/file.txt" to "/Shares/dst/file.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path        |
      | Alice | /           |
      | Alice | /src        |
      | Alice | /dst        |
      | Brian | /           |
      | Brian | /Shares     |
      | Brian | /Shares/src |
      | Brian | /Shares/dst |
    Examples:
      | dav_version |
      | old         |
      | new         |

  @issue-product-280
  Scenario Outline: sharer moving a folder from one folder to an other changes the etags of both folders for all collaborators
    Given user "Brian" has been created with default attributes and without skeleton files
    And using <dav_version> DAV path
    And user "Alice" has created folder "/src"
    And user "Alice" has created folder "/dst"
    And user "Alice" has created folder "/src/toMove"
    And user "Alice" has shared folder "/src" with user "Brian"
    And user "Brian" has accepted share "/src" offered by user "Alice"
    And user "Alice" has shared folder "/dst" with user "Brian"
    And user "Brian" has accepted share "/dst" offered by user "Alice"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/src"
    And user "Alice" has stored etag of element "/dst"
    And user "Brian" has stored etag of element "/"
    And user "Brian" has stored etag of element "/Shares"
    And user "Brian" has stored etag of element "/Shares/src"
    And user "Brian" has stored etag of element "/Shares/dst"
    When user "Alice" moves folder "/src/toMove" to "/dst/toMove" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path        |
      | Alice | /           |
      | Alice | /src        |
      | Alice | /dst        |
      | Brian | /           |
      | Brian | /Shares     |
      | Brian | /Shares/src |
      | Brian | /Shares/dst |
    Examples:
      | dav_version |
      | old         |
      | new         |


  Scenario Outline: sharee moving a folder from one folder to an other changes the etags of both folders for all collaborators
    Given user "Brian" has been created with default attributes and without skeleton files
    And using <dav_version> DAV path
    And user "Alice" has created folder "/src"
    And user "Alice" has created folder "/dst"
    And user "Alice" has created folder "/src/toMove"
    And user "Alice" has shared folder "/src" with user "Brian"
    And user "Brian" has accepted share "/src" offered by user "Alice"
    And user "Alice" has shared folder "/dst" with user "Brian"
    And user "Brian" has accepted share "/dst" offered by user "Alice"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/src"
    And user "Alice" has stored etag of element "/dst"
    And user "Brian" has stored etag of element "/"
    And user "Brian" has stored etag of element "/Shares"
    And user "Brian" has stored etag of element "/Shares/src"
    And user "Brian" has stored etag of element "/Shares/dst"
    When user "Brian" moves folder "/Shares/src/toMove" to "/Shares/dst/toMove" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path        |
      | Alice | /           |
      | Alice | /src        |
      | Alice | /dst        |
      | Brian | /           |
      | Brian | /Shares     |
      | Brian | /Shares/src |
      | Brian | /Shares/dst |
    Examples:
      | dav_version |
      | old         |
      | new         |


  Scenario Outline: renaming a file in a publicly shared folder changes its etag for the sharer
    Given using <dav_version> DAV path
    And user "Alice" has created folder "/upload"
    And user "Alice" has uploaded file with content "uploaded content" to "/upload/file.txt"
    And user "Alice" has created a public link share with settings
      | path        | upload |
      | permissions | change |
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    When the public renames file "file.txt" to "renamed.txt" from the last public link share using the new public WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path    |
      | Alice | /       |
      | Alice | /upload |
    Examples:
      | dav_version |
      | old         |
      | new         |

    @skipOnRevaMaster
    Examples:
      | dav_version |
      | spaces      |


  Scenario Outline: renaming a folder in a publicly shared folder changes its etag for the sharer
    Given using <dav_version> DAV path
    And user "Alice" has created folder "/upload"
    And user "Alice" has created folder "/upload/sub"
    And user "Alice" has created a public link share with settings
      | path        | upload |
      | permissions | change |
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    When the public renames folder "sub" to "renamed" from the last public link share using the new public WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path    |
      | Alice | /       |
      | Alice | /upload |
    Examples:
      | dav_version |
      | old         |
      | new         |

    @skipOnRevaMaster
    Examples:
      | dav_version |
      | spaces      |
