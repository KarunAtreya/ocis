@api
Feature: propagation of etags when creating folders
  As a client app
  I want metadata (etags) of parent folders to change when a sub-folder is created
  So that the client app can know to re-scan and sync the content of the folder(s)

  Background:
    Given user "Alice" has been created with default attributes and without skeleton files

  @issue-product-280
  Scenario Outline: creating a folder inside a folder changes its etag
    Given using <dav_version> DAV path
    And user "Alice" has created folder "/folder"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/folder"
    When user "Alice" creates folder "/folder/new" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path    |
      | Alice | /       |
      | Alice | /folder |
    Examples:
      | dav_version |
      | old         |
      | new         |

    @skipOnRevaMaster
    Examples:
      | dav_version |
      | spaces      |


  Scenario Outline: creating an invalid folder inside a folder should not change any etags
    Given using <dav_version> DAV path
    And user "Alice" has created folder "/folder"
    And user "Alice" has created folder "/folder/sub"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/folder"
    And user "Alice" has stored etag of element "/folder/sub"
    When user "Alice" creates folder "/folder/sub/.." using the WebDAV API
    Then the HTTP status code should be "405"
    And these etags should not have changed:
      | user  | path        |
      | Alice | /           |
      | Alice | /folder     |
      | Alice | /folder/sub |
    Examples:
      | dav_version |
      | old         |
      | new         |

    @skipOnRevaMaster
    Examples:
      | dav_version |
      | spaces      |

  @issue-product-280
  Scenario Outline: sharee creating a folder inside a folder received as a share changes its etag for all collaborators
    Given user "Brian" has been created with default attributes and without skeleton files
    And using <dav_version> DAV path
    And user "Alice" has created folder "/folder"
    And user "Alice" has shared folder "/folder" with user "Brian"
    And user "Brian" has accepted share "/folder" offered by user "Alice"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/folder"
    And user "Brian" has stored etag of element "/"
    And user "Brian" has stored etag of element "/Shares"
    And user "Brian" has stored etag of element "/Shares/folder"
    When user "Brian" creates folder "/Shares/folder/new" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path           |
      | Alice | /              |
      | Alice | /folder        |
      | Brian | /              |
      | Brian | /Shares        |
      | Brian | /Shares/folder |
    Examples:
      | dav_version |
      | old         |
      | new         |

  @issue-product-280
  Scenario Outline: sharer creating a folder inside a shared folder changes etag for all collaborators
    Given user "Brian" has been created with default attributes and without skeleton files
    And using <dav_version> DAV path
    And user "Alice" has created folder "/folder"
    And user "Alice" has shared folder "/folder" with user "Brian"
    And user "Brian" has accepted share "/folder" offered by user "Alice"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/folder"
    And user "Brian" has stored etag of element "/"
    And user "Brian" has stored etag of element "/Shares"
    And user "Brian" has stored etag of element "/Shares/folder"
    When user "Alice" creates folder "/folder/new" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path           |
      | Alice | /              |
      | Alice | /folder        |
      | Brian | /              |
      | Brian | /Shares        |
      | Brian | /Shares/folder |
    Examples:
      | dav_version |
      | old         |
      | new         |

  @issue-product-280
  Scenario Outline: creating a folder in a publicly shared folder changes its etag for the sharer
    Given using <dav_version> DAV path
    And user "Alice" has created folder "/folder"
    And user "Alice" has created a public link share with settings
      | path        | folder |
      | permissions | create |
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/folder"
    When the public creates folder "created-by-public" using the new public WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path    |
      | Alice | /       |
      | Alice | /folder |
    Examples:
      | dav_version |
      | old         |
      | new         |

  @skipOnRevaMaster
    Examples:
      | dav_version |
      | spaces      |
