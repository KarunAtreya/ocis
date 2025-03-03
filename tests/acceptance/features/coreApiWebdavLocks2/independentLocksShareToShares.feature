@api @issue-1284
Feature: independent locks - make sure all locks are independent and don't interact with other items that have the same name
  As a user
  I want to independently lock resources shared with me
  So that resources with same name in other received shares will not be locked

  Background:
    Given user "Alice" has been created with default attributes and without skeleton files
    And user "Brian" has been created with default attributes and without skeleton files


  Scenario Outline: locking a file in a received share does not lock other items with the same name in other received shares (shares from different users)
    Given using <dav-path> DAV path
    And user "Carol" has been created with default attributes and without skeleton files
    And user "Alice" has created folder "FromAlice"
    And user "Brian" has created folder "FromBrian"
    And user "Alice" has uploaded file "filesForUpload/textfile.txt" to "/FromAlice/textfile0.txt"
    And user "Brian" has uploaded file "filesForUpload/textfile.txt" to "/FromBrian/textfile0.txt"
    And user "Alice" has shared folder "FromAlice" with user "Carol"
    And user "Brian" has shared folder "FromBrian" with user "Carol"
    And user "Carol" has accepted share "/FromAlice" offered by user "Alice"
    And user "Carol" has accepted share "/FromBrian" offered by user "Brian"
    When user "Carol" locks file "/Shares/FromBrian/textfile0.txt" using the WebDAV API setting the following properties
      | lockscope | <lock-scope> |
    Then the HTTP status code should be "200"
    And user "Carol" should be able to upload file "filesForUpload/lorem.txt" to "/Shares/FromAlice/textfile0.txt"
    But user "Carol" should not be able to upload file "filesForUpload/lorem.txt" to "/Shares/FromBrian/textfile0.txt"
    Examples:
      | dav-path | lock-scope |
      | old      | shared     |
      | old      | exclusive  |
      | new      | shared     |
      | new      | exclusive  |

    @skipOnRevaMaster
    Examples:
      | dav-path | lock-scope |
      | spaces   | shared     |
      | spaces   | exclusive  |


  Scenario Outline: locking a file in a received share does not lock other items with the same name in other received shares (shares from same user)
    Given using <dav-path> DAV path
    And user "Alice" has created folder "locked/"
    And user "Alice" has created folder "notlocked/"
    And user "Alice" has uploaded file "filesForUpload/textfile.txt" to "/locked/textfile0.txt"
    And user "Alice" has uploaded file "filesForUpload/textfile.txt" to "/notlocked/textfile0.txt"
    And user "Alice" has shared folder "locked" with user "Brian"
    And user "Alice" has shared folder "notlocked" with user "Brian"
    And user "Brian" has accepted share "/locked" offered by user "Alice"
    And user "Brian" has accepted share "/notlocked" offered by user "Alice"
    When user "Brian" locks file "/Shares/locked/textfile0.txt" using the WebDAV API setting the following properties
      | lockscope | <lock-scope> |
    Then the HTTP status code should be "200"
    And user "Brian" should be able to upload file "filesForUpload/lorem.txt" to "/Shares/notlocked/textfile0.txt"
    But user "Brian" should not be able to upload file "filesForUpload/lorem.txt" to "/Shares/locked/textfile0.txt"
    Examples:
      | dav-path | lock-scope |
      | old      | shared     |
      | old      | exclusive  |
      | new      | shared     |
      | new      | exclusive  |

    @skipOnRevaMaster
    Examples:
      | dav-path | lock-scope |
      | spaces   | shared     |
      | spaces   | exclusive  |
