# Changelog for [unreleased] (UNRELEASED)

The following sections list the changes for unreleased.

[unreleased]: https://github.com/owncloud/ocis/compare/v3.0.0...master

## Summary

* Bugfix - Add default store to postprocessing: [#6578](https://github.com/owncloud/ocis/pull/6578)
* Bugfix - Add token to LinkAccessedEvent: [#6554](https://github.com/owncloud/ocis/pull/6554)
* Bugfix - Add missing timestamps: [#6515](https://github.com/owncloud/ocis/pull/6515)
* Bugfix - Don't connect to ldap on startup: [#6565](https://github.com/owncloud/ocis/pull/6565)
* Bugfix - Handle the bad request status: [#6469](https://github.com/owncloud/ocis/pull/6469)
* Bugfix - Fix the oidc role assigner: [#6605](https://github.com/owncloud/ocis/pull/6605)
* Enhancement - Add 'ocis decomposedfs check-treesize' command: [#6556](https://github.com/owncloud/ocis/pull/6556)
* Enhancement - Add IDs to graph resource logging: [#6593](https://github.com/owncloud/ocis/pull/6593)
* Enhancement - Add permissions to report: [#6528](https://github.com/owncloud/ocis/pull/6528)
* Enhancement - Add more metadata to the remote item: [#6300](https://github.com/owncloud/ocis/pull/6300)
* Enhancement - We added the storage id to the audit log for spaces: [#6548](https://github.com/owncloud/ocis/pull/6548)
* Enhancement - Make the post logout redirect uri configurable: [#6583](https://github.com/owncloud/ocis/pull/6583)
* Enhancement - Make the app provider service name configurable: [#6482](https://github.com/owncloud/ocis/pull/6482)
* Enhancement - Add old & new values to audit logs: [#6537](https://github.com/owncloud/ocis/pull/6537)
* Enhancement - Update reva: [#6529](https://github.com/owncloud/ocis/pull/6529)
* Enhancement - Use reva client selectors: [#6452](https://github.com/owncloud/ocis/pull/6452)
* Enhancement - Add companion URL config: [#6453](https://github.com/owncloud/ocis/pull/6453)
* Enhancement - Add imprint and privacy url config: [#6462](https://github.com/owncloud/ocis/pull/6462)
* Enhancement - Add logged out url config: [#6549](https://github.com/owncloud/ocis/pull/6549)
* Enhancement - Fix envvar defaults: [#6516](https://github.com/owncloud/ocis/pull/6516)
* Enhancement - Skip if the simulink is a directory: [#6574](https://github.com/owncloud/ocis/pull/6574)
* Enhancement - Fix the groupname validation: [#6490](https://github.com/owncloud/ocis/pull/6490)
* Enhancement - Fix the username validation: [#6437](https://github.com/owncloud/ocis/pull/6437)
* Enhancement - Move proxy to service tracerprovider: [#6591](https://github.com/owncloud/ocis/pull/6591)
* Enhancement - Add functionality to retry postprocessing: [#6500](https://github.com/owncloud/ocis/pull/6500)
* Enhancement - Update go-micro kubernetes registry: [#6457](https://github.com/owncloud/ocis/pull/6457)
* Enhancement - Update web to v7.0.1: [#6470](https://github.com/owncloud/ocis/pull/6470)
* Enhancement - Allow disabling wopi chat: [#6544](https://github.com/owncloud/ocis/pull/6544)

## Details

* Bugfix - Add default store to postprocessing: [#6578](https://github.com/owncloud/ocis/pull/6578)

   Postprocessing did not have a default store especially `database` and `table` are needed to
   talk to nats-js

   https://github.com/owncloud/ocis/pull/6578

* Bugfix - Add token to LinkAccessedEvent: [#6554](https://github.com/owncloud/ocis/pull/6554)

   We added the link token to the LinkAccessedEvent

   https://github.com/owncloud/ocis/issues/3753
   https://github.com/owncloud/ocis/pull/6554
   https://github.com/cs3org/reva/pull/3993

* Bugfix - Add missing timestamps: [#6515](https://github.com/owncloud/ocis/pull/6515)

   We have added missing timestamps to the audit service

   https://github.com/owncloud/ocis/issues/3753
   https://github.com/owncloud/ocis/pull/6515

* Bugfix - Don't connect to ldap on startup: [#6565](https://github.com/owncloud/ocis/pull/6565)

   This leads to misleading error messages. Instead we connect on first request

   https://github.com/owncloud/ocis/pull/6565

* Bugfix - Handle the bad request status: [#6469](https://github.com/owncloud/ocis/pull/6469)

   Handle the bad request status for the CreateStorageSpace function

   https://github.com/owncloud/ocis/issues/6414
   https://github.com/owncloud/ocis/pull/6469
   https://github.com/cs3org/reva/pull/3948

* Bugfix - Fix the oidc role assigner: [#6605](https://github.com/owncloud/ocis/pull/6605)

   The update role method did not allow to set a role when the user already has two roles. This makes
   no sense as the user is supposed to have only one and the update will fix that. We still log an error
   level log to make the admin aware of that.

   https://github.com/owncloud/ocis/pull/6605

* Enhancement - Add 'ocis decomposedfs check-treesize' command: [#6556](https://github.com/owncloud/ocis/pull/6556)

   We added a 'ocis decomposedfs check-treesize' command for checking (and reparing) the
   treesize metadata of a storage space.

   https://github.com/owncloud/ocis/pull/6556

* Enhancement - Add IDs to graph resource logging: [#6593](https://github.com/owncloud/ocis/pull/6593)

   Graph access logs were unsuable as they didn't contain IDs to match them to a request

   https://github.com/owncloud/ocis/pull/6593

* Enhancement - Add permissions to report: [#6528](https://github.com/owncloud/ocis/pull/6528)

   The webdav REPORT endpoint only returned permissions for personal spaces and shares. Now also
   for project spaces.

   https://github.com/owncloud/ocis/pull/6528

* Enhancement - Add more metadata to the remote item: [#6300](https://github.com/owncloud/ocis/pull/6300)

   We added the drive alias, the space name and the relative path to the remote item. This is needed
   to resolve shared files directly on the source space.

   https://github.com/owncloud/ocis/pull/6300

* Enhancement - We added the storage id to the audit log for spaces: [#6548](https://github.com/owncloud/ocis/pull/6548)

   We added the storage id to the audit log for spaces

   https://github.com/owncloud/ocis/issues/3753
   https://github.com/owncloud/ocis/pull/6548

* Enhancement - Make the post logout redirect uri configurable: [#6583](https://github.com/owncloud/ocis/pull/6583)

   We added a config option to change the redirect uri after the logout action of the web client.

   https://github.com/owncloud/ocis/issues/6536
   https://github.com/owncloud/ocis/pull/6583

* Enhancement - Make the app provider service name configurable: [#6482](https://github.com/owncloud/ocis/pull/6482)

   We needed to make the service name of the app provider configurable. This needs to be changed
   when using more than one app provider. Each of them needs be found by a unique service name.
   Possible examples are: `app-provider-collabora`, `app-provider-onlyoffice`,
   `app-provider-office365`.

   https://github.com/owncloud/ocis/pull/6482

* Enhancement - Add old & new values to audit logs: [#6537](https://github.com/owncloud/ocis/pull/6537)

   We have added old & new values to the audit logs We have added the missing events for role changes

   https://github.com/owncloud/ocis/pull/6537

* Enhancement - Update reva: [#6529](https://github.com/owncloud/ocis/pull/6529)

   Update reva to latest edge

  *   Bugfix [cs3org/reva#3963](https://github.com/cs3org/reva/pull/3963): Treesize interger overflows
  *   Bugfix [cs3org/reva#3943](https://github.com/cs3org/reva/pull/3943): When removing metadata always use correct database and table
  *   Bugfix [cs3org/reva#3978](https://github.com/cs3org/reva/pull/3978): Decomposedfs no longer os.Stats when reading node metadata
  *   Bugfix [cs3org/reva#3959](https://github.com/cs3org/reva/pull/3959): Drop unnecessary stat
  *   Bugfix [cs3org/reva#3948](https://github.com/cs3org/reva/pull/3948): Handle the bad request status
  *   Bugfix [cs3org/reva#3955](https://github.com/cs3org/reva/pull/3955): Fix panic
  *   Bugfix [cs3org/reva#3977](https://github.com/cs3org/reva/pull/3977): Prevent direct access to trash items
  *   Bugfix [cs3org/reva#3933](https://github.com/cs3org/reva/pull/3933): Concurrently invalidate mtime cache in jsoncs3 share manager
  *   Bugfix [cs3org/reva#3985](https://github.com/cs3org/reva/pull/3985): Reduce jsoncs3 lock congestion
  *   Bugfix [cs3org/reva#3960](https://github.com/cs3org/reva/pull/3960): Add trace span details
  *   Bugfix [cs3org/reva#3951](https://github.com/cs3org/reva/pull/3951): Link context in metadata client
  *   Bugfix [cs3org/reva#3950](https://github.com/cs3org/reva/pull/3950): Use plain otel tracing in metadata client
  *   Bugfix [cs3org/reva#3975](https://github.com/cs3org/reva/pull/3975): Decomposedfs now resolves the parent without an os.Stat
  *   Change [cs3org/reva#3947](https://github.com/cs3org/reva/pull/3947): Bump golangci-lint to 1.51.2
  *   Change [cs3org/reva#3945](https://github.com/cs3org/reva/pull/3945): Revert golangci-lint back to 1.50.1
  *   Enhancement [cs3org/reva#3966](https://github.com/cs3org/reva/pull/3966): Add space metadata to ocs shares list
  *   Enhancement [cs3org/reva#3953](https://github.com/cs3org/reva/pull/3953): Client selector pool
  *   Enhancement [cs3org/reva#3941](https://github.com/cs3org/reva/pull/3941): Adding tracing for jsoncs3
  *   Enhancement [cs3org/reva#3965](https://github.com/cs3org/reva/pull/3965): ResumePostprocessing Event
  *   Enhancement [cs3org/reva#3981](https://github.com/cs3org/reva/pull/3981): We have updated the UserFeatureChangedEvent to reflect value changes
  *   Enhancement [cs3org/reva#3986](https://github.com/cs3org/reva/pull/3986): Allow disabling wopi chat

   https://github.com/owncloud/ocis/pull/6529
   https://github.com/owncloud/ocis/pull/6544
   https://github.com/owncloud/ocis/pull/6507
   https://github.com/owncloud/ocis/pull/6572
   https://github.com/owncloud/ocis/pull/6590

* Enhancement - Use reva client selectors: [#6452](https://github.com/owncloud/ocis/pull/6452)

   Use reva client selectors instead of the static clients, this introduces the ocis service
   registry in reva. The service discovery now resolves reva services by name and the client
   selectors pick a random registered service node.

   https://github.com/owncloud/ocis/pull/6452
   https://github.com/cs3org/reva/pull/3939
   https://github.com/cs3org/reva/pull/3953

* Enhancement - Add companion URL config: [#6453](https://github.com/owncloud/ocis/pull/6453)

   Introduce a config to set the Uppy Companion URL via `WEB_OPTION_UPLOAD_COMPANION_URL`.

   https://github.com/owncloud/ocis/pull/6453

* Enhancement - Add imprint and privacy url config: [#6462](https://github.com/owncloud/ocis/pull/6462)

   Introduce a config to set the imprint and privacy url via `WEB_OPTION_IMPRINT_URL` and
   `WEB_OPTION_PRIVACY_URL`.

   https://github.com/owncloud/ocis/pull/6462

* Enhancement - Add logged out url config: [#6549](https://github.com/owncloud/ocis/pull/6549)

   Introduce a config to set the more button url on the access denied page in web via
   `WEB_OPTION_ACCESS_DENIED_HELP_URL`.

   https://github.com/owncloud/ocis/pull/6549

* Enhancement - Fix envvar defaults: [#6516](https://github.com/owncloud/ocis/pull/6516)

   Defaults for the envvar OCIS_LDAP_DISABLE_USER_MECHANISM were not used consistently,
   correct is `attribute`.

   https://github.com/owncloud/ocis/issues/6513
   https://github.com/owncloud/ocis/pull/6516

* Enhancement - Skip if the simulink is a directory: [#6574](https://github.com/owncloud/ocis/pull/6574)

   Skip the error if the simulink is pointed to a directory

   https://github.com/owncloud/ocis/issues/6567
   https://github.com/owncloud/ocis/pull/6574

* Enhancement - Fix the groupname validation: [#6490](https://github.com/owncloud/ocis/pull/6490)

   Fixed the ability to create a group with an empty name

   https://github.com/owncloud/ocis/issues/5050
   https://github.com/owncloud/ocis/pull/6490

* Enhancement - Fix the username validation: [#6437](https://github.com/owncloud/ocis/pull/6437)

   Fix the username validation when an admin update the user

   https://github.com/owncloud/ocis/issues/6436
   https://github.com/owncloud/ocis/pull/6437

* Enhancement - Move proxy to service tracerprovider: [#6591](https://github.com/owncloud/ocis/pull/6591)

   This moves the proxy to initialise a service tracer provider at service initialisation time,
   instead of using a package global tracer provider.

   https://github.com/owncloud/ocis/pull/6591

* Enhancement - Add functionality to retry postprocessing: [#6500](https://github.com/owncloud/ocis/pull/6500)

   Adds a ctl command to manually retry failed postprocessing on uploads

   https://github.com/owncloud/ocis/pull/6500

* Enhancement - Update go-micro kubernetes registry: [#6457](https://github.com/owncloud/ocis/pull/6457)

   https://github.com/owncloud/ocis/pull/6457
   https://github.com/go-micro/plugins/pull/114
   https://github.com/go-micro/plugins/pull/113

* Enhancement - Update web to v7.0.1: [#6470](https://github.com/owncloud/ocis/pull/6470)

   Tags: web

   We updated ownCloud Web to v7.0.1. Please refer to the changelog (linked) for details on the web
   release.

   ## Summary * Bugfix [owncloud/web#9153](https://github.com/owncloud/web/pull/9153):
   Reduce space preloading

   https://github.com/owncloud/ocis/pull/6470
   https://github.com/owncloud/web/releases/tag/v7.0.1

* Enhancement - Allow disabling wopi chat: [#6544](https://github.com/owncloud/ocis/pull/6544)

   Add a configreva for the new reva disable-chat feature

   https://github.com/owncloud/ocis/pull/6544
# Changelog for [3.0.0] (2023-06-06)

The following sections list the changes for 3.0.0.

[3.0.0]: https://github.com/owncloud/ocis/compare/v2.0.0...v3.0.0

## Summary

* Bugfix - Return 425 on Thumbnails: [#5300](https://github.com/owncloud/ocis/pull/5300)
* Bugfix - Allow selected updates on graph users: [#6233](https://github.com/owncloud/ocis/pull/6233)
* Bugfix - Disassociate users from deleted school: [#5343](https://github.com/owncloud/ocis/pull/5343)
* Bugfix - Fix error message when disabling users: [#6435](https://github.com/owncloud/ocis/pull/6435)
* Bugfix - Fix default role assignment for demo users: [#3432](https://github.com/owncloud/ocis/issues/3432)
* Bugfix - Empty exact list while searching for a sharee: [#6398](https://github.com/owncloud/ocis/pull/6398)
* Bugfix - Reduced default TTL of user and group caches in graph API: [#6320](https://github.com/owncloud/ocis/issues/6320)
* Bugfix - Fix so that PATCH requests for groups actually updates the group name: [#5949](https://github.com/owncloud/ocis/pull/5949)
* Bugfix - Use UUID attribute for computing "sub" claim in lico idp: [#904](https://github.com/owncloud/ocis/issues/904)
* Bugfix - Hide the existence of space when deleting/updating: [#5031](https://github.com/owncloud/ocis/issues/5031)
* Bugfix - Fix OIDC auth cache: [#5997](https://github.com/owncloud/ocis/pull/5997)
* Bugfix - Fix the empty string givenName attribute when creating user: [#5431](https://github.com/owncloud/ocis/issues/5431)
* Bugfix - Fix Postprocessing events: [#5269](https://github.com/owncloud/ocis/pull/5269)
* Bugfix - Fix Search reindexing performance regression: [#6085](https://github.com/owncloud/ocis/pull/6085)
* Bugfix - Fix Search tag indexing: [#5405](https://github.com/owncloud/ocis/pull/5405)
* Bugfix - Fix the wrong status code when appRoleAssignments is forbidden: [#6037](https://github.com/owncloud/ocis/issues/6037)
* Bugfix - Fix user type config for user provider: [#6027](https://github.com/owncloud/ocis/pull/6027)
* Bugfix - Fix userlog panic: [#6114](https://github.com/owncloud/ocis/pull/6114)
* Bugfix - Fix Logout Url config name: [#6227](https://github.com/owncloud/ocis/pull/6227)
* Bugfix - Add missing CORS config: [#5987](https://github.com/owncloud/ocis/pull/5987)
* Bugfix - Add missing response to blocked requests: [#6277](https://github.com/owncloud/ocis/pull/6277)
* Bugfix - Populate expanded properties: [#5421](https://github.com/owncloud/ocis/pull/5421)
* Bugfix - Add portrait thumbnail resolutions: [#5656](https://github.com/owncloud/ocis/pull/5656)
* Bugfix - Trace proxy middlewares: [#6313](https://github.com/owncloud/ocis/pull/6313)
* Bugfix - Update the default admin role: [#6310](https://github.com/owncloud/ocis/pull/6310)
* Bugfix - Fix authenticate headers for API requests: [#5992](https://github.com/owncloud/ocis/pull/5992)
* Change - Bump libregraph lico: [#5768](https://github.com/owncloud/ocis/pull/5768)
* Change - Updated Cache Configuration: [#5829](https://github.com/owncloud/ocis/pull/5829)
* Change - Remove the settings ui: [#5463](https://github.com/owncloud/ocis/pull/5463)
* Change - Do not share versions: [#5531](https://github.com/owncloud/ocis/pull/5531)
* Change - We renamed the guest role to user light: [#6456](https://github.com/owncloud/ocis/pull/6456)
* Enhancement - Add specific result to antivirus for debugging: [#6265](https://github.com/owncloud/ocis/pull/6265)
* Enhancement - Add debug server to audit: [#6178](https://github.com/owncloud/ocis/pull/6178)
* Enhancement - Add debug server to idm: [#6153](https://github.com/owncloud/ocis/pull/6153)
* Enhancement - Add debug server to postprocessing: [#6203](https://github.com/owncloud/ocis/pull/6203)
* Enhancement - Add debug server to userlog: [#6202](https://github.com/owncloud/ocis/pull/6202)
* Enhancement - Add 'ocis decomposedfs metadata' command: [#5858](https://github.com/owncloud/ocis/pull/5858)
* Enhancement - Add debug server to eventhistory: [#6204](https://github.com/owncloud/ocis/pull/6204)
* Enhancement - Add global env variable extractor: [#5164](https://github.com/owncloud/ocis/pull/5164)
* Enhancement - Add the email HTML templates: [#6147](https://github.com/owncloud/ocis/pull/6147)
* Enhancement - Open Debug endpoint for Notifications: [#5002](https://github.com/owncloud/ocis/issues/5002)
* Enhancement - Add MessageRichParameters: [#5927](https://github.com/owncloud/ocis/pull/5927)
* Enhancement - Add webfinger service: [#5373](https://github.com/owncloud/ocis/pull/5373)
* Enhancement - Async Postprocessing: [#5207](https://github.com/owncloud/ocis/pull/5207)
* Enhancement - Automate md creation: [#5901](https://github.com/owncloud/ocis/pull/5901)
* Enhancement - Add more logging to av service: [#5973](https://github.com/owncloud/ocis/pull/5973)
* Enhancement - Return Bad Request when requesting GDPR export for another user: [#6123](https://github.com/owncloud/ocis/pull/6123)
* Enhancement - Add endpoints to upload a custom logo: [#5735](https://github.com/owncloud/ocis/pull/5735)
* Enhancement - Bump go-ldap version: [#6004](https://github.com/owncloud/ocis/pull/6004)
* Enhancement - Bump libre-graph-api-go: [#5309](https://github.com/owncloud/ocis/pull/5309)
* Enhancement - Update Reva to version 2.14.0: [#6448](https://github.com/owncloud/ocis/pull/6448)
* Enhancement - Collect global envvars: [#5367](https://github.com/owncloud/ocis/pull/5367)
* Enhancement - Make the settings bundles part of the service config: [#5589](https://github.com/owncloud/ocis/pull/5589)
* Enhancement - Configure GRPC in ocs: [#6022](https://github.com/owncloud/ocis/pull/6022)
* Enhancement - Default LDAP write to true: [#6362](https://github.com/owncloud/ocis/pull/6362)
* Enhancement - Disable Notifications: [#6137](https://github.com/owncloud/ocis/pull/6137)
* Enhancement - Drive group permissions: [#5312](https://github.com/owncloud/ocis/pull/5312)
* Enhancement - Make the group members addition limit configurable: [#5357](https://github.com/owncloud/ocis/pull/5357)
* Enhancement - Allow username to be changed: [#5509](https://github.com/owncloud/ocis/pull/5509)
* Enhancement - Graph Drives IdentitySet displayName: [#5347](https://github.com/owncloud/ocis/pull/5347)
* Enhancement - Make the LDAP base DN for new groups configurable: [#5974](https://github.com/owncloud/ocis/pull/5974)
* Enhancement - Update to go 1.20 to use memlimit: [#5732](https://github.com/owncloud/ocis/pull/5732)
* Enhancement - Display surname and givenName attributes: [#5388](https://github.com/owncloud/ocis/pull/5388)
* Enhancement - Extended search: [#5221](https://github.com/owncloud/ocis/pull/5221)
* Enhancement - Resource tags: [#5227](https://github.com/owncloud/ocis/pull/5227)
* Enhancement - Allow users to be disabled: [#5588](https://github.com/owncloud/ocis/pull/5588)
* Enhancement - Web config additions: [#6032](https://github.com/owncloud/ocis/pull/6032)
* Enhancement - Eventhistory service: [#5600](https://github.com/owncloud/ocis/pull/5600)
* Enhancement - Expiration Notifications: [#5330](https://github.com/owncloud/ocis/pull/5330)
* Enhancement - Fix to prevent the email X-Site scripting: [#6429](https://github.com/owncloud/ocis/pull/6429)
* Enhancement - Fix preview or viewing of shared animated GIFs: [#6386](https://github.com/owncloud/ocis/pull/6386)
* Enhancement - Fix err when the user share the locked file: [#6358](https://github.com/owncloud/ocis/pull/6358)
* Enhancement - Add fulltextsearch capabilty: [#6366](https://github.com/owncloud/ocis/pull/6366)
* Enhancement - GDPR Export: [#6064](https://github.com/owncloud/ocis/pull/6064)
* Enhancement - Make graph/education API errors more consistent: [#5682](https://github.com/owncloud/ocis/pull/5682)
* Enhancement - Graph user capabilities: [#6339](https://github.com/owncloud/ocis/pull/6339)
* Enhancement - Configurable ID Cache: [#6353](https://github.com/owncloud/ocis/pull/6353)
* Enhancement - Add endpoint to list permissions: [#5594](https://github.com/owncloud/ocis/pull/5594)
* Enhancement - Notifications: [#6038](https://github.com/owncloud/ocis/pull/6038)
* Enhancement - Open Debug endpoint for Nats: [#5002](https://github.com/owncloud/ocis/issues/5002)
* Enhancement - No Notifications for own actions: [#5871](https://github.com/owncloud/ocis/pull/5871)
* Enhancement - Notify about policies: [#5912](https://github.com/owncloud/ocis/pull/5912)
* Enhancement - Add otlp tracing exporter: [#5132](https://github.com/owncloud/ocis/pull/5132)
* Enhancement - Add a capability for the Personal Data export: [#5984](https://github.com/owncloud/ocis/pull/5984)
* Enhancement - Introduce policies-service: [#5714](https://github.com/owncloud/ocis/pull/5714)
* Enhancement - Better config for postprocessing service: [#5457](https://github.com/owncloud/ocis/pull/5457)
* Enhancement - Add Store to `postprocessing`: [#6281](https://github.com/owncloud/ocis/pull/6281)
* Enhancement - Add config option to enforce passwords on public links: [#5848](https://github.com/owncloud/ocis/pull/5848)
* Enhancement - Add new permission for public links: [#5690](https://github.com/owncloud/ocis/pull/5690)
* Enhancement - Remove the email logo: [#6359](https://github.com/owncloud/ocis/issues/6359)
* Enhancement - Remove quota from share jails api responses: [#6309](https://github.com/owncloud/ocis/pull/6309)
* Enhancement - Rename permissions: [#3922](https://github.com/cs3org/reva/pull/3922)
* Enhancement - Added possibility to assign roles based on OIDC claims: [#6048](https://github.com/owncloud/ocis/pull/6048)
* Enhancement - Added option to configure default quota per role: [#5616](https://github.com/owncloud/ocis/pull/5616)
* Enhancement - Add optional services to the runtime: [#6071](https://github.com/owncloud/ocis/pull/6071)
* Enhancement - Add new SetProjectSpaceQuota permission: [#5660](https://github.com/owncloud/ocis/pull/5660)
* Enhancement - Add expiration to user and group shares: [#5389](https://github.com/owncloud/ocis/pull/5389)
* Enhancement - Space Management permissions: [#5441](https://github.com/owncloud/ocis/pull/5441)
* Enhancement - Cli to purge expired trash-bin items: [#5500](https://github.com/owncloud/ocis/pull/5500)
* Enhancement - Unify CA Cert envvars: [#6392](https://github.com/owncloud/ocis/pull/6392)
* Enhancement - Update web to v7.0.0-rc.37: [#6294](https://github.com/owncloud/ocis/pull/6294)
* Enhancement - Update web to v7.0.0-rc.38: [#6375](https://github.com/owncloud/ocis/pull/6375)
* Enhancement - Update web to v7.0.0: [#6438](https://github.com/owncloud/ocis/pull/6438)
* Enhancement - Use Accept-Language Header: [#5918](https://github.com/owncloud/ocis/pull/5918)
* Enhancement - Use gotext master: [#5867](https://github.com/owncloud/ocis/pull/5867)
* Enhancement - Userlog: [#5699](https://github.com/owncloud/ocis/pull/5699)
* Enhancement - Userlog Service: [#5610](https://github.com/owncloud/ocis/pull/5610)
* Enhancement - Determine the users language to translate via Transifex: [#6089](https://github.com/owncloud/ocis/pull/6089)
* Enhancement - Web options configuration: [#6188](https://github.com/owncloud/ocis/pull/6188)

## Details

* Bugfix - Return 425 on Thumbnails: [#5300](https://github.com/owncloud/ocis/pull/5300)

   Return `425` on thumbnails `GET` when file is processing. Pass `425` also through webdav
   endpoint

   https://github.com/owncloud/ocis/pull/5300

* Bugfix - Allow selected updates on graph users: [#6233](https://github.com/owncloud/ocis/pull/6233)

   We are now allowing a couple of update request to complete even if
   GRAPH_LDAP_SERVER_WRITE_ENABLED=false:

  *   When using a group to disable users (OCIS_LDAP_DISABLE_USER_MECHANISM=group) updates to the accountEnabled property of a user will be allowed
  *   When a distinct base dn for new groups is configured ( GRAPH_LDAP_GROUP_CREATE_BASE_DN is set to a different value than GRAPH_LDAP_GROUP_BASE_DN), allow the creation/update of local groups.

   https://github.com/owncloud/ocis/pull/6233

* Bugfix - Disassociate users from deleted school: [#5343](https://github.com/owncloud/ocis/pull/5343)

   When a school is deleted, users should be disassociated from it.

   https://github.com/owncloud/ocis/issues/5246
   https://github.com/owncloud/ocis/pull/5343

* Bugfix - Fix error message when disabling users: [#6435](https://github.com/owncloud/ocis/pull/6435)

   When we disable users by adding them to a group we do not need to update the user entry.

   https://github.com/owncloud/ocis/pull/6435

* Bugfix - Fix default role assignment for demo users: [#3432](https://github.com/owncloud/ocis/issues/3432)

   The roles-assignments for demo users where duplicated with every restart of the settings
   service.

   https://github.com/owncloud/ocis/issues/3432

* Bugfix - Empty exact list while searching for a sharee: [#6398](https://github.com/owncloud/ocis/pull/6398)

   We fixed a bug in the sharing api, it always returns an empty exact list while searching for a
   sharee

   https://github.com/owncloud/ocis/issues/4265
   https://github.com/owncloud/ocis/pull/6398
   https://github.com/cs3org/reva/pull/3877

* Bugfix - Reduced default TTL of user and group caches in graph API: [#6320](https://github.com/owncloud/ocis/issues/6320)

   We reduced the default TTL of the cache for user and group information on the /drives endpoints
   to 60 seconds. This fixes in issue where outdated information was show on the spaces list for a
   very long time.

   https://github.com/owncloud/ocis/issues/6320

* Bugfix - Fix so that PATCH requests for groups actually updates the group name: [#5949](https://github.com/owncloud/ocis/pull/5949)

   https://github.com/owncloud/ocis/pull/5949

* Bugfix - Use UUID attribute for computing "sub" claim in lico idp: [#904](https://github.com/owncloud/ocis/issues/904)

   By default the LDAP backend for lico uses the User DN for computing the "sub" claim of a user. This
   caused the "sub" claim to stay the same even if a user was deleted and recreated (and go a new UUID
   assgined with that). We now use the user's unique id (`owncloudUUID` by default) for computing
   the `sub` claim. So that user's recreated with the same name will be treated as different users
   by the IDP.

   https://github.com/owncloud/ocis/issues/904
   https://github.com/owncloud/ocis/pull/6326
   https://github.com/owncloud/ocis/pull/6338
   https://github.com/owncloud/ocis/pull/6420

* Bugfix - Hide the existence of space when deleting/updating: [#5031](https://github.com/owncloud/ocis/issues/5031)

   The "code": "notAllowed" changed to "code": "itemNotFound"

   https://github.com/owncloud/ocis/issues/5031
   https://github.com/owncloud/ocis/pull/6220

* Bugfix - Fix OIDC auth cache: [#5997](https://github.com/owncloud/ocis/pull/5997)

   We've fixed an issue rendering the OIDC auth cache useless.

   https://github.com/owncloud/ocis/pull/5997

* Bugfix - Fix the empty string givenName attribute when creating user: [#5431](https://github.com/owncloud/ocis/issues/5431)

   Omitempty givenName attribute when creating user

   https://github.com/owncloud/ocis/issues/5431
   https://github.com/owncloud/ocis/pull/6259

* Bugfix - Fix Postprocessing events: [#5269](https://github.com/owncloud/ocis/pull/5269)

   Postprocessing service did not want to play with non-tls events. That is fixed now

   https://github.com/owncloud/ocis/pull/5269

* Bugfix - Fix Search reindexing performance regression: [#6085](https://github.com/owncloud/ocis/pull/6085)

   We've fixed a regression in the search service reindexing step, causing the whole space to be
   reindexed instead of just the changed resources.

   https://github.com/owncloud/ocis/pull/6085

* Bugfix - Fix Search tag indexing: [#5405](https://github.com/owncloud/ocis/pull/5405)

   We've fixed an issue where search is not able to index tags for space resources.

   https://github.com/owncloud/ocis/pull/5405

* Bugfix - Fix the wrong status code when appRoleAssignments is forbidden: [#6037](https://github.com/owncloud/ocis/issues/6037)

   Fix the wrong status code when appRoleAssignments is forbidden in the
   CreateAppRoleAssignment and DeleteAppRoleAssignment methods.

   https://github.com/owncloud/ocis/issues/6037
   https://github.com/owncloud/ocis/pull/6276

* Bugfix - Fix user type config for user provider: [#6027](https://github.com/owncloud/ocis/pull/6027)

   We needed to provide a default value for the user type property in the user provider.

   https://github.com/owncloud/ocis/pull/6027

* Bugfix - Fix userlog panic: [#6114](https://github.com/owncloud/ocis/pull/6114)

   Userlog services paniced because of `nil` ctx. That is fixed now

   https://github.com/owncloud/ocis/pull/6114

* Bugfix - Fix Logout Url config name: [#6227](https://github.com/owncloud/ocis/pull/6227)

   We fixed the yaml and json name of the logout url option.

   https://github.com/owncloud/ocis/pull/6227

* Bugfix - Add missing CORS config: [#5987](https://github.com/owncloud/ocis/pull/5987)

   The graph, userlog and ocdav services had no CORS config options.

   https://github.com/owncloud/ocis/pull/5987

* Bugfix - Add missing response to blocked requests: [#6277](https://github.com/owncloud/ocis/pull/6277)

   We added the missing response body to requests which were blocked by the policy engine.

   https://github.com/owncloud/ocis/pull/6277

* Bugfix - Populate expanded properties: [#5421](https://github.com/owncloud/ocis/pull/5421)

   We now return an empty array when an expanded relation has no entries. This makes consuming the
   responses a little easier.

   https://github.com/owncloud/ocis/issues/5419
   https://github.com/owncloud/ocis/pull/5421
   https://github.com/owncloud/ocis/pull/5426

* Bugfix - Add portrait thumbnail resolutions: [#5656](https://github.com/owncloud/ocis/pull/5656)

   Add portrait-orientation resolutions to the thumbnail service's default configuration.
   This prevents portrait photos from being heavily cropped into landscape resolutions in the
   web viewer.

   https://github.com/owncloud/ocis/pull/5656

* Bugfix - Trace proxy middlewares: [#6313](https://github.com/owncloud/ocis/pull/6313)

   We moved trace initialization to an early middleware to also trace requests made by other proxy
   middlewares.

   https://github.com/owncloud/ocis/pull/6313

* Bugfix - Update the default admin role: [#6310](https://github.com/owncloud/ocis/pull/6310)

   The admin role was missing two permissions. We added them to make the space admin role a subset of
   the admin role. This matches better with the default user expectations.

   https://github.com/owncloud/ocis/pull/6310

* Bugfix - Fix authenticate headers for API requests: [#5992](https://github.com/owncloud/ocis/pull/5992)

   We changed the www-authenticate header which should not be sent when the `XMLHttpRequest`
   header is set.

   https://github.com/owncloud/ocis/issues/5986
   https://github.com/owncloud/ocis/pull/5992

* Change - Bump libregraph lico: [#5768](https://github.com/owncloud/ocis/pull/5768)

   We updated lico to the latest version * Update to 0.59.4 - upstream dropped the kc and cookie
   backends

   https://github.com/owncloud/ocis/pull/5768

* Change - Updated Cache Configuration: [#5829](https://github.com/owncloud/ocis/pull/5829)

   We updated all cache related environment vars to more closely follow the go micro naming
   pattern: - `{service}_CACHE_STORE_TYPE` becomes `{service}_CACHE_STORE` or
   `{service}_PERSISTENT_STORE` - `{service}_CACHE_STORE_ADDRESS(ES)` becomes
   `{service}_CACHE_STORE_NODES` - The `mem` store implementation name changes to `memory` -
   In yaml files the cache `type` becomes `store` We introduced `redis-sentinel` as a store
   implementation.

   https://github.com/owncloud/ocis/pull/5829

* Change - Remove the settings ui: [#5463](https://github.com/owncloud/ocis/pull/5463)

   With ownCloud Web having transitioned to Vue 3 recently, we would have had to port the settings
   ui as well. The decision was made to discontinue the settings ui instead. As a result all traces
   of the settings ui have been removed.

   The only user facing setting that ever existed in the settings service is now integrated into
   the `account` page of ownCloud Web (click on top right user menu, then on your username to reach
   the account page).

   https://github.com/owncloud/ocis/pull/5463

* Change - Do not share versions: [#5531](https://github.com/owncloud/ocis/pull/5531)

   We changed the default behavior of shares: Share receivers have no access to versions. People
   in spaces with the "Editor" or "Manager" role can still see versions and work with them.

   https://github.com/owncloud/ocis/pull/5531

* Change - We renamed the guest role to user light: [#6456](https://github.com/owncloud/ocis/pull/6456)

   We needed to rename the "Guest" role to "User Light" because the naming was creating
   confusions. The roles are not bound to a user type.

   https://github.com/owncloud/ocis/issues/6058
   https://github.com/owncloud/ocis/pull/6456

* Enhancement - Add specific result to antivirus for debugging: [#6265](https://github.com/owncloud/ocis/pull/6265)

   We added the ability to define a specific result for the virus scanner via env-var
   (ANTIVIRUS_DEBUG_SCAN_OUTCOME)

   https://github.com/owncloud/ocis/pull/6265

* Enhancement - Add debug server to audit: [#6178](https://github.com/owncloud/ocis/pull/6178)

   We added a debug server to audit.

   https://github.com/owncloud/ocis/issues/5002
   https://github.com/owncloud/ocis/pull/6178

* Enhancement - Add debug server to idm: [#6153](https://github.com/owncloud/ocis/pull/6153)

   We added a debug server to idm.

   https://github.com/owncloud/ocis/issues/5003
   https://github.com/owncloud/ocis/pull/6153

* Enhancement - Add debug server to postprocessing: [#6203](https://github.com/owncloud/ocis/pull/6203)

   We added a debug server to postprocessing.

   https://github.com/owncloud/ocis/issues/5002
   https://github.com/owncloud/ocis/pull/6203

* Enhancement - Add debug server to userlog: [#6202](https://github.com/owncloud/ocis/pull/6202)

   We added a debug server to userlog.

   https://github.com/owncloud/ocis/issues/5002
   https://github.com/owncloud/ocis/pull/6202

* Enhancement - Add 'ocis decomposedfs metadata' command: [#5858](https://github.com/owncloud/ocis/pull/5858)

   We added a 'ocis decomposedfs metadata' command for inspecting and manipulating node
   metadata.

   https://github.com/owncloud/ocis/pull/5858

* Enhancement - Add debug server to eventhistory: [#6204](https://github.com/owncloud/ocis/pull/6204)

   We added a debug server to eventhistory.

   https://github.com/owncloud/ocis/issues/5002
   https://github.com/owncloud/ocis/pull/6204

* Enhancement - Add global env variable extractor: [#5164](https://github.com/owncloud/ocis/pull/5164)

   We have added a little tool that will extract global env vars, that are loaded only through
   os.Getenv for documentation purposes

   https://github.com/owncloud/ocis/issues/4916
   https://github.com/owncloud/ocis/pull/5164

* Enhancement - Add the email HTML templates: [#6147](https://github.com/owncloud/ocis/pull/6147)

   Add the email HTML templates

   https://github.com/owncloud/ocis/issues/6146
   https://github.com/owncloud/ocis/pull/6147

* Enhancement - Open Debug endpoint for Notifications: [#5002](https://github.com/owncloud/ocis/issues/5002)

   We added a debug server to the notifications service

   https://github.com/owncloud/ocis/issues/5002
   https://github.com/owncloud/ocis/pull/6155

* Enhancement - Add MessageRichParameters: [#5927](https://github.com/owncloud/ocis/pull/5927)

   Adds the messageRichParameters to virus and policies notifications

   https://github.com/owncloud/ocis/pull/5927

* Enhancement - Add webfinger service: [#5373](https://github.com/owncloud/ocis/pull/5373)

   Adds a webfinger service to redirect ocis clients

   https://github.com/owncloud/ocis/issues/6102
   https://github.com/owncloud/ocis/pull/5373
   https://github.com/owncloud/ocis/pull/6110

* Enhancement - Async Postprocessing: [#5207](https://github.com/owncloud/ocis/pull/5207)

   Provides functionality for async postprocessing. This will allow the system to do the
   postprocessing (virusscan, copying of bytes to their final destination, ...) asynchronous
   to the users request. Major change when active.

   https://github.com/owncloud/ocis/pull/5207

* Enhancement - Automate md creation: [#5901](https://github.com/owncloud/ocis/pull/5901)

   Automatically create `_index.md` files from the services `README.md`

   https://github.com/owncloud/ocis/pull/5901

* Enhancement - Add more logging to av service: [#5973](https://github.com/owncloud/ocis/pull/5973)

   We need more debug logging in some situations to understand the state of a virus scan.

   https://github.com/owncloud/ocis/pull/5973

* Enhancement - Return Bad Request when requesting GDPR export for another user: [#6123](https://github.com/owncloud/ocis/pull/6123)

   This is an enhancement, not security related as the requested uid is never used

   https://github.com/owncloud/ocis/pull/6123

* Enhancement - Add endpoints to upload a custom logo: [#5735](https://github.com/owncloud/ocis/pull/5735)

   Added endpoints to upload and reset custom logos. The files are stored under the
   `WEB_ASSET_PATH` which defaults to `$OCIS_BASE_DATA_PATH/web/assets`.

   https://github.com/owncloud/ocis/pull/5735
   https://github.com/owncloud/ocis/pull/5559

* Enhancement - Bump go-ldap version: [#6004](https://github.com/owncloud/ocis/pull/6004)

   Use master version of go-ldap to get rid of nasty `=` bug. See
   https://github.com/go-ldap/ldap/issues/416

   https://github.com/owncloud/ocis/pull/6004

* Enhancement - Bump libre-graph-api-go: [#5309](https://github.com/owncloud/ocis/pull/5309)

   We fixed a couple of issues in libre-graph-api-go package.

  * rename drive permission grantedTo to grantedToIdentities to be ms graph spec compatible.
  * drive.name is a required property now.
  * add group property to the identitySet.

   https://github.com/owncloud/ocis/pull/5309
   https://github.com/owncloud/ocis/pull/5312

* Enhancement - Update Reva to version 2.14.0: [#6448](https://github.com/owncloud/ocis/pull/6448)

   Changelog for reva 2.14.0 (2023-06-05) =======================================

  *   Bugfix [cs3org/reva#3919](https://github.com/cs3org/reva/pull/3919): We added missing timestamps to events
  *   Bugfix [cs3org/reva#3911](https://github.com/cs3org/reva/pull/3911): Clean IDCache properly
  *   Bugfix [cs3org/reva#3896](https://github.com/cs3org/reva/pull/3896): Do not lose old revisions when overwriting a file during copy
  *   Bugfix [cs3org/reva#3918](https://github.com/cs3org/reva/pull/3918): Dont enumerate users
  *   Bugfix [cs3org/reva#3902](https://github.com/cs3org/reva/pull/3902): Do not try to use the cache for empty node
  *   Bugfix [cs3org/reva#3877](https://github.com/cs3org/reva/pull/3877): Empty exact list while searching for a sharee
  *   Bugfix [cs3org/reva#3906](https://github.com/cs3org/reva/pull/3906): Fix preflight requests
  *   Bugfix [cs3org/reva#3934](https://github.com/cs3org/reva/pull/3934): Fix the space editor permissions
  *   Bugfix [cs3org/reva#3899](https://github.com/cs3org/reva/pull/3899): Harden uploads
  *   Bugfix [cs3org/reva#3917](https://github.com/cs3org/reva/pull/3917): Prevent last space manager from leaving
  *   Bugfix [cs3org/reva#3866](https://github.com/cs3org/reva/pull/3866): Fix public link lookup performance
  *   Bugfix [cs3org/reva#3904](https://github.com/cs3org/reva/pull/3904): Improve performance of directory listings
  *   Enhancement [cs3org/reva#3893](https://github.com/cs3org/reva/pull/3893): Cleanup Space Delete permissions
  *   Enhancement [cs3org/reva#3894](https://github.com/cs3org/reva/pull/3894): Fix err when the user share the locked file
  *   Enhancement [cs3org/reva#3913](https://github.com/cs3org/reva/pull/3913): Introduce FullTextSearch Capability
  *   Enhancement [cs3org/reva#3898](https://github.com/cs3org/reva/pull/3898): Add Graph User capabilities
  *   Enhancement [cs3org/reva#3496](https://github.com/cs3org/reva/pull/3496): Add otlp tracing exporter
  *   Enhancement [cs3org/reva#3922](https://github.com/cs3org/reva/pull/3922): Rename permissions

   Changelog for reva 2.13.3 (2023-05-17) =======================================

  *   Bugfix [cs3org/reva#3890](https://github.com/cs3org/reva/pull/3890): Bring back public link sharing of project space roots
  *   Bugfix [cs3org/reva#3888](https://github.com/cs3org/reva/pull/3888): We fixed a bug that unnecessarily fetched all members of a group
  *   Bugfix [cs3org/reva#3886](https://github.com/cs3org/reva/pull/3886): Decomposedfs no longer deadlocks when cache is disabled
  *   Bugfix [cs3org/reva#3892](https://github.com/cs3org/reva/pull/3892): Fix public links
  *   Bugfix [cs3org/reva#3876](https://github.com/cs3org/reva/pull/3876): Remove go-micro/store/redis specific workaround
  *   Bugfix [cs3org/reva#3889](https://github.com/cs3org/reva/pull/3889): Update space root mtime when changing space metadata
  *   Bugfix [cs3org/reva#3836](https://github.com/cs3org/reva/pull/3836): Fix spaceID in the decomposedFS
  *   Bugfix [cs3org/reva#3867](https://github.com/cs3org/reva/pull/3867): Restore last version after positive result
  *   Bugfix [cs3org/reva#3849](https://github.com/cs3org/reva/pull/3849): Prevent sharing space roots and personal spaces
  *   Enhancement [cs3org/reva#3865](https://github.com/cs3org/reva/pull/3865): Remove unneccessary code from gateway
  *   Enhancement [cs3org/reva#3895](https://github.com/cs3org/reva/pull/3895): Add missing expiry date to shares

   Changelog for reva 2.13.2 (2023-05-08) =======================================

  *   Bugfix [cs3org/reva#3845](https://github.com/cs3org/reva/pull/3845): Fix propagation
  *   Bugfix [cs3org/reva#3856](https://github.com/cs3org/reva/pull/3856): Fix response code
  *   Bugfix [cs3org/reva#3857](https://github.com/cs3org/reva/pull/3857): Fix trashbin purge

   Changelog for reva 2.13.1 (2023-05-03) =======================================

  *   Bugfix [cs3org/reva#3843](https://github.com/cs3org/reva/pull/3843): Allow scope check to impersonate space owners

   Changelog for reva 2.13.0 (2023-05-02) =======================================

  *   Bugfix [cs3org/reva#3570](https://github.com/cs3org/reva/pull/3570): Return 425 on HEAD
  *   Bugfix [cs3org/reva#3830](https://github.com/cs3org/reva/pull/3830): Be more robust when logging errors
  *   Bugfix [cs3org/reva#3815](https://github.com/cs3org/reva/pull/3815): Bump micro redis store
  *   Bugfix [cs3org/reva#3596](https://github.com/cs3org/reva/pull/3596): Cache CreateHome calls
  *   Bugfix [cs3org/reva#3823](https://github.com/cs3org/reva/pull/3823): Deny correctly in decomposedfs
  *   Bugfix [cs3org/reva#3826](https://github.com/cs3org/reva/pull/3826): Add by group index to decomposedfs
  *   Bugfix [cs3org/reva#3618](https://github.com/cs3org/reva/pull/3618): Drain body on failed put
  *   Bugfix [cs3org/reva#3685](https://github.com/cs3org/reva/pull/3685): Send fileid on copy
  *   Bugfix [cs3org/reva#3688](https://github.com/cs3org/reva/pull/3688): Return 425 on GET
  *   Bugfix [cs3org/reva#3755](https://github.com/cs3org/reva/pull/3755): Fix app provider language validation
  *   Bugfix [cs3org/reva#3800](https://github.com/cs3org/reva/pull/3800): Fix building for freebsd
  *   Bugfix [cs3org/reva#3700](https://github.com/cs3org/reva/pull/3700): Fix caching
  *   Bugfix [cs3org/reva#3535](https://github.com/cs3org/reva/pull/3535): Fix ceph driver storage fs implementation
  *   Bugfix [cs3org/reva#3764](https://github.com/cs3org/reva/pull/3764): Fix missing CORS config in ocdav service
  *   Bugfix [cs3org/reva#3710](https://github.com/cs3org/reva/pull/3710): Fix error when try to delete space without permission
  *   Bugfix [cs3org/reva#3822](https://github.com/cs3org/reva/pull/3822): Fix deleting spaces
  *   Bugfix [cs3org/reva#3718](https://github.com/cs3org/reva/pull/3718): Fix revad-eos docker image which was failing to build
  *   Bugfix [cs3org/reva#3559](https://github.com/cs3org/reva/pull/3559): Fix build on freebsd
  *   Bugfix [cs3org/reva#3696](https://github.com/cs3org/reva/pull/3696): Fix ldap filters when checking for enabled users
  *   Bugfix [cs3org/reva#3767](https://github.com/cs3org/reva/pull/3767): Decode binary UUID when looking up a users group memberships
  *   Bugfix [cs3org/reva#3741](https://github.com/cs3org/reva/pull/3741): Fix listing shares to multiple groups
  *   Bugfix [cs3org/reva#3834](https://github.com/cs3org/reva/pull/3834): Return correct error during MKCOL
  *   Bugfix [cs3org/reva#3841](https://github.com/cs3org/reva/pull/3841): Fix nil pointer and improve logging
  *   Bugfix [cs3org/reva#3831](https://github.com/cs3org/reva/pull/3831): Ignore 'null' mtime on tus upload
  *   Bugfix [cs3org/reva#3758](https://github.com/cs3org/reva/pull/3758): Fix public links with enforced password
  *   Bugfix [cs3org/reva#3814](https://github.com/cs3org/reva/pull/3814): Fix stat cache access
  *   Bugfix [cs3org/reva#3650](https://github.com/cs3org/reva/pull/3650): FreeBSD xattr support
  *   Bugfix [cs3org/reva#3827](https://github.com/cs3org/reva/pull/3827): Initialize user cache for decomposedfs
  *   Bugfix [cs3org/reva#3818](https://github.com/cs3org/reva/pull/3818): Invalidate cache when deleting space
  *   Bugfix [cs3org/reva#3812](https://github.com/cs3org/reva/pull/3812): Filemetadata Cache now deletes keys without listing them first
  *   Bugfix [cs3org/reva#3817](https://github.com/cs3org/reva/pull/3817): Pipeline cache deletes
  *   Bugfix [cs3org/reva#3711](https://github.com/cs3org/reva/pull/3711): Replace ini metadata backend by messagepack backend
  *   Bugfix [cs3org/reva#3828](https://github.com/cs3org/reva/pull/3828): Send quota when listing spaces in decomposedfs
  *   Bugfix [cs3org/reva#3681](https://github.com/cs3org/reva/pull/3681): Fix etag of "empty" shares jail
  *   Bugfix [cs3org/reva#3748](https://github.com/cs3org/reva/pull/3748): Prevent service from panicking
  *   Bugfix [cs3org/reva#3816](https://github.com/cs3org/reva/pull/3816): Write Metadata once
  *   Change [cs3org/reva#3641](https://github.com/cs3org/reva/pull/3641): Hide file versions for share receivers
  *   Change [cs3org/reva#3820](https://github.com/cs3org/reva/pull/3820): Streamline stores
  *   Enhancement [cs3org/reva#3732](https://github.com/cs3org/reva/pull/3732): Make method for detecting the metadata backend public
  *   Enhancement [cs3org/reva#3789](https://github.com/cs3org/reva/pull/3789): Add capabilities indicating if user attributes are read-only
  *   Enhancement [cs3org/reva#3792](https://github.com/cs3org/reva/pull/3792): Add a prometheus gauge to keep track of active uploads and downloads
  *   Enhancement [cs3org/reva#3637](https://github.com/cs3org/reva/pull/3637): Add an ID to each events
  *   Enhancement [cs3org/reva#3704](https://github.com/cs3org/reva/pull/3704): Add more information to events
  *   Enhancement [cs3org/reva#3744](https://github.com/cs3org/reva/pull/3744): Add LDAP user type attribute
  *   Enhancement [cs3org/reva#3806](https://github.com/cs3org/reva/pull/3806): Decomposedfs now supports filtering spaces by owner
  *   Enhancement [cs3org/reva#3730](https://github.com/cs3org/reva/pull/3730): Antivirus
  *   Enhancement [cs3org/reva#3531](https://github.com/cs3org/reva/pull/3531): Async Postprocessing
  *   Enhancement [cs3org/reva#3571](https://github.com/cs3org/reva/pull/3571): Async Upload Improvements
  *   Enhancement [cs3org/reva#3801](https://github.com/cs3org/reva/pull/3801): Cache node ids
  *   Enhancement [cs3org/reva#3690](https://github.com/cs3org/reva/pull/3690): Check set project space quota permission
  *   Enhancement [cs3org/reva#3686](https://github.com/cs3org/reva/pull/3686): User disabling functionality
  *   Enhancement [cs3org/reva#3505](https://github.com/cs3org/reva/pull/3505): Fix eosgrpc package
  *   Enhancement [cs3org/reva#3575](https://github.com/cs3org/reva/pull/3575): Fix skip group grant index cleanup
  *   Enhancement [cs3org/reva#3564](https://github.com/cs3org/reva/pull/3564): Fix tag pkg
  *   Enhancement [cs3org/reva#3756](https://github.com/cs3org/reva/pull/3756): Prepare for GDPR export
  *   Enhancement [cs3org/reva#3612](https://github.com/cs3org/reva/pull/3612): Group feature changed event added
  *   Enhancement [cs3org/reva#3729](https://github.com/cs3org/reva/pull/3729): Improve decomposedfs performance, esp. with network fs/cache
  *   Enhancement [cs3org/reva#3697](https://github.com/cs3org/reva/pull/3697): Improve the ini file metadata backend
  *   Enhancement [cs3org/reva#3819](https://github.com/cs3org/reva/pull/3819): Allow creating internal links without permission
  *   Enhancement [cs3org/reva#3740](https://github.com/cs3org/reva/pull/3740): Limit concurrency in decomposedfs
  *   Enhancement [cs3org/reva#3569](https://github.com/cs3org/reva/pull/3569): Always list shares jail when listing spaces
  *   Enhancement [cs3org/reva#3788](https://github.com/cs3org/reva/pull/3788): Make resharing configurable
  *   Enhancement [cs3org/reva#3674](https://github.com/cs3org/reva/pull/3674): Introduce ini file based metadata backend
  *   Enhancement [cs3org/reva#3728](https://github.com/cs3org/reva/pull/3728): Automatically migrate file metadata from xattrs to messagepack
  *   Enhancement [cs3org/reva#3807](https://github.com/cs3org/reva/pull/3807): Name Validation
  *   Enhancement [cs3org/reva#3574](https://github.com/cs3org/reva/pull/3574): Opaque space group
  *   Enhancement [cs3org/reva#3598](https://github.com/cs3org/reva/pull/3598): Pass estream to Storage Providers
  *   Enhancement [cs3org/reva#3763](https://github.com/cs3org/reva/pull/3763): Add a capability for personal data export
  *   Enhancement [cs3org/reva#3577](https://github.com/cs3org/reva/pull/3577): Prepare for SSE
  *   Enhancement [cs3org/reva#3731](https://github.com/cs3org/reva/pull/3731): Add config option to enforce passwords on public links
  *   Enhancement [cs3org/reva#3693](https://github.com/cs3org/reva/pull/3693): Enforce the PublicLink.Write permission
  *   Enhancement [cs3org/reva#3497](https://github.com/cs3org/reva/pull/3497): Introduce owncloud 10 publiclink manager
  *   Enhancement [cs3org/reva#3714](https://github.com/cs3org/reva/pull/3714): Add global max quota option and quota for CreateHome
  *   Enhancement [cs3org/reva#3759](https://github.com/cs3org/reva/pull/3759): Set correct share type when listing shares
  *   Enhancement [cs3org/reva#3594](https://github.com/cs3org/reva/pull/3594): Add expiration to user and group shares
  *   Enhancement [cs3org/reva#3580](https://github.com/cs3org/reva/pull/3580): Share expired event
  *   Enhancement [cs3org/reva#3620](https://github.com/cs3org/reva/pull/3620): Allow a new ShareType `SpaceMembershipGroup`
  *   Enhancement [cs3org/reva#3609](https://github.com/cs3org/reva/pull/3609): Space Management Permissions
  *   Enhancement [cs3org/reva#3655](https://github.com/cs3org/reva/pull/3655): Add expiration date to space memberships
  *   Enhancement [cs3org/reva#3697](https://github.com/cs3org/reva/pull/3697): Add support for redis sentinel caches
  *   Enhancement [cs3org/reva#3552](https://github.com/cs3org/reva/pull/3552): Suppress tusd logs
  *   Enhancement [cs3org/reva#3555](https://github.com/cs3org/reva/pull/3555): Tags
  *   Enhancement [cs3org/reva#3785](https://github.com/cs3org/reva/pull/3785): Increase unit test coverage in the ocdav service
  *   Enhancement [cs3org/reva#3739](https://github.com/cs3org/reva/pull/3739): Try to rename uploaded files to their final position
  *   Enhancement [cs3org/reva#3610](https://github.com/cs3org/reva/pull/3610): Walk and log chi routes

   https://github.com/owncloud/ocis/pull/6448
   https://github.com/owncloud/ocis/pull/6447
   https://github.com/owncloud/ocis/pull/6381
   https://github.com/owncloud/ocis/pull/6305
   https://github.com/owncloud/ocis/pull/6339
   https://github.com/owncloud/ocis/pull/6205
   https://github.com/owncloud/ocis/pull/6186

* Enhancement - Collect global envvars: [#5367](https://github.com/owncloud/ocis/pull/5367)

   Compose a list of all envvars living in more than 1 service

   https://github.com/owncloud/ocis/pull/5367

* Enhancement - Make the settings bundles part of the service config: [#5589](https://github.com/owncloud/ocis/pull/5589)

   We added the settings bundles to the config. The default roles are still unchanged. You can now
   override the defaults by replacing the whole bundles list via json config files. The config
   file is loaded from a specified path which can be configured with `SETTINGS_BUNDLES_PATH`.

   https://github.com/owncloud/ocis/pull/5589
   https://github.com/owncloud/ocis/pull/5607

* Enhancement - Configure GRPC in ocs: [#6022](https://github.com/owncloud/ocis/pull/6022)

   Fixes a panic in ocs when running not in single binary

   https://github.com/owncloud/ocis/pull/6022

* Enhancement - Default LDAP write to true: [#6362](https://github.com/owncloud/ocis/pull/6362)

   Default `OCIS_LDAP_SERVER_WRITE_ENABLED` to true

   https://github.com/owncloud/ocis/pull/6362

* Enhancement - Disable Notifications: [#6137](https://github.com/owncloud/ocis/pull/6137)

   Introduce new setting to disable notifications

   https://github.com/owncloud/ocis/pull/6137

* Enhancement - Drive group permissions: [#5312](https://github.com/owncloud/ocis/pull/5312)

   We've updated the libregraph.Drive response to contain group permissions.

   https://github.com/owncloud/ocis/pull/5312

* Enhancement - Make the group members addition limit configurable: [#5357](https://github.com/owncloud/ocis/pull/5357)

   It's now possible to configure the limit of group members addition by PATCHing
   `/graph/v1.0/groups/{groupID}`. It still defaults to 20 as defined in the spec but it can be
   configured via `.graph.api.group_members_patch_limit` in `ocis.yaml` or via the
   `GRAPH_GROUP_MEMBERS_PATCH_LIMIT` environment variable.

   https://github.com/owncloud/ocis/issues/5262
   https://github.com/owncloud/ocis/pull/5357

* Enhancement - Allow username to be changed: [#5509](https://github.com/owncloud/ocis/pull/5509)

   When OnPremisesSamAccountName is present in a PATCH on `{apiRoot}/users/{userID}` it will
   change the username of the user. This also changes the references to this user in the groups.

   https://github.com/owncloud/ocis/issues/4988
   https://github.com/owncloud/ocis/pull/5509

* Enhancement - Graph Drives IdentitySet displayName: [#5347](https://github.com/owncloud/ocis/pull/5347)

   We've added the IdentitySet displayName property to the group and user sets for the graph
   drives endpoint. The values for groups and users get cached.

   https://github.com/owncloud/ocis/pull/5347
   https://github.com/owncloud/web/pull/8178

* Enhancement - Make the LDAP base DN for new groups configurable: [#5974](https://github.com/owncloud/ocis/pull/5974)

   The LDAP backend for the Graph service introduced a new config option for setting the Parent DN
   for new groups created via the `/groups/` endpoint. (`GRAPH_LDAP_GROUP_CREATE_BASE_DN`)

   It defaults to the value of `GRAPH_LDAP_GROUP_BASE_DN`. If set to a different value the
   `GRAPH_LDAP_GROUP_CREATE_BASE_DN` needs to be a subordinate DN of
   `GRAPH_LDAP_GROUP_BASE_DN`.

   All existing groups with a DN outside the `GRAPH_LDAP_GROUP_CREATE_BASE_DN` tree will be
   treated as read-only groups. So it is not possible to edit these groups.

   https://github.com/owncloud/ocis/pull/5974

* Enhancement - Update to go 1.20 to use memlimit: [#5732](https://github.com/owncloud/ocis/pull/5732)

   We updated to go 1.20 which allows setting GOMEMLIMIT, which we by default set to 0.9.

   https://github.com/owncloud/ocis/pull/5732

* Enhancement - Display surname and givenName attributes: [#5388](https://github.com/owncloud/ocis/pull/5388)

   When querying the graph API, the surname and givenName attributes are now displayed for users.

   https://github.com/owncloud/ocis/issues/5386
   https://github.com/owncloud/ocis/pull/5388

* Enhancement - Extended search: [#5221](https://github.com/owncloud/ocis/pull/5221)

   Provides multiple enhancement to the search implementation. * content extraction, search
   now supports apache tika to extract resource contents. * search engine, underlying search
   engine is swappable now. * event consumers, the number of event consumers can now be set, which
   improves the speed of the individual tasks

   https://github.com/owncloud/ocis/issues/5184
   https://github.com/owncloud/ocis/pull/5221

* Enhancement - Resource tags: [#5227](https://github.com/owncloud/ocis/pull/5227)

   We've added the ability to tag resources via the graph api. Tags can be added (put request) and
   removed (delete request) from a resource, a list of available tags can also be requested by
   sending a get request to the graph endpoint.

   https://github.com/owncloud/ocis/issues/5184
   https://github.com/owncloud/ocis/pull/5227
   https://github.com/owncloud/ocis/pull/5271

* Enhancement - Allow users to be disabled: [#5588](https://github.com/owncloud/ocis/pull/5588)

   By setting the `accountEnabled` property to `false` for a user via the graph API. Users can be
   disabled (i.e. they can no longer login)

   https://github.com/owncloud/ocis/pull/5588
   https://github.com/owncloud/ocis/pull/5620

* Enhancement - Web config additions: [#6032](https://github.com/owncloud/ocis/pull/6032)

   We've added config keys for defining additional css, scripts and translations for ownCloud
   Web.

   https://github.com/owncloud/ocis/pull/6032

* Enhancement - Eventhistory service: [#5600](https://github.com/owncloud/ocis/pull/5600)

   Introduces the `eventhistory` service. It is a service that stores events and provides a grpc
   API to retrieve them.

   https://github.com/owncloud/ocis/pull/5600

* Enhancement - Expiration Notifications: [#5330](https://github.com/owncloud/ocis/pull/5330)

   Send emails to the user informing that a share or a space membership expires.

   https://github.com/owncloud/ocis/pull/5330

* Enhancement - Fix to prevent the email X-Site scripting: [#6429](https://github.com/owncloud/ocis/pull/6429)

   Fix to prevent the email notification X-Site scripting

   https://github.com/owncloud/ocis/issues/6411
   https://github.com/owncloud/ocis/pull/6429

* Enhancement - Fix preview or viewing of shared animated GIFs: [#6386](https://github.com/owncloud/ocis/pull/6386)

   Fix preview or viewing of shared animated GIFs

   https://github.com/owncloud/ocis/issues/5418
   https://github.com/owncloud/ocis/pull/6386

* Enhancement - Fix err when the user share the locked file: [#6358](https://github.com/owncloud/ocis/pull/6358)

   Fix unexpected behavior when the user try to share the locked file

   https://github.com/owncloud/ocis/issues/6197
   https://github.com/owncloud/ocis/pull/6358

* Enhancement - Add fulltextsearch capabilty: [#6366](https://github.com/owncloud/ocis/pull/6366)

   It needs an extra envvar `FRONTEND_FULL_TEXT_SEARCH_ENABLED`

   https://github.com/owncloud/ocis/pull/6366

* Enhancement - GDPR Export: [#6064](https://github.com/owncloud/ocis/pull/6064)

   Adds an endpoint to collect all data that is related to a user

   https://github.com/owncloud/ocis/pull/6064
   https://github.com/owncloud/ocis/pull/5950

* Enhancement - Make graph/education API errors more consistent: [#5682](https://github.com/owncloud/ocis/pull/5682)

   Aligned the error messages when creating schools and classes fail and changed the response
   code from 500 to 409.

   https://github.com/owncloud/ocis/issues/5660
   https://github.com/owncloud/ocis/pull/5682

* Enhancement - Graph user capabilities: [#6339](https://github.com/owncloud/ocis/pull/6339)

   Adds capablities to show if users are writeable in LDAP so clients can block their specific
   fields

   https://github.com/owncloud/ocis/pull/6339

* Enhancement - Configurable ID Cache: [#6353](https://github.com/owncloud/ocis/pull/6353)

   Makes the integrated idcache (used to reduce reads from disc) configurable with the general
   cache envvars

   https://github.com/owncloud/ocis/pull/6353

* Enhancement - Add endpoint to list permissions: [#5594](https://github.com/owncloud/ocis/pull/5594)

   We added 'https://cloud.ocis.test/api/v0/settings/permissions-list' to retrieve all
   permissions of the logged in user.

   https://github.com/owncloud/ocis/pull/5594
   https://github.com/owncloud/ocis/pull/5571

* Enhancement - Notifications: [#6038](https://github.com/owncloud/ocis/pull/6038)

   Make Emails translatable via transifex The transifex translation add in to the email
   templates. The optional environment variable NOTIFICATIONS_TRANSLATION_PATH added to
   config. The optional global environment variable OCIS_TRANSLATION_PATH added to
   notifications and userlog config.

   https://github.com/owncloud/ocis/issues/6025
   https://github.com/owncloud/ocis/pull/6038

* Enhancement - Open Debug endpoint for Nats: [#5002](https://github.com/owncloud/ocis/issues/5002)

   We added a debug server to nats

   https://github.com/owncloud/ocis/issues/5002
   https://github.com/owncloud/ocis/pull/6139

* Enhancement - No Notifications for own actions: [#5871](https://github.com/owncloud/ocis/pull/5871)

   Don't send notifications on space events when the user has executed them herself.

   https://github.com/owncloud/ocis/pull/5871

* Enhancement - Notify about policies: [#5912](https://github.com/owncloud/ocis/pull/5912)

   Notify the user when a file was deleted due to policies (policies service)

   https://github.com/owncloud/ocis/pull/5912

* Enhancement - Add otlp tracing exporter: [#5132](https://github.com/owncloud/ocis/pull/5132)

   We can now configure otlp to send traces using the otlp exporter.

   https://github.com/owncloud/ocis/pull/5132
   https://github.com/cs3org/reva/pull/3496

* Enhancement - Add a capability for the Personal Data export: [#5984](https://github.com/owncloud/ocis/pull/5984)

   Adds a capability for the personal data export endpoint

   https://github.com/owncloud/ocis/pull/5984

* Enhancement - Introduce policies-service: [#5714](https://github.com/owncloud/ocis/pull/5714)

   Introduces policies service. The policies-service provides a new grpc api which can be used to
   return whether a requested operation is allowed or not. Open Policy Agent is used to determine
   the set of rules of what is permitted and what is not.

   2 further levels of authorization build on this:

  * Proxy Authorization
  * Event Authorization (needs async post-processing enabled)

   The simplest authorization layer is in the proxy, since every request is processed here, only
   simple decisions that can be processed quickly are made here, more complex queries such as file
   evaluation are explicitly excluded in this layer.

   The next layer is event-based as a pipeline step in asynchronous post-processing, since
   processing at this point is asynchronous, the operations there can also take longer and be more
   expensive, the bytes of a file can be examined here as an example.

   Since the base block is a grpc api, it is also possible to use it directly. The policies are
   written in the [rego query
   language](https://www.openpolicyagent.org/docs/latest/policy-language/).

   https://github.com/owncloud/ocis/issues/5580
   https://github.com/owncloud/ocis/pull/5714

* Enhancement - Better config for postprocessing service: [#5457](https://github.com/owncloud/ocis/pull/5457)

   The postprocessing service is now individually configurable. This is achieved by allowing a
   list of postprocessing steps that are processed in order of their appearance in the
   `POSTPROCESSING_STEPS` envvar.

   https://github.com/owncloud/ocis/pull/5457

* Enhancement - Add Store to `postprocessing`: [#6281](https://github.com/owncloud/ocis/pull/6281)

   Add a gomicro store for the postprocessing service. Needed to run multiple postprocessing
   instances

   https://github.com/owncloud/ocis/pull/6281

* Enhancement - Add config option to enforce passwords on public links: [#5848](https://github.com/owncloud/ocis/pull/5848)

   Added a new config option to enforce passwords on public links with "Uploader, Editor,
   Contributor" roles.

   The new options are: `OCIS_SHARING_PUBLIC_WRITEABLE_SHARE_MUST_HAVE_PASSWORD`,
   `SHARING_PUBLIC_WRITEABLE_SHARE_MUST_HAVE_PASSWORD` and
   `FRONTEND_OCS_PUBLIC_WRITEABLE_SHARE_MUST_HAVE_PASSWORD`. Check the docs on how to
   properly set them.

   https://github.com/owncloud/ocis/pull/5848
   https://github.com/owncloud/ocis/pull/5785
   https://github.com/owncloud/ocis/pull/5720

* Enhancement - Add new permission for public links: [#5690](https://github.com/owncloud/ocis/pull/5690)

   Added a new permission 'PublicLink.Write' to check if a user can create or update public links.

   https://github.com/owncloud/ocis/pull/5690

* Enhancement - Remove the email logo: [#6359](https://github.com/owncloud/ocis/issues/6359)

   Remove the email logo

   https://github.com/owncloud/ocis/issues/6359
   https://github.com/owncloud/ocis/pull/6361

* Enhancement - Remove quota from share jails api responses: [#6309](https://github.com/owncloud/ocis/pull/6309)

   We have removed the quota object from api responses for share jails, which would permanently
   show exceeded due to restrictions in the permission system.

   https://github.com/owncloud/ocis/issues/4472
   https://github.com/owncloud/ocis/pull/6309

* Enhancement - Rename permissions: [#3922](https://github.com/cs3org/reva/pull/3922)

   Rename permissions to be consistent and future proof

   https://github.com/cs3org/reva/pull/3922
   https://github.com/owncloud/ocis/pull/6418

* Enhancement - Added possibility to assign roles based on OIDC claims: [#6048](https://github.com/owncloud/ocis/pull/6048)

   OCIS can now be configured to update a user's role assignment from the values of a claim provided
   via the IDPs userinfo endpoint. The claim name and the mapping between claim values and ocis
   role name can be configured via the configuration of the proxy service. Example:

   ```yaml role_assignment: driver: oidc oidc_role_mapper: role_claim: ocisRoles
   role_mapping: - role_name: admin claim_value: myAdminRole - role_name: spaceadmin
   claim_value: mySpaceAdminRole - role_name: user claim_value: myUserRole - role_name:
   guest: claim_value: myGuestRole ```

   https://github.com/owncloud/ocis/pull/6048

* Enhancement - Added option to configure default quota per role: [#5616](https://github.com/owncloud/ocis/pull/5616)

   Admins can assign default quotas to users with certain roles by adding the following config to
   the `proxy.yaml`. E.g.: ``` role_quotas: d7beeea8-8ff4-406b-8fb6-ab2dd81e6b11: 2300000
   ```

   It maps a role ID to the quota in bytes.

   https://github.com/owncloud/ocis/pull/5616

* Enhancement - Add optional services to the runtime: [#6071](https://github.com/owncloud/ocis/pull/6071)

   Make it possible to start optional services in the ocis runtime. Instead of using
   `OCIS_RUN_SERVICES` to define all services we can now use `OCIS_ADD_RUN_SERVICES` to add a
   comma separated list of additional services which are not started in the single process by
   default.

   https://github.com/owncloud/ocis/pull/6071

* Enhancement - Add new SetProjectSpaceQuota permission: [#5660](https://github.com/owncloud/ocis/pull/5660)

   Additionally to `set-space-quota` for setting quota on personal spaces we now have
   `Drive.ReadWriteQuota.Project` for setting project spaces quota

   https://github.com/owncloud/ocis/pull/5660

* Enhancement - Add expiration to user and group shares: [#5389](https://github.com/owncloud/ocis/pull/5389)

   Added expiration to user and group shares.

   https://github.com/owncloud/ocis/pull/5389

* Enhancement - Space Management permissions: [#5441](https://github.com/owncloud/ocis/pull/5441)

   We added new space management permissions. `space-properties` will allow changing space
   properties (name, description, ...). `space-ability` will allow enabling and disabling
   spaces

   https://github.com/owncloud/ocis/pull/5441

* Enhancement - Cli to purge expired trash-bin items: [#5500](https://github.com/owncloud/ocis/pull/5500)

   Introduction of a new cli command to purge old trash-bin items. The command is part of the
   `storage-users` service and can be used as follows:

   `ocis storage-users trash-bin purge-expired`.

   The `purge-expired` command configuration is done in the `ocis`configuration or as usual by
   using environment variables.

   ENV `STORAGE_USERS_PURGE_TRASH_BIN_USER_ID` is used to obtain space trash-bin
   information and takes the system admin user as the default `OCIS_ADMIN_USER_ID`. It should be
   noted, that this is only set by default in the single binary. The command only considers spaces
   to which the user has access and delete permission.

   ENV `STORAGE_USERS_PURGE_TRASH_BIN_PERSONAL_DELETE_BEFORE` has a default value of `30
   days`, which means the command will delete all files older than `30 days`. The value is
   human-readable, valid values are `24h`, `60m`, `60s` etc. `0` is equivalent to disable and
   prevents the deletion of `personal space` trash-bin files.

   ENV `STORAGE_USERS_PURGE_TRASH_BIN_PROJECT_DELETE_BEFORE` has a default value of `30
   days`, which means the command will delete all files older than `30 days`. The value is
   human-readable, valid values are `24h`, `60m`, `60s` etc. `0` is equivalent to disable and
   prevents the deletion of `project space` trash-bin files.

   Likewise, only spaces of the type `project` and `personal` are taken into account. Spaces of
   type `virtual`, for example, are ignored.

   https://github.com/owncloud/ocis/issues/5499
   https://github.com/owncloud/ocis/pull/5500

* Enhancement - Unify CA Cert envvars: [#6392](https://github.com/owncloud/ocis/pull/6392)

   Introduce a global `OCIS_EVENTS_TLS_ROOT_CA_CERTIFICATE` to avoid needing to configure
   all `{SERVICENAME}_EVENTS_TLS_ROOT_CA_CERTIFICATE` envvars

   https://github.com/owncloud/ocis/pull/6392

* Enhancement - Update web to v7.0.0-rc.37: [#6294](https://github.com/owncloud/ocis/pull/6294)

   Tags: web

   We updated ownCloud Web to v7.0.0-rc.37. Please refer to the changelog (linked) for details on
   the web release.

  * Bugfix [owncloud/web#6423](https://github.com/owncloud/web/issues/6423): Archiver in protected public links
  * Bugfix [owncloud/web#6434](https://github.com/owncloud/web/issues/6434): Endless lazy loading indicator after sorting file table
  * Bugfix [owncloud/web#6731](https://github.com/owncloud/web/issues/6731): Layout with long breadcrumb
  * Bugfix [owncloud/web#6768](https://github.com/owncloud/web/issues/6768): Pagination after increasing items per page
  * Bugfix [owncloud/web#7513](https://github.com/owncloud/web/issues/7513): Calendar popup position in right sidebar
  * Bugfix [owncloud/web#7655](https://github.com/owncloud/web/issues/7655): Loading shares in deep nested folders
  * Bugfix [owncloud/web#7925](https://github.com/owncloud/web/pull/7925): "Paste"-action without write permissions
  * Bugfix [owncloud/web#7926](https://github.com/owncloud/web/pull/7926): Include spaces in the list info
  * Bugfix [owncloud/web#7958](https://github.com/owncloud/web/pull/7958): Prevent deletion of own account
  * Bugfix [owncloud/web#7966](https://github.com/owncloud/web/pull/7966): UI fixes for sorting and quickactions
  * Bugfix [owncloud/web#7969](https://github.com/owncloud/web/pull/7969): Space quota not displayed after creation
  * Bugfix [owncloud/web#8026](https://github.com/owncloud/web/pull/8026): Text editor appearance
  * Bugfix [owncloud/web#8040](https://github.com/owncloud/web/pull/8040): Reverting versions for read-only shares
  * Bugfix [owncloud/web#8045](https://github.com/owncloud/web/pull/8045): Resolving drives in search
  * Bugfix [owncloud/web#8054](https://github.com/owncloud/web/issues/8054): Search repeating no results message
  * Bugfix [owncloud/web#8058](https://github.com/owncloud/web/pull/8058): Current year selection in the date picker
  * Bugfix [owncloud/web#8061](https://github.com/owncloud/web/pull/8061): Omit "page"-query in breadcrumb navigation
  * Bugfix [owncloud/web#8080](https://github.com/owncloud/web/pull/8080): Left sidebar navigation item text flickers on transition
  * Bugfix [owncloud/web#8081](https://github.com/owncloud/web/issues/8081): Space member disappearing
  * Bugfix [owncloud/web#8083](https://github.com/owncloud/web/issues/8083): Re-using space images
  * Bugfix [owncloud/web#8148](https://github.com/owncloud/web/issues/8148): Show space members despite deleted entries
  * Bugfix [owncloud/web#8158](https://github.com/owncloud/web/issues/8158): Search bar input appearance
  * Bugfix [owncloud/web#8265](https://github.com/owncloud/web/pull/8265): Application menu active display on hover
  * Bugfix [owncloud/web#8276](https://github.com/owncloud/web/pull/8276): Loading additional user data
  * Bugfix [owncloud/web#8300](https://github.com/owncloud/web/pull/8300): Re-loading space members panel
  * Bugfix [owncloud/web#8326](https://github.com/owncloud/web/pull/8326): Editing users who never logged in
  * Bugfix [owncloud/web#8340](https://github.com/owncloud/web/pull/8340): Cancel custom permissions
  * Bugfix [owncloud/web#8411](https://github.com/owncloud/web/issues/8411): Drop menus with limited vertical screen space
  * Bugfix [owncloud/web#8420](https://github.com/owncloud/web/issues/8420): Token renewal in vue router hash mode
  * Bugfix [owncloud/web#8434](https://github.com/owncloud/web/issues/8434): Accessing route in admin-settings with insufficient permissions
  * Bugfix [owncloud/web#8479](https://github.com/owncloud/web/issues/8479): "Show more"-action in shares panel
  * Bugfix [owncloud/web#8480](https://github.com/owncloud/web/pull/8480): Paste action conflict dialog broken
  * Bugfix [owncloud/web#8498](https://github.com/owncloud/web/pull/8498): PDF display issue - Update CSP object-src policy
  * Bugfix [owncloud/web#8508](https://github.com/owncloud/web/pull/8508): Remove fuzzy search results
  * Bugfix [owncloud/web#8523](https://github.com/owncloud/web/issues/8523): Space image upload
  * Bugfix [owncloud/web#8549](https://github.com/owncloud/web/issues/8549): Batch context actions in admin settings
  * Bugfix [owncloud/web#8554](https://github.com/owncloud/web/pull/8554): Height of dropdown no-option
  * Bugfix [owncloud/web#8576](https://github.com/owncloud/web/pull/8576): De-duplicate event handling to prevent errors on Draw-io
  * Bugfix [owncloud/web#8585](https://github.com/owncloud/web/issues/8585): Users without role assignment
  * Bugfix [owncloud/web#8587](https://github.com/owncloud/web/issues/8587): Password enforced check for public links
  * Bugfix [owncloud/web#8592](https://github.com/owncloud/web/issues/8592): Group members sorting
  * Bugfix [owncloud/web#8694](https://github.com/owncloud/web/pull/8694): Broken re-login after logout
  * Bugfix [owncloud/web#8695](https://github.com/owncloud/web/issues/8695): Open files in external app
  * Bugfix [owncloud/web#8756](https://github.com/owncloud/web/pull/8756): Copy link to clipboard text
  * Bugfix [owncloud/web#8758](https://github.com/owncloud/web/pull/8758): Preview controls colors
  * Bugfix [owncloud/web#8776](https://github.com/owncloud/web/issues/8776): Selection reset on action click
  * Bugfix [owncloud/web#8814](https://github.com/owncloud/web/pull/8814): Share recipient container exceed
  * Bugfix [owncloud/web#8825](https://github.com/owncloud/web/pull/8825): Remove drop target in read-only folders
  * Bugfix [owncloud/web#8827](https://github.com/owncloud/web/pull/8827): Opening context menu via keyboard
  * Bugfix [owncloud/web#8834](https://github.com/owncloud/web/issues/8834): Hide upload hint in empty read-only folders
  * Bugfix [owncloud/web#8864](https://github.com/owncloud/web/pull/8864): Public link empty password stays forever
  * Bugfix [owncloud/web#8880](https://github.com/owncloud/web/issues/8880): Sidebar header after deleting resource
  * Bugfix [owncloud/web#8928](https://github.com/owncloud/web/issues/8928): Infinite login redirect
  * Bugfix [owncloud/web#8987](https://github.com/owncloud/web/pull/8987): Limit amount of concurrent tus requests
  * Bugfix [owncloud/web#8992](https://github.com/owncloud/web/pull/8992): Personal space name after language change
  * Bugfix [owncloud/web#9004](https://github.com/owncloud/web/issues/9004): Endless loading when encountering a public link error
  * Bugfix [owncloud/web#9015](https://github.com/owncloud/web/pull/9015): Prevent "virtual" spaces from being displayed in the UI
  * Change [owncloud/web#6661](https://github.com/owncloud/web/issues/6661): Streamline new tab handling in extensions
  * Change [owncloud/web#7948](https://github.com/owncloud/web/issues/7948): Update Vue to v3.2
  * Change [owncloud/web#8431](https://github.com/owncloud/web/pull/8431): Remove permission manager
  * Change [owncloud/web#8455](https://github.com/owncloud/web/pull/8455): Configurable extension autosave
  * Change [owncloud/web#8563](https://github.com/owncloud/web/pull/8563): Theme colors
  * Enhancement [owncloud/web#6183](https://github.com/owncloud/web/issues/6183): Global loading indicator
  * Enhancement [owncloud/web#7388](https://github.com/owncloud/web/pull/7388): Add tag support
  * Enhancement [owncloud/web#7721](https://github.com/owncloud/web/issues/7721): Improve performance when loading folders and share indicators
  * Enhancement [owncloud/web#7942](https://github.com/owncloud/web/pull/7942): Warn users when using unsupported browsers
  * Enhancement [owncloud/web#7965](https://github.com/owncloud/web/pull/7965): Optional Contributor role and configurable resharing permissions
  * Enhancement [owncloud/web#7968](https://github.com/owncloud/web/pull/7968): Group and user creation forms submit on enter
  * Enhancement [owncloud/web#7976](https://github.com/owncloud/web/pull/7976): Add switch to enable condensed resource table
  * Enhancement [owncloud/web#7977](https://github.com/owncloud/web/pull/7977): Introduce zoom and rotate to the preview app
  * Enhancement [owncloud/web#7983](https://github.com/owncloud/web/pull/7983): Conflict dialog UX
  * Enhancement [owncloud/web#7991](https://github.com/owncloud/web/pull/7991): Add tiles view for resource display
  * Enhancement [owncloud/web#7994](https://github.com/owncloud/web/pull/7994): Introduce full screen mode to the preview app
  * Enhancement [owncloud/web#7995](https://github.com/owncloud/web/pull/7995): Enable autoplay in the preview app
  * Enhancement [owncloud/web#8008](https://github.com/owncloud/web/issues/8008): Don't open sidebar when copying quicklink
  * Enhancement [owncloud/web#8021](https://github.com/owncloud/web/pull/8021): Access right sidebar panels via URL
  * Enhancement [owncloud/web#8051](https://github.com/owncloud/web/pull/8051): Introduce image preloading to the preview app
  * Enhancement [owncloud/web#8055](https://github.com/owncloud/web/pull/8055): Retry failed uploads on re-upload
  * Enhancement [owncloud/web#8056](https://github.com/owncloud/web/pull/8056): Increase Searchbar height
  * Enhancement [owncloud/web#8057](https://github.com/owncloud/web/pull/8057): Show text file icon for empty text files
  * Enhancement [owncloud/web#8132](https://github.com/owncloud/web/pull/8132): Update libre-graph-api to v1.0
  * Enhancement [owncloud/web#8136](https://github.com/owncloud/web/pull/8136): Make clipboard copy available to more browsers
  * Enhancement [owncloud/web#8161](https://github.com/owncloud/web/pull/8161): Space group members
  * Enhancement [owncloud/web#8161](https://github.com/owncloud/web/pull/8161): Space group shares
  * Enhancement [owncloud/web#8166](https://github.com/owncloud/web/issues/8166): Show upload speed
  * Enhancement [owncloud/web#8175](https://github.com/owncloud/web/pull/8175): Rename "user management" app
  * Enhancement [owncloud/web#8178](https://github.com/owncloud/web/pull/8178): Spaces list in admin settings
  * Enhancement [owncloud/web#8261](https://github.com/owncloud/web/pull/8261): Admin settings users section uses graph api for role assignments
  * Enhancement [owncloud/web#8279](https://github.com/owncloud/web/pull/8279): Move user group select to edit panel
  * Enhancement [owncloud/web#8280](https://github.com/owncloud/web/pull/8280): Add support for multiple clients in `theme.json`
  * Enhancement [owncloud/web#8294](https://github.com/owncloud/web/pull/8294): Move language selection to user account page
  * Enhancement [owncloud/web#8306](https://github.com/owncloud/web/pull/8306): Show selectable groups only
  * Enhancement [owncloud/web#8317](https://github.com/owncloud/web/pull/8317): Add context menu to groups
  * Enhancement [owncloud/web#8320](https://github.com/owncloud/web/pull/8320): Space member expiration
  * Enhancement [owncloud/web#8320](https://github.com/owncloud/web/pull/8320): Update SDK to v3.1.0-alpha.3
  * Enhancement [owncloud/web#8324](https://github.com/owncloud/web/pull/8324): Add context menu to users
  * Enhancement [owncloud/web#8331](https://github.com/owncloud/web/pull/8331): Admin settings users section details improvement
  * Enhancement [owncloud/web#8354](https://github.com/owncloud/web/issues/8354): Add `ItemFilter` component
  * Enhancement [owncloud/web#8356](https://github.com/owncloud/web/pull/8356): Slight improvement of key up/down performance
  * Enhancement [owncloud/web#8363](https://github.com/owncloud/web/issues/8363): Admin settings general section
  * Enhancement [owncloud/web#8375](https://github.com/owncloud/web/pull/8375): Add appearance section in general settings
  * Enhancement [owncloud/web#8377](https://github.com/owncloud/web/issues/8377): User group filter
  * Enhancement [owncloud/web#8387](https://github.com/owncloud/web/pull/8387): Batch edit quota in admin panel
  * Enhancement [owncloud/web#8398](https://github.com/owncloud/web/pull/8398): Use standardized layout for file/space action list
  * Enhancement [owncloud/web#8425](https://github.com/owncloud/web/issues/8425): Add dark ownCloud logo
  * Enhancement [owncloud/web#8432](https://github.com/owncloud/web/pull/8432): Inject customizations
  * Enhancement [owncloud/web#8433](https://github.com/owncloud/web/pull/8433): User settings login field
  * Enhancement [owncloud/web#8441](https://github.com/owncloud/web/pull/8441): Skeleton App
  * Enhancement [owncloud/web#8449](https://github.com/owncloud/web/pull/8449): Configurable top bar
  * Enhancement [owncloud/web#8450](https://github.com/owncloud/web/pull/8450): Rework notification bell
  * Enhancement [owncloud/web#8455](https://github.com/owncloud/web/pull/8455): Autosave content changes in text editor
  * Enhancement [owncloud/web#8473](https://github.com/owncloud/web/pull/8473): Update CERN links
  * Enhancement [owncloud/web#8489](https://github.com/owncloud/web/pull/8489): Respect max quota
  * Enhancement [owncloud/web#8492](https://github.com/owncloud/web/pull/8492): User role filter
  * Enhancement [owncloud/web#8503](https://github.com/owncloud/web/issues/8503): Beautify file version list
  * Enhancement [owncloud/web#8515](https://github.com/owncloud/web/pull/8515): Introduce trashbin overview
  * Enhancement [owncloud/web#8518](https://github.com/owncloud/web/pull/8518): Make notifications work with oCIS
  * Enhancement [owncloud/web#8541](https://github.com/owncloud/web/pull/8541): Public link permission `PublicLink.Write.all`
  * Enhancement [owncloud/web#8553](https://github.com/owncloud/web/pull/8553): Add and remove users from groups batch actions
  * Enhancement [owncloud/web#8554](https://github.com/owncloud/web/pull/8554): Beautify form inputs
  * Enhancement [owncloud/web#8557](https://github.com/owncloud/web/issues/8557): Rework mobile navigation
  * Enhancement [owncloud/web#8566](https://github.com/owncloud/web/pull/8566): QuickActions role configurable
  * Enhancement [owncloud/web#8612](https://github.com/owncloud/web/issues/8612): Add `Accept-Language` header to all outgoing requests
  * Enhancement [owncloud/web#8630](https://github.com/owncloud/web/pull/8630): Add logout url
  * Enhancement [owncloud/web#8652](https://github.com/owncloud/web/pull/8652): Enable guest users
  * Enhancement [owncloud/web#8711](https://github.com/owncloud/web/pull/8711): Remove placeholder, add customizable label
  * Enhancement [owncloud/web#8713](https://github.com/owncloud/web/pull/8713): Context helper read more link configurable
  * Enhancement [owncloud/web#8715](https://github.com/owncloud/web/pull/8715): Enable rename groups
  * Enhancement [owncloud/web#8730](https://github.com/owncloud/web/pull/8730): Create Space from selection
  * Enhancement [owncloud/web#8738](https://github.com/owncloud/web/issues/8738): GDPR export
  * Enhancement [owncloud/web#8762](https://github.com/owncloud/web/pull/8762): Stop bootstrapping application earlier in anonymous contexts
  * Enhancement [owncloud/web#8766](https://github.com/owncloud/web/pull/8766): Add support for read-only groups
  * Enhancement [owncloud/web#8790](https://github.com/owncloud/web/pull/8790): Custom translations
  * Enhancement [owncloud/web#8797](https://github.com/owncloud/web/pull/8797): Font family in theming
  * Enhancement [owncloud/web#8806](https://github.com/owncloud/web/pull/8806): Preview app sorting
  * Enhancement [owncloud/web#8820](https://github.com/owncloud/web/pull/8820): Adjust missing reshare permissions message
  * Enhancement [owncloud/web#8822](https://github.com/owncloud/web/pull/8822): Fix quicklink icon alignment
  * Enhancement [owncloud/web#8826](https://github.com/owncloud/web/pull/8826): Admin settings groups members panel
  * Enhancement [owncloud/web#8868](https://github.com/owncloud/web/pull/8868): Respect user read-only configuration by the server
  * Enhancement [owncloud/web#8876](https://github.com/owncloud/web/pull/8876): Update roles and permissions names, labels, texts and icons
  * Enhancement [owncloud/web#8882](https://github.com/owncloud/web/pull/8882): Layout of Share role and expiration date dropdown
  * Enhancement [owncloud/web#8883](https://github.com/owncloud/web/issues/8883): Webfinger redirect app
  * Enhancement [owncloud/web#8898](https://github.com/owncloud/web/pull/8898): Rename "Quicklink" to "link"
  * Enhancement [owncloud/web#8911](https://github.com/owncloud/web/pull/8911): Add notification setting to account page

   https://github.com/owncloud/ocis/pull/6294
   https://github.com/owncloud/web/releases/tag/v7.0.0-rc.37

* Enhancement - Update web to v7.0.0-rc.38: [#6375](https://github.com/owncloud/ocis/pull/6375)

   Tags: web

   We updated ownCloud Web to v7.0.0-rc.38. Please refer to the changelog (linked) for details on
   the web release.

  * Bugfix [owncloud/web#6423](https://github.com/owncloud/web/issues/6423): Archiver in protected public links
  * Bugfix [owncloud/web#6434](https://github.com/owncloud/web/issues/6434): Endless lazy loading indicator after sorting file table
  * Bugfix [owncloud/web#6731](https://github.com/owncloud/web/issues/6731): Layout with long breadcrumb
  * Bugfix [owncloud/web#6768](https://github.com/owncloud/web/issues/6768): Pagination after increasing items per page
  * Bugfix [owncloud/web#7513](https://github.com/owncloud/web/issues/7513): Calendar popup position in right sidebar
  * Bugfix [owncloud/web#7655](https://github.com/owncloud/web/issues/7655): Loading shares in deep nested folders
  * Bugfix [owncloud/web#7925](https://github.com/owncloud/web/pull/7925): "Paste"-action without write permissions
  * Bugfix [owncloud/web#7926](https://github.com/owncloud/web/pull/7926): Include spaces in the list info
  * Bugfix [owncloud/web#7958](https://github.com/owncloud/web/pull/7958): Prevent deletion of own account
  * Bugfix [owncloud/web#7966](https://github.com/owncloud/web/pull/7966): UI fixes for sorting and quickactions
  * Bugfix [owncloud/web#7969](https://github.com/owncloud/web/pull/7969): Space quota not displayed after creation
  * Bugfix [owncloud/web#8026](https://github.com/owncloud/web/pull/8026): Text editor appearance
  * Bugfix [owncloud/web#8040](https://github.com/owncloud/web/pull/8040): Reverting versions for read-only shares
  * Bugfix [owncloud/web#8045](https://github.com/owncloud/web/pull/8045): Resolving drives in search
  * Bugfix [owncloud/web#8054](https://github.com/owncloud/web/issues/8054): Search repeating no results message
  * Bugfix [owncloud/web#8058](https://github.com/owncloud/web/pull/8058): Current year selection in the date picker
  * Bugfix [owncloud/web#8061](https://github.com/owncloud/web/pull/8061): Omit "page"-query in breadcrumb navigation
  * Bugfix [owncloud/web#8080](https://github.com/owncloud/web/pull/8080): Left sidebar navigation item text flickers on transition
  * Bugfix [owncloud/web#8081](https://github.com/owncloud/web/issues/8081): Space member disappearing
  * Bugfix [owncloud/web#8083](https://github.com/owncloud/web/issues/8083): Re-using space images
  * Bugfix [owncloud/web#8148](https://github.com/owncloud/web/issues/8148): Show space members despite deleted entries
  * Bugfix [owncloud/web#8158](https://github.com/owncloud/web/issues/8158): Search bar input appearance
  * Bugfix [owncloud/web#8265](https://github.com/owncloud/web/pull/8265): Application menu active display on hover
  * Bugfix [owncloud/web#8276](https://github.com/owncloud/web/pull/8276): Loading additional user data
  * Bugfix [owncloud/web#8300](https://github.com/owncloud/web/pull/8300): Re-loading space members panel
  * Bugfix [owncloud/web#8326](https://github.com/owncloud/web/pull/8326): Editing users who never logged in
  * Bugfix [owncloud/web#8340](https://github.com/owncloud/web/pull/8340): Cancel custom permissions
  * Bugfix [owncloud/web#8411](https://github.com/owncloud/web/issues/8411): Drop menus with limited vertical screen space
  * Bugfix [owncloud/web#8420](https://github.com/owncloud/web/issues/8420): Token renewal in vue router hash mode
  * Bugfix [owncloud/web#8434](https://github.com/owncloud/web/issues/8434): Accessing route in admin-settings with insufficient permissions
  * Bugfix [owncloud/web#8479](https://github.com/owncloud/web/issues/8479): "Show more"-action in shares panel
  * Bugfix [owncloud/web#8480](https://github.com/owncloud/web/pull/8480): Paste action conflict dialog broken
  * Bugfix [owncloud/web#8498](https://github.com/owncloud/web/pull/8498): PDF display issue - Update CSP object-src policy
  * Bugfix [owncloud/web#8508](https://github.com/owncloud/web/pull/8508): Remove fuzzy search results
  * Bugfix [owncloud/web#8523](https://github.com/owncloud/web/issues/8523): Space image upload
  * Bugfix [owncloud/web#8549](https://github.com/owncloud/web/issues/8549): Batch context actions in admin settings
  * Bugfix [owncloud/web#8554](https://github.com/owncloud/web/pull/8554): Height of dropdown no-option
  * Bugfix [owncloud/web#8576](https://github.com/owncloud/web/pull/8576): De-duplicate event handling to prevent errors on Draw-io
  * Bugfix [owncloud/web#8585](https://github.com/owncloud/web/issues/8585): Users without role assignment
  * Bugfix [owncloud/web#8587](https://github.com/owncloud/web/issues/8587): Password enforced check for public links
  * Bugfix [owncloud/web#8592](https://github.com/owncloud/web/issues/8592): Group members sorting
  * Bugfix [owncloud/web#8694](https://github.com/owncloud/web/pull/8694): Broken re-login after logout
  * Bugfix [owncloud/web#8695](https://github.com/owncloud/web/issues/8695): Open files in external app
  * Bugfix [owncloud/web#8756](https://github.com/owncloud/web/pull/8756): Copy link to clipboard text
  * Bugfix [owncloud/web#8758](https://github.com/owncloud/web/pull/8758): Preview controls colors
  * Bugfix [owncloud/web#8776](https://github.com/owncloud/web/issues/8776): Selection reset on action click
  * Bugfix [owncloud/web#8814](https://github.com/owncloud/web/pull/8814): Share recipient container exceed
  * Bugfix [owncloud/web#8825](https://github.com/owncloud/web/pull/8825): Remove drop target in read-only folders
  * Bugfix [owncloud/web#8827](https://github.com/owncloud/web/pull/8827): Opening context menu via keyboard
  * Bugfix [owncloud/web#8834](https://github.com/owncloud/web/issues/8834): Hide upload hint in empty read-only folders
  * Bugfix [owncloud/web#8864](https://github.com/owncloud/web/pull/8864): Public link empty password stays forever
  * Bugfix [owncloud/web#8880](https://github.com/owncloud/web/issues/8880): Sidebar header after deleting resource
  * Bugfix [owncloud/web#8928](https://github.com/owncloud/web/issues/8928): Infinite login redirect
  * Bugfix [owncloud/web#8987](https://github.com/owncloud/web/pull/8987): Limit amount of concurrent tus requests
  * Bugfix [owncloud/web#8992](https://github.com/owncloud/web/pull/8992): Personal space name after language change
  * Bugfix [owncloud/web#9004](https://github.com/owncloud/web/issues/9004): Endless loading when encountering a public link error
  * Bugfix [owncloud/web#9015](https://github.com/owncloud/web/pull/9015): Prevent "virtual" spaces from being displayed in the UI
  * Bugfix [owncloud/web#9022](https://github.com/owncloud/web/issues/9022): Spaces in search results
  * Bugfix [owncloud/web#9061](https://github.com/owncloud/web/issues/9061): Resource not found and No content message at the same time
  * Change [owncloud/web#6661](https://github.com/owncloud/web/issues/6661): Streamline new tab handling in extensions
  * Change [owncloud/web#7948](https://github.com/owncloud/web/issues/7948): Update Vue to v3.2
  * Change [owncloud/web#8431](https://github.com/owncloud/web/pull/8431): Remove permission manager
  * Change [owncloud/web#8455](https://github.com/owncloud/web/pull/8455): Configurable extension autosave
  * Change [owncloud/web#8563](https://github.com/owncloud/web/pull/8563): Theme colors
  * Enhancement [owncloud/web#6183](https://github.com/owncloud/web/issues/6183): Global loading indicator
  * Enhancement [owncloud/web#7388](https://github.com/owncloud/web/pull/7388): Add tag support
  * Enhancement [owncloud/web#7721](https://github.com/owncloud/web/issues/7721): Improve performance when loading folders and share indicators
  * Enhancement [owncloud/web#7942](https://github.com/owncloud/web/pull/7942): Warn users when using unsupported browsers
  * Enhancement [owncloud/web#7965](https://github.com/owncloud/web/pull/7965): Optional Contributor role and configurable resharing permissions
  * Enhancement [owncloud/web#7968](https://github.com/owncloud/web/pull/7968): Group and user creation forms submit on enter
  * Enhancement [owncloud/web#7976](https://github.com/owncloud/web/pull/7976): Add switch to enable condensed resource table
  * Enhancement [owncloud/web#7977](https://github.com/owncloud/web/pull/7977): Introduce zoom and rotate to the preview app
  * Enhancement [owncloud/web#7983](https://github.com/owncloud/web/pull/7983): Conflict dialog UX
  * Enhancement [owncloud/web#7991](https://github.com/owncloud/web/pull/7991): Add tiles view for resource display
  * Enhancement [owncloud/web#7994](https://github.com/owncloud/web/pull/7994): Introduce full screen mode to the preview app
  * Enhancement [owncloud/web#7995](https://github.com/owncloud/web/pull/7995): Enable autoplay in the preview app
  * Enhancement [owncloud/web#8008](https://github.com/owncloud/web/issues/8008): Don't open sidebar when copying quicklink
  * Enhancement [owncloud/web#8021](https://github.com/owncloud/web/pull/8021): Access right sidebar panels via URL
  * Enhancement [owncloud/web#8051](https://github.com/owncloud/web/pull/8051): Introduce image preloading to the preview app
  * Enhancement [owncloud/web#8055](https://github.com/owncloud/web/pull/8055): Retry failed uploads on re-upload
  * Enhancement [owncloud/web#8056](https://github.com/owncloud/web/pull/8056): Increase Searchbar height
  * Enhancement [owncloud/web#8057](https://github.com/owncloud/web/pull/8057): Show text file icon for empty text files
  * Enhancement [owncloud/web#8132](https://github.com/owncloud/web/pull/8132): Update libre-graph-api to v1.0
  * Enhancement [owncloud/web#8136](https://github.com/owncloud/web/pull/8136): Make clipboard copy available to more browsers
  * Enhancement [owncloud/web#8161](https://github.com/owncloud/web/pull/8161): Space group members
  * Enhancement [owncloud/web#8161](https://github.com/owncloud/web/pull/8161): Space group shares
  * Enhancement [owncloud/web#8166](https://github.com/owncloud/web/issues/8166): Show upload speed
  * Enhancement [owncloud/web#8175](https://github.com/owncloud/web/pull/8175): Rename "user management" app
  * Enhancement [owncloud/web#8178](https://github.com/owncloud/web/pull/8178): Spaces list in admin settings
  * Enhancement [owncloud/web#8261](https://github.com/owncloud/web/pull/8261): Admin settings users section uses graph api for role assignments
  * Enhancement [owncloud/web#8279](https://github.com/owncloud/web/pull/8279): Move user group select to edit panel
  * Enhancement [owncloud/web#8280](https://github.com/owncloud/web/pull/8280): Add support for multiple clients in `theme.json`
  * Enhancement [owncloud/web#8294](https://github.com/owncloud/web/pull/8294): Move language selection to user account page
  * Enhancement [owncloud/web#8306](https://github.com/owncloud/web/pull/8306): Show selectable groups only
  * Enhancement [owncloud/web#8317](https://github.com/owncloud/web/pull/8317): Add context menu to groups
  * Enhancement [owncloud/web#8320](https://github.com/owncloud/web/pull/8320): Space member expiration
  * Enhancement [owncloud/web#8320](https://github.com/owncloud/web/pull/8320): Update SDK to v3.1.0-alpha.3
  * Enhancement [owncloud/web#8324](https://github.com/owncloud/web/pull/8324): Add context menu to users
  * Enhancement [owncloud/web#8331](https://github.com/owncloud/web/pull/8331): Admin settings users section details improvement
  * Enhancement [owncloud/web#8354](https://github.com/owncloud/web/issues/8354): Add `ItemFilter` component
  * Enhancement [owncloud/web#8356](https://github.com/owncloud/web/pull/8356): Slight improvement of key up/down performance
  * Enhancement [owncloud/web#8363](https://github.com/owncloud/web/issues/8363): Admin settings general section
  * Enhancement [owncloud/web#8375](https://github.com/owncloud/web/pull/8375): Add appearance section in general settings
  * Enhancement [owncloud/web#8377](https://github.com/owncloud/web/issues/8377): User group filter
  * Enhancement [owncloud/web#8387](https://github.com/owncloud/web/pull/8387): Batch edit quota in admin panel
  * Enhancement [owncloud/web#8398](https://github.com/owncloud/web/pull/8398): Use standardized layout for file/space action list
  * Enhancement [owncloud/web#8425](https://github.com/owncloud/web/issues/8425): Add dark ownCloud logo
  * Enhancement [owncloud/web#8432](https://github.com/owncloud/web/pull/8432): Inject customizations
  * Enhancement [owncloud/web#8433](https://github.com/owncloud/web/pull/8433): User settings login field
  * Enhancement [owncloud/web#8441](https://github.com/owncloud/web/pull/8441): Skeleton App
  * Enhancement [owncloud/web#8449](https://github.com/owncloud/web/pull/8449): Configurable top bar
  * Enhancement [owncloud/web#8450](https://github.com/owncloud/web/pull/8450): Rework notification bell
  * Enhancement [owncloud/web#8455](https://github.com/owncloud/web/pull/8455): Autosave content changes in text editor
  * Enhancement [owncloud/web#8473](https://github.com/owncloud/web/pull/8473): Update CERN links
  * Enhancement [owncloud/web#8489](https://github.com/owncloud/web/pull/8489): Respect max quota
  * Enhancement [owncloud/web#8492](https://github.com/owncloud/web/pull/8492): User role filter
  * Enhancement [owncloud/web#8503](https://github.com/owncloud/web/issues/8503): Beautify file version list
  * Enhancement [owncloud/web#8515](https://github.com/owncloud/web/pull/8515): Introduce trashbin overview
  * Enhancement [owncloud/web#8518](https://github.com/owncloud/web/pull/8518): Make notifications work with oCIS
  * Enhancement [owncloud/web#8541](https://github.com/owncloud/web/pull/8541): Public link permission `PublicLink.Write.all`
  * Enhancement [owncloud/web#8553](https://github.com/owncloud/web/pull/8553): Add and remove users from groups batch actions
  * Enhancement [owncloud/web#8554](https://github.com/owncloud/web/pull/8554): Beautify form inputs
  * Enhancement [owncloud/web#8557](https://github.com/owncloud/web/issues/8557): Rework mobile navigation
  * Enhancement [owncloud/web#8566](https://github.com/owncloud/web/pull/8566): QuickActions role configurable
  * Enhancement [owncloud/web#8612](https://github.com/owncloud/web/issues/8612): Add `Accept-Language` header to all outgoing requests
  * Enhancement [owncloud/web#8630](https://github.com/owncloud/web/pull/8630): Add logout url
  * Enhancement [owncloud/web#8652](https://github.com/owncloud/web/pull/8652): Enable guest users
  * Enhancement [owncloud/web#8711](https://github.com/owncloud/web/pull/8711): Remove placeholder, add customizable label
  * Enhancement [owncloud/web#8713](https://github.com/owncloud/web/pull/8713): Context helper read more link configurable
  * Enhancement [owncloud/web#8715](https://github.com/owncloud/web/pull/8715): Enable rename groups
  * Enhancement [owncloud/web#8730](https://github.com/owncloud/web/pull/8730): Create Space from selection
  * Enhancement [owncloud/web#8738](https://github.com/owncloud/web/issues/8738): GDPR export
  * Enhancement [owncloud/web#8762](https://github.com/owncloud/web/pull/8762): Stop bootstrapping application earlier in anonymous contexts
  * Enhancement [owncloud/web#8766](https://github.com/owncloud/web/pull/8766): Add support for read-only groups
  * Enhancement [owncloud/web#8790](https://github.com/owncloud/web/pull/8790): Custom translations
  * Enhancement [owncloud/web#8797](https://github.com/owncloud/web/pull/8797): Font family in theming
  * Enhancement [owncloud/web#8806](https://github.com/owncloud/web/pull/8806): Preview app sorting
  * Enhancement [owncloud/web#8820](https://github.com/owncloud/web/pull/8820): Adjust missing reshare permissions message
  * Enhancement [owncloud/web#8822](https://github.com/owncloud/web/pull/8822): Fix quicklink icon alignment
  * Enhancement [owncloud/web#8826](https://github.com/owncloud/web/pull/8826): Admin settings groups members panel
  * Enhancement [owncloud/web#8868](https://github.com/owncloud/web/pull/8868): Respect user read-only configuration by the server
  * Enhancement [owncloud/web#8876](https://github.com/owncloud/web/pull/8876): Update roles and permissions names, labels, texts and icons
  * Enhancement [owncloud/web#8882](https://github.com/owncloud/web/pull/8882): Layout of Share role and expiration date dropdown
  * Enhancement [owncloud/web#8883](https://github.com/owncloud/web/issues/8883): Webfinger redirect app
  * Enhancement [owncloud/web#8898](https://github.com/owncloud/web/pull/8898): Rename "Quicklink" to "link"
  * Enhancement [owncloud/web#8911](https://github.com/owncloud/web/pull/8911): Add notification setting to account page
  * Enhancement [owncloud/web#9070](https://github.com/owncloud/web/pull/9070): Disable change password capability
  * Enhancement [owncloud/web#9070](https://github.com/owncloud/web/pull/9070): Disable create user and delete user via capabilities
  * Enhancement [owncloud/web#9076](https://github.com/owncloud/web/pull/9076): Show detailed error messages while upload fails

   https://github.com/owncloud/ocis/pull/6375
   https://github.com/owncloud/web/releases/tag/v7.0.0-rc.38

* Enhancement - Update web to v7.0.0: [#6438](https://github.com/owncloud/ocis/pull/6438)

   Tags: web

   We updated ownCloud Web to v7.0.0. Please refer to the changelog (linked) for details on the web
   release.

   ## Breaking changes * BREAKING CHANGE for developers and admins in
   [owncloud/web#7948](https://github.com/owncloud/web/issues/7948): we've updated
   Vue.js to version 3. Existing apps that have not been updated to Vue.js version 3 will not be
   compatible anymore. * BREAKING CHANGE for admins in
   [owncloud/web#8563](https://github.com/owncloud/web/pull/8563): we've introduced
   contrast colors in our theming. In case you have created a custom `theme.json` it needs to be
   adjusted accordingly: `-contrast` color values need to be added to all `swatches`, e.g. to
   `swatch-brand-contrast`. See https://owncloud.dev/clients/web/theming/#colors

   ## Summary * Bugfix
   [owncloud/web#6423](https://github.com/owncloud/web/issues/6423): Archiver in
   protected public links * Bugfix
   [owncloud/web#6434](https://github.com/owncloud/web/issues/6434): Endless lazy
   loading indicator after sorting file table * Bugfix
   [owncloud/web#6731](https://github.com/owncloud/web/issues/6731): Layout with long
   breadcrumb * Bugfix
   [owncloud/web#6768](https://github.com/owncloud/web/issues/6768): Pagination after
   increasing items per page * Bugfix
   [owncloud/web#7513](https://github.com/owncloud/web/issues/7513): Calendar popup
   position in right sidebar * Bugfix
   [owncloud/web#7655](https://github.com/owncloud/web/issues/7655): Loading shares in
   deep nested folders * Bugfix
   [owncloud/web#7925](https://github.com/owncloud/web/pull/7925): "Paste"-action
   without write permissions * Bugfix
   [owncloud/web#7926](https://github.com/owncloud/web/pull/7926): Include spaces in
   the list info * Bugfix
   [owncloud/web#7958](https://github.com/owncloud/web/pull/7958): Prevent deletion of
   own account * Bugfix [owncloud/web#7966](https://github.com/owncloud/web/pull/7966):
   UI fixes for sorting and quickactions * Bugfix
   [owncloud/web#7969](https://github.com/owncloud/web/pull/7969): Space quota not
   displayed after creation * Bugfix
   [owncloud/web#8026](https://github.com/owncloud/web/pull/8026): Text editor
   appearance * Bugfix [owncloud/web#8040](https://github.com/owncloud/web/pull/8040):
   Reverting versions for read-only shares * Bugfix
   [owncloud/web#8045](https://github.com/owncloud/web/pull/8045): Resolving drives in
   search * Bugfix [owncloud/web#8054](https://github.com/owncloud/web/issues/8054):
   Search repeating no results message * Bugfix
   [owncloud/web#8058](https://github.com/owncloud/web/pull/8058): Current year
   selection in the date picker * Bugfix
   [owncloud/web#8061](https://github.com/owncloud/web/pull/8061): Omit "page"-query
   in breadcrumb navigation * Bugfix
   [owncloud/web#8080](https://github.com/owncloud/web/pull/8080): Left sidebar
   navigation item text flickers on transition * Bugfix
   [owncloud/web#8081](https://github.com/owncloud/web/issues/8081): Space member
   disappearing * Bugfix
   [owncloud/web#8083](https://github.com/owncloud/web/issues/8083): Re-using space
   images * Bugfix [owncloud/web#8148](https://github.com/owncloud/web/issues/8148):
   Show space members despite deleted entries * Bugfix
   [owncloud/web#8158](https://github.com/owncloud/web/issues/8158): Search bar input
   appearance * Bugfix [owncloud/web#8265](https://github.com/owncloud/web/pull/8265):
   Application menu active display on hover * Bugfix
   [owncloud/web#8276](https://github.com/owncloud/web/pull/8276): Loading additional
   user data * Bugfix [owncloud/web#8300](https://github.com/owncloud/web/pull/8300):
   Re-loading space members panel * Bugfix
   [owncloud/web#8326](https://github.com/owncloud/web/pull/8326): Editing users who
   never logged in * Bugfix
   [owncloud/web#8340](https://github.com/owncloud/web/pull/8340): Cancel custom
   permissions * Bugfix
   [owncloud/web#8411](https://github.com/owncloud/web/issues/8411): Drop menus with
   limited vertical screen space * Bugfix
   [owncloud/web#8420](https://github.com/owncloud/web/issues/8420): Token renewal in
   vue router hash mode * Bugfix
   [owncloud/web#8434](https://github.com/owncloud/web/issues/8434): Accessing route
   in admin-settings with insufficient permissions * Bugfix
   [owncloud/web#8479](https://github.com/owncloud/web/issues/8479): "Show
   more"-action in shares panel * Bugfix
   [owncloud/web#8480](https://github.com/owncloud/web/pull/8480): Paste action
   conflict dialog broken * Bugfix
   [owncloud/web#8498](https://github.com/owncloud/web/pull/8498): PDF display issue -
   Update CSP object-src policy * Bugfix
   [owncloud/web#8508](https://github.com/owncloud/web/pull/8508): Remove fuzzy search
   results * Bugfix [owncloud/web#8523](https://github.com/owncloud/web/issues/8523):
   Space image upload * Bugfix
   [owncloud/web#8549](https://github.com/owncloud/web/issues/8549): Batch context
   actions in admin settings * Bugfix
   [owncloud/web#8554](https://github.com/owncloud/web/pull/8554): Height of dropdown
   no-option * Bugfix [owncloud/web#8576](https://github.com/owncloud/web/pull/8576):
   De-duplicate event handling to prevent errors on Draw-io * Bugfix
   [owncloud/web#8585](https://github.com/owncloud/web/issues/8585): Users without
   role assignment * Bugfix
   [owncloud/web#8587](https://github.com/owncloud/web/issues/8587): Password
   enforced check for public links * Bugfix
   [owncloud/web#8592](https://github.com/owncloud/web/issues/8592): Group members
   sorting * Bugfix [owncloud/web#8694](https://github.com/owncloud/web/pull/8694):
   Broken re-login after logout * Bugfix
   [owncloud/web#8695](https://github.com/owncloud/web/issues/8695): Open files in
   external app * Bugfix
   [owncloud/web#8756](https://github.com/owncloud/web/pull/8756): Copy link to
   clipboard text * Bugfix
   [owncloud/web#8758](https://github.com/owncloud/web/pull/8758): Preview controls
   colors * Bugfix [owncloud/web#8776](https://github.com/owncloud/web/issues/8776):
   Selection reset on action click * Bugfix
   [owncloud/web#8814](https://github.com/owncloud/web/pull/8814): Share recipient
   container exceed * Bugfix
   [owncloud/web#8825](https://github.com/owncloud/web/pull/8825): Remove drop target
   in read-only folders * Bugfix
   [owncloud/web#8827](https://github.com/owncloud/web/pull/8827): Opening context
   menu via keyboard * Bugfix
   [owncloud/web#8834](https://github.com/owncloud/web/issues/8834): Hide upload hint
   in empty read-only folders * Bugfix
   [owncloud/web#8864](https://github.com/owncloud/web/pull/8864): Public link empty
   password stays forever * Bugfix
   [owncloud/web#8880](https://github.com/owncloud/web/issues/8880): Sidebar header
   after deleting resource * Bugfix
   [owncloud/web#8928](https://github.com/owncloud/web/issues/8928): Infinite login
   redirect * Bugfix [owncloud/web#8987](https://github.com/owncloud/web/pull/8987):
   Limit amount of concurrent tus requests * Bugfix
   [owncloud/web#8992](https://github.com/owncloud/web/pull/8992): Personal space name
   after language change * Bugfix
   [owncloud/web#9004](https://github.com/owncloud/web/issues/9004): Endless loading
   when encountering a public link error * Bugfix
   [owncloud/web#9009](https://github.com/owncloud/web/pull/9009): Public link file
   previews * Bugfix [owncloud/web#9014](https://github.com/owncloud/web/issues/9014):
   Empty file list after deleting resources * Bugfix
   [owncloud/web#9015](https://github.com/owncloud/web/pull/9015): Prevent "virtual"
   spaces from being displayed in the UI * Bugfix
   [owncloud/web#9020](https://github.com/owncloud/web/issues/9020): Sidebar for
   spaces on "Shared via link"-page * Bugfix
   [owncloud/web#9022](https://github.com/owncloud/web/issues/9022): Spaces in search
   results * Bugfix [owncloud/web#9030](https://github.com/owncloud/web/issues/9030):
   Share indicator loading after pasting resources * Bugfix
   [owncloud/web#9050](https://github.com/owncloud/web/issues/9050): Preview app mime
   type detection * Bugfix
   [owncloud/web#9061](https://github.com/owncloud/web/issues/9061): Resource not
   found and No content message at the same time * Bugfix
   [owncloud/web#9080](https://github.com/owncloud/web/issues/9080): Incorrect pause
   state in upload info * Bugfix
   [owncloud/web#9131](https://github.com/owncloud/web/pull/9131): Select all checkbox
   * Bugfix [owncloud/web#9144](https://github.com/owncloud/web/pull/9144):
   Notifications link overflow * Change
   [owncloud/web#6661](https://github.com/owncloud/web/issues/6661): Streamline new
   tab handling in extensions * Change
   [owncloud/web#7948](https://github.com/owncloud/web/issues/7948): Update Vue to v3.2
   * Change [owncloud/web#8431](https://github.com/owncloud/web/pull/8431): Remove
   permission manager * Change
   [owncloud/web#8455](https://github.com/owncloud/web/pull/8455): Configurable
   extension autosave * Change
   [owncloud/web#8563](https://github.com/owncloud/web/pull/8563): Theme colors *
   Enhancement [owncloud/web#6183](https://github.com/owncloud/web/issues/6183):
   Global loading indicator * Enhancement
   [owncloud/web#7388](https://github.com/owncloud/web/pull/7388): Add tag support *
   Enhancement [owncloud/web#7721](https://github.com/owncloud/web/issues/7721):
   Improve performance when loading folders and share indicators * Enhancement
   [owncloud/web#7942](https://github.com/owncloud/web/pull/7942): Warn users when
   using unsupported browsers * Enhancement
   [owncloud/web#7965](https://github.com/owncloud/web/pull/7965): Optional
   Contributor role and configurable resharing permissions * Enhancement
   [owncloud/web#7968](https://github.com/owncloud/web/pull/7968): Group and user
   creation forms submit on enter * Enhancement
   [owncloud/web#7976](https://github.com/owncloud/web/pull/7976): Add switch to enable
   condensed resource table * Enhancement
   [owncloud/web#7977](https://github.com/owncloud/web/pull/7977): Introduce zoom and
   rotate to the preview app * Enhancement
   [owncloud/web#7983](https://github.com/owncloud/web/pull/7983): Conflict dialog UX *
   Enhancement [owncloud/web#7991](https://github.com/owncloud/web/pull/7991): Add
   tiles view for resource display * Enhancement
   [owncloud/web#7994](https://github.com/owncloud/web/pull/7994): Introduce full
   screen mode to the preview app * Enhancement
   [owncloud/web#7995](https://github.com/owncloud/web/pull/7995): Enable autoplay in
   the preview app * Enhancement
   [owncloud/web#8008](https://github.com/owncloud/web/issues/8008): Don't open
   sidebar when copying quicklink * Enhancement
   [owncloud/web#8021](https://github.com/owncloud/web/pull/8021): Access right
   sidebar panels via URL * Enhancement
   [owncloud/web#8051](https://github.com/owncloud/web/pull/8051): Introduce image
   preloading to the preview app * Enhancement
   [owncloud/web#8055](https://github.com/owncloud/web/pull/8055): Retry failed
   uploads on re-upload * Enhancement
   [owncloud/web#8056](https://github.com/owncloud/web/pull/8056): Increase Searchbar
   height * Enhancement
   [owncloud/web#8057](https://github.com/owncloud/web/pull/8057): Show text file icon
   for empty text files * Enhancement
   [owncloud/web#8132](https://github.com/owncloud/web/pull/8132): Update
   libre-graph-api to v1.0 * Enhancement
   [owncloud/web#8136](https://github.com/owncloud/web/pull/8136): Make clipboard copy
   available to more browsers * Enhancement
   [owncloud/web#8161](https://github.com/owncloud/web/pull/8161): Space group members
   * Enhancement [owncloud/web#8161](https://github.com/owncloud/web/pull/8161): Space
   group shares * Enhancement
   [owncloud/web#8166](https://github.com/owncloud/web/issues/8166): Show upload speed
   * Enhancement [owncloud/web#8175](https://github.com/owncloud/web/pull/8175):
   Rename "user management" app * Enhancement
   [owncloud/web#8178](https://github.com/owncloud/web/pull/8178): Spaces list in admin
   settings * Enhancement
   [owncloud/web#8261](https://github.com/owncloud/web/pull/8261): Admin settings
   users section uses graph api for role assignments * Enhancement
   [owncloud/web#8279](https://github.com/owncloud/web/pull/8279): Move user group
   select to edit panel * Enhancement
   [owncloud/web#8280](https://github.com/owncloud/web/pull/8280): Add support for
   multiple clients in `theme.json` * Enhancement
   [owncloud/web#8294](https://github.com/owncloud/web/pull/8294): Move language
   selection to user account page * Enhancement
   [owncloud/web#8306](https://github.com/owncloud/web/pull/8306): Show selectable
   groups only * Enhancement
   [owncloud/web#8317](https://github.com/owncloud/web/pull/8317): Add context menu to
   groups * Enhancement
   [owncloud/web#8320](https://github.com/owncloud/web/pull/8320): Space member
   expiration * Enhancement
   [owncloud/web#8320](https://github.com/owncloud/web/pull/8320): Update SDK to
   v3.1.0-alpha.3 * Enhancement
   [owncloud/web#8324](https://github.com/owncloud/web/pull/8324): Add context menu to
   users * Enhancement [owncloud/web#8331](https://github.com/owncloud/web/pull/8331):
   Admin settings users section details improvement * Enhancement
   [owncloud/web#8354](https://github.com/owncloud/web/issues/8354): Add `ItemFilter`
   component * Enhancement
   [owncloud/web#8356](https://github.com/owncloud/web/pull/8356): Slight improvement
   of key up/down performance * Enhancement
   [owncloud/web#8363](https://github.com/owncloud/web/issues/8363): Admin settings
   general section * Enhancement
   [owncloud/web#8375](https://github.com/owncloud/web/pull/8375): Add appearance
   section in general settings * Enhancement
   [owncloud/web#8377](https://github.com/owncloud/web/issues/8377): User group filter
   * Enhancement [owncloud/web#8387](https://github.com/owncloud/web/pull/8387): Batch
   edit quota in admin panel * Enhancement
   [owncloud/web#8398](https://github.com/owncloud/web/pull/8398): Use standardized
   layout for file/space action list * Enhancement
   [owncloud/web#8425](https://github.com/owncloud/web/issues/8425): Add dark ownCloud
   logo * Enhancement [owncloud/web#8432](https://github.com/owncloud/web/pull/8432):
   Inject customizations * Enhancement
   [owncloud/web#8433](https://github.com/owncloud/web/pull/8433): User settings login
   field * Enhancement [owncloud/web#8441](https://github.com/owncloud/web/pull/8441):
   Skeleton App * Enhancement
   [owncloud/web#8449](https://github.com/owncloud/web/pull/8449): Configurable top
   bar * Enhancement [owncloud/web#8450](https://github.com/owncloud/web/pull/8450):
   Rework notification bell * Enhancement
   [owncloud/web#8455](https://github.com/owncloud/web/pull/8455): Autosave content
   changes in text editor * Enhancement
   [owncloud/web#8473](https://github.com/owncloud/web/pull/8473): Update CERN links *
   Enhancement [owncloud/web#8489](https://github.com/owncloud/web/pull/8489):
   Respect max quota * Enhancement
   [owncloud/web#8492](https://github.com/owncloud/web/pull/8492): User role filter *
   Enhancement [owncloud/web#8503](https://github.com/owncloud/web/issues/8503):
   Beautify file version list * Enhancement
   [owncloud/web#8515](https://github.com/owncloud/web/pull/8515): Introduce trashbin
   overview * Enhancement
   [owncloud/web#8518](https://github.com/owncloud/web/pull/8518): Make notifications
   work with oCIS * Enhancement
   [owncloud/web#8541](https://github.com/owncloud/web/pull/8541): Public link
   permission `PublicLink.Write.all` * Enhancement
   [owncloud/web#8553](https://github.com/owncloud/web/pull/8553): Add and remove users
   from groups batch actions * Enhancement
   [owncloud/web#8554](https://github.com/owncloud/web/pull/8554): Beautify form
   inputs * Enhancement
   [owncloud/web#8557](https://github.com/owncloud/web/issues/8557): Rework mobile
   navigation * Enhancement
   [owncloud/web#8566](https://github.com/owncloud/web/pull/8566): QuickActions role
   configurable * Enhancement
   [owncloud/web#8612](https://github.com/owncloud/web/issues/8612): Add
   `Accept-Language` header to all outgoing requests * Enhancement
   [owncloud/web#8630](https://github.com/owncloud/web/pull/8630): Add logout url *
   Enhancement [owncloud/web#8652](https://github.com/owncloud/web/pull/8652): Enable
   guest users * Enhancement
   [owncloud/web#8711](https://github.com/owncloud/web/pull/8711): Remove
   placeholder, add customizable label * Enhancement
   [owncloud/web#8713](https://github.com/owncloud/web/pull/8713): Context helper read
   more link configurable * Enhancement
   [owncloud/web#8715](https://github.com/owncloud/web/pull/8715): Enable rename
   groups * Enhancement
   [owncloud/web#8730](https://github.com/owncloud/web/pull/8730): Create Space from
   selection * Enhancement
   [owncloud/web#8738](https://github.com/owncloud/web/issues/8738): GDPR export *
   Enhancement [owncloud/web#8762](https://github.com/owncloud/web/pull/8762): Stop
   bootstrapping application earlier in anonymous contexts * Enhancement
   [owncloud/web#8766](https://github.com/owncloud/web/pull/8766): Add support for
   read-only groups * Enhancement
   [owncloud/web#8790](https://github.com/owncloud/web/pull/8790): Custom
   translations * Enhancement
   [owncloud/web#8797](https://github.com/owncloud/web/pull/8797): Font family in
   theming * Enhancement
   [owncloud/web#8806](https://github.com/owncloud/web/pull/8806): Preview app sorting
   * Enhancement [owncloud/web#8820](https://github.com/owncloud/web/pull/8820):
   Adjust missing reshare permissions message * Enhancement
   [owncloud/web#8822](https://github.com/owncloud/web/pull/8822): Fix quicklink icon
   alignment * Enhancement
   [owncloud/web#8826](https://github.com/owncloud/web/pull/8826): Admin settings
   groups members panel * Enhancement
   [owncloud/web#8868](https://github.com/owncloud/web/pull/8868): Respect user
   read-only configuration by the server * Enhancement
   [owncloud/web#8876](https://github.com/owncloud/web/pull/8876): Update roles and
   permissions names, labels, texts and icons * Enhancement
   [owncloud/web#8882](https://github.com/owncloud/web/pull/8882): Layout of Share role
   and expiration date dropdown * Enhancement
   [owncloud/web#8883](https://github.com/owncloud/web/issues/8883): Webfinger
   redirect app * Enhancement
   [owncloud/web#8898](https://github.com/owncloud/web/pull/8898): Rename "Quicklink"
   to "link" * Enhancement
   [owncloud/web#8911](https://github.com/owncloud/web/pull/8911): Add notification
   setting to account page * Enhancement
   [owncloud/web#9048](https://github.com/owncloud/web/issues/9048): Support
   pagination in admin settings app * Enhancement
   [owncloud/web#9070](https://github.com/owncloud/web/pull/9070): Disable change
   password capability * Enhancement
   [owncloud/web#9070](https://github.com/owncloud/web/pull/9070): Disable create user
   and delete user via capabilities * Enhancement
   [owncloud/web#9076](https://github.com/owncloud/web/pull/9076): Show detailed error
   messages while upload fails

   https://github.com/owncloud/ocis/pull/6438
   https://github.com/owncloud/web/releases/tag/v7.0.0

* Enhancement - Use Accept-Language Header: [#5918](https://github.com/owncloud/ocis/pull/5918)

   Use the `Accept-Language` header instead of the custom `Prefered-Language`

   https://github.com/owncloud/ocis/pull/5918

* Enhancement - Use gotext master: [#5867](https://github.com/owncloud/ocis/pull/5867)

   We needed to use forked version until our upstream changes were merged

   https://github.com/owncloud/ocis/pull/5867

* Enhancement - Userlog: [#5699](https://github.com/owncloud/ocis/pull/5699)

   Enhance userlog service with proper api and messages

   https://github.com/owncloud/ocis/pull/5699

* Enhancement - Userlog Service: [#5610](https://github.com/owncloud/ocis/pull/5610)

   Introduces userlog service. It stores eventIDs the user is interested in and provides an API to
   retrieve the events.

   https://github.com/owncloud/ocis/pull/5610

* Enhancement - Determine the users language to translate via Transifex: [#6089](https://github.com/owncloud/ocis/pull/6089)

   https://github.com/owncloud/ocis/issues/6087
   https://github.com/owncloud/ocis/pull/6089
   Enhance
   userlog
   service
   with
   proper
   api
   and
   messages

* Enhancement - Web options configuration: [#6188](https://github.com/owncloud/ocis/pull/6188)

   Hardcode web options instead of using a generic `map[string]interface{}`

   https://github.com/owncloud/ocis/pull/6188
# Changelog for [2.0.0] (2022-11-30)

The following sections list the changes for 2.0.0.

[2.0.0]: https://github.com/owncloud/ocis/compare/v1.20.0...v2.0.0

## Summary

* Bugfix - Fix configuration of mimetypes for the app registry: [#4411](https://github.com/owncloud/ocis/pull/4411)
* Bugfix - Disable default expiration for public links: [#4445](https://github.com/owncloud/ocis/issues/4445)
* Bugfix - Show help for some commands when unconfigured: [#4405](https://github.com/owncloud/ocis/pull/4405)
* Bugfix - Translations on login page: [#7550](https://github.com/owncloud/web/issues/7550)
* Bugfix - Autocreate IDP private key also if file exists but is empty: [#4394](https://github.com/owncloud/ocis/pull/4394)
* Bugfix - Rename extensions to services (leftover occurrences): [#4407](https://github.com/owncloud/ocis/pull/4407)
* Bugfix - Fix DN parsing issues and sizelimit handling in libregraph/idm: [#3631](https://github.com/owncloud/ocis/issues/3631)
* Bugfix - Lower IDP token lifespans: [#5077](https://github.com/owncloud/ocis/pull/5077)
* Bugfix - Remove runtime kill and run commands: [#3740](https://github.com/owncloud/ocis/pull/3740)
* Bugfix - Check permissions when deleting Space: [#3709](https://github.com/owncloud/ocis/pull/3709)
* Bugfix - Do not reindex a space twice at the same time: [#5001](https://github.com/owncloud/ocis/pull/5001)
* Bugfix - Disable federation capabilities: [#4864](https://github.com/owncloud/ocis/pull/4864)
* Bugfix - Decomposedfs increase filelock duration factor: [#5130](https://github.com/owncloud/ocis/pull/5130)
* Bugfix - Find spaces by their name: [#5044](https://github.com/owncloud/ocis/pull/5044)
* Bugfix - Logging in on the wrong account when an email address is not unique: [#4039](https://github.com/owncloud/ocis/issues/4039)
* Bugfix - Allow empty environment variables: [#3892](https://github.com/owncloud/ocis/pull/3892)
* Bugfix - Remove unused transfer secret from app provider: [#3798](https://github.com/owncloud/ocis/pull/3798)
* Bugfix - Fix authentication for autoprovisioned users: [#4616](https://github.com/owncloud/ocis/issues/4616)
* Bugfix - Bring back the settings UI in Web: [#4691](https://github.com/owncloud/ocis/pull/4691)
* Bugfix - Fix cache stat table config: [#4732](https://github.com/owncloud/ocis/pull/4732)
* Bugfix - Adjust cache related configuration options: [#5087](https://github.com/owncloud/ocis/pull/5087)
* Bugfix - Make IDP secrets configurable via environment variables: [#3744](https://github.com/owncloud/ocis/pull/3744)
* Bugfix - CSP rules for silent token refresh in iframe: [#4031](https://github.com/owncloud/ocis/pull/4031)
* Bugfix - Enable debug server by default: [#3827](https://github.com/owncloud/ocis/pull/3827)
* Bugfix - Rework default role provisioning: [#3900](https://github.com/owncloud/ocis/issues/3900)
* Bugfix - Fix search index getting out of sync: [#3851](https://github.com/owncloud/ocis/pull/3851)
* Bugfix - Change the default value for PROXY_OIDC_INSECURE to false: [#4601](https://github.com/owncloud/ocis/pull/4601)
* Bugfix - Fix sharing jsoncs3 driver options: [#4593](https://github.com/owncloud/ocis/pull/4593)
* Bugfix - Inconsistency env var naming for LDAP filter configuration: [#3890](https://github.com/owncloud/ocis/issues/3890)
* Bugfix - Fix LDAP insecure options: [#3897](https://github.com/owncloud/ocis/pull/3897)
* Bugfix - Fix handling of invalid LDAP users and groups: [#4274](https://github.com/owncloud/ocis/issues/4274)
* Bugfix - Fix logging levels: [#4102](https://github.com/owncloud/ocis/pull/4102)
* Bugfix - Don't run auth-bearer service by default: [#4692](https://github.com/owncloud/ocis/issues/4692)
* Bugfix - Fix notifications service settings: [#4652](https://github.com/owncloud/ocis/pull/4652)
* Bugfix - Fix notifications Web UI url: [#4998](https://github.com/owncloud/ocis/pull/4998)
* Bugfix - Fix `OCIS_RUN_SERVICES`: [#4133](https://github.com/owncloud/ocis/pull/4133)
* Bugfix - Fix the OIDC provider cache: [#4600](https://github.com/owncloud/ocis/pull/4600)
* Bugfix - Fix permissions in REPORT: [#4520](https://github.com/owncloud/ocis/pull/4520)
* Bugfix - Set default name for public link via capabilities: [#3834](https://github.com/owncloud/ocis/pull/3834)
* Bugfix - Remove legacy accounts proxy routes: [#3831](https://github.com/owncloud/ocis/pull/3831)
* Bugfix - Fix unused config option `GRAPH_SPACES_INSECURE`: [#55555](https://github.com/owncloud/ocis/pull/55555)
* Bugfix - Remove unused configuration options: [#3973](https://github.com/owncloud/ocis/pull/3973)
* Bugfix - Remove static ocs user backend config: [#4077](https://github.com/owncloud/ocis/pull/4077)
* Bugfix - Remove unused OCS storage configuration: [#3955](https://github.com/owncloud/ocis/pull/3955)
* Bugfix - Fix the `ocis search` command: [#3796](https://github.com/owncloud/ocis/pull/3796)
* Bugfix - Rename search env variable for the grpc server address: [#3800](https://github.com/owncloud/ocis/pull/3800)
* Bugfix - Fix search in received shares: [#4308](https://github.com/owncloud/ocis/issues/4308)
* Bugfix - Fix search report: [#7557](https://github.com/owncloud/web/issues/7557)
* Bugfix - Render webdav permissions as string in search report: [#4575](https://github.com/owncloud/ocis/issues/4575)
* Bugfix - Fix make sensitive config values in the proxy's debug server: [#4086](https://github.com/owncloud/ocis/pull/4086)
* Bugfix - Fix the idm and settings extensions' admin user id configuration option: [#3799](https://github.com/owncloud/ocis/pull/3799)
* Bugfix - Mail notifications for group shares: [#4714](https://github.com/owncloud/ocis/pull/4714)
* Bugfix - Substring search for sharees: [#547](https://github.com/owncloud/ocis/issues/547)
* Bugfix - Fix configuration validation for extensions' server commands: [#3911](https://github.com/owncloud/ocis/pull/3911)
* Bugfix - Fix startup error logging: [#4093](https://github.com/owncloud/ocis/pull/4093)
* Bugfix - Disable cache for selected static web assets: [#4809](https://github.com/owncloud/ocis/pull/4809)
* Bugfix - Fix multiple storage-users env variables: [#3802](https://github.com/owncloud/ocis/pull/3802)
* Bugfix - Thumbnails for `/dav/xxx?preview=1` requests: [#3567](https://github.com/owncloud/ocis/pull/3567)
* Bugfix - Fix unfindable entities from shares/publicshares: [#4651](https://github.com/owncloud/ocis/pull/4651)
* Bugfix - Fix unrestricted quota on the graphAPI: [#4363](https://github.com/owncloud/ocis/pull/4363)
* Bugfix - Fix user autoprovisioning: [#3893](https://github.com/owncloud/ocis/issues/3893)
* Bugfix - Fix version info: [#3953](https://github.com/owncloud/ocis/pull/3953)
* Bugfix - Fix version number in status page: [#3788](https://github.com/owncloud/ocis/issues/3788)
* Bugfix - Fix CORS in frontend service: [#4948](https://github.com/owncloud/ocis/pull/4948)
* Bugfix - Graph service now forwards trace context: [#4582](https://github.com/owncloud/ocis/pull/4582)
* Bugfix - Fix the webdav URL of drive roots: [#3706](https://github.com/owncloud/ocis/issues/3706)
* Bugfix - Idp: Check if CA certificate if present: [#3623](https://github.com/owncloud/ocis/issues/3623)
* Bugfix - Fix graph endpoint: [#3925](https://github.com/owncloud/ocis/issues/3925)
* Bugfix - Initial role assignment with external IDM: [#5045](https://github.com/owncloud/ocis/issues/5045)
* Bugfix - Escape DN attribute value: [#4117](https://github.com/owncloud/ocis/pull/4117)
* Bugfix - Make IDP only wait for certs when using LDAP: [#3965](https://github.com/owncloud/ocis/pull/3965)
* Bugfix - Make ocdav service behave properly: [#3957](https://github.com/owncloud/ocis/pull/3957)
* Bugfix - Make storage users mount ids unique by default: [#5091](https://github.com/owncloud/ocis/pull/5091)
* Bugfix - Return proper errors when ocs/cloud/users is using the cs3 backend: [#3483](https://github.com/owncloud/ocis/issues/3483)
* Bugfix - Polish search: [#4094](https://github.com/owncloud/ocis/pull/4094)
* Bugfix - Fix the shareroot path in REPORT responses: [#4859](https://github.com/owncloud/ocis/pull/4859)
* Bugfix - Remove the storage-users event configuration: [#4825](https://github.com/owncloud/ocis/pull/4825)
* Bugfix - Trigger a rescan of spaces in the search index when items have changed: [#4777](https://github.com/owncloud/ocis/pull/4777)
* Bugfix - Save Katherine: [#3823](https://github.com/owncloud/ocis/issues/3823)
* Bugfix - Fix permission check in settings service: [#4890](https://github.com/owncloud/ocis/pull/4890)
* Bugfix - Fix Thumbnails for IDs without a trailing path: [#3791](https://github.com/owncloud/ocis/pull/3791)
* Bugfix - Space Creators can hand over spaces: [#4244](https://github.com/owncloud/ocis/pull/4244)
* Bugfix - Make tokeninfo endpoint unprotected: [#4715](https://github.com/owncloud/ocis/pull/4715)
* Bugfix - Update reva to version 2.12.0: [#5092](https://github.com/owncloud/ocis/pull/5092)
* Bugfix - URL encode the webdav url in the graph API: [#3597](https://github.com/owncloud/ocis/pull/3597)
* Bugfix - Store user passwords hashed in idm: [#3778](https://github.com/owncloud/ocis/issues/3778)
* Bugfix - Fix wopi access to public shares: [#4631](https://github.com/owncloud/ocis/pull/4631)
* Change - Update ocis packages and imports to V2: [#3678](https://github.com/owncloud/ocis/pull/3678)
* Change - Build service frontends with pnpm instead of yarn: [#4878](https://github.com/owncloud/ocis/pull/4878)
* Change - Load configuration files just from one directory: [#3587](https://github.com/owncloud/ocis/pull/3587)
* Change - Reduce permissions on docker image predeclared volumes: [#3641](https://github.com/owncloud/ocis/pull/3641)
* Change - Introduce `ocis init` and remove all default secrets: [#3551](https://github.com/owncloud/ocis/pull/3551)
* Change - Rename "uploads purge" command to "uploads clean": [#4403](https://github.com/owncloud/ocis/pull/4403)
* Change - Enable private links by default: [#4599](https://github.com/owncloud/ocis/pull/4599/)
* Change - The `glauth` and `accounts` services are removed: [#3685](https://github.com/owncloud/ocis/pull/3685)
* Change - Reduce drives in graph /me/drives API: [#3629](https://github.com/owncloud/ocis/pull/3629)
* Change - Switched default configuration to use libregraph/idm: [#3331](https://github.com/owncloud/ocis/pull/3331)
* Change - Rename MetadataUserID: [#3671](https://github.com/owncloud/ocis/pull/3671)
* Change - Use new space ID util functions: [#3648](https://github.com/owncloud/ocis/pull/3648)
* Change - Prevent access to disabled space: [#3779](https://github.com/owncloud/ocis/pull/3779)
* Change - Rename serviceUser to systemUser: [#3673](https://github.com/owncloud/ocis/pull/3673)
* Change - Use the spaceID on the cs3 resource: [#4748](https://github.com/owncloud/ocis/pull/4748)
* Change - Split MachineAuth from SystemUser: [#3672](https://github.com/owncloud/ocis/pull/3672)
* Enhancement - Add capability for alias links: [#3983](https://github.com/owncloud/ocis/issues/3983)
* Enhancement - Add curl to the oCIS OCI image: [#4751](https://github.com/owncloud/ocis/pull/4751)
* Enhancement - Add deprecation annotation: [#3917](https://github.com/owncloud/ocis/issues/3917)
* Enhancement - Add drives field to users endpoint: [#4072](https://github.com/owncloud/ocis/pull/4072)
* Enhancement - Add Email templating: [#4564](https://github.com/owncloud/ocis/pull/4564)
* Enhancement - Add FRONTEND_ENABLE_RESHARING env variable: [#4023](https://github.com/owncloud/ocis/pull/4023)
* Enhancement - We added e-mail subject templating: [#4799](https://github.com/owncloud/ocis/pull/4799)
* Enhancement - Add number of total matches to the search result: [#4189](https://github.com/owncloud/ocis/issues/4189)
* Enhancement - Add tracing to search: [#5113](https://github.com/owncloud/ocis/pull/5113)
* Enhancement - Add webURL to space root: [#4588](https://github.com/owncloud/ocis/pull/4588)
* Enhancement - Align service naming: [#3606](https://github.com/owncloud/ocis/pull/3606)
* Enhancement - Add acting user to the audit log: [#3753](https://github.com/owncloud/ocis/issues/3753)
* Enhancement - Configurable max lock cycles: [#4965](https://github.com/owncloud/ocis/pull/4965)
* Enhancement - Allow to configuring the reva cache store: [#4627](https://github.com/owncloud/ocis/pull/4627)
* Enhancement - Add audit events for created containers: [#3941](https://github.com/owncloud/ocis/pull/3941)
* Enhancement - Add support for REPORT requests to /dav/spaces URLs: [#4661](https://github.com/owncloud/ocis/pull/4661)
* Enhancement - Don't setup demo role assignments on default: [#3661](https://github.com/owncloud/ocis/issues/3661)
* Enhancement - Introduce "delete-all-spaces" permission: [#4196](https://github.com/owncloud/ocis/issues/4196)
* Enhancement - Deny access to resources: [#4903](https://github.com/owncloud/ocis/pull/4903)
* Enhancement - Improve validation of OIDC access tokens: [#3841](https://github.com/owncloud/ocis/issues/3841)
* Enhancement - Add /app/open-with-web endpoint: [#4376](https://github.com/owncloud/ocis/pull/4376)
* Enhancement - Add previewFileMimeTypes to web default config: [#4414](https://github.com/owncloud/ocis/pull/4414)
* Enhancement - Added language option to the app provider: [#4399](https://github.com/owncloud/ocis/pull/4399)
* Enhancement - Improve error log for "could not get user by claim" error: [#4227](https://github.com/owncloud/ocis/pull/4227)
* Enhancement - Improve login screen design: [#4500](https://github.com/owncloud/ocis/pull/4500)
* Enhancement - Add configuration options for mail authentication and encryption: [#4443](https://github.com/owncloud/ocis/pull/4443)
* Enhancement - Introduce service registry cache: [#3833](https://github.com/owncloud/ocis/pull/3833)
* Enhancement - Reintroduce user autoprovisioning in proxy: [#3860](https://github.com/owncloud/ocis/pull/3860)
* Enhancement - Allow to configure applications in Web: [#4578](https://github.com/owncloud/ocis/pull/4578)
* Enhancement - Added command to reset administrator password: [#4084](https://github.com/owncloud/ocis/issues/4084)
* Enhancement - Disable the color logging in docker compose examples: [#871](https://github.com/owncloud/ocis/issues/871)
* Enhancement - Allow providing list of services NOT to start: [#4254](https://github.com/owncloud/ocis/pull/4254)
* Enhancement - Introduce insecure flag for smtp email notifications: [#4279](https://github.com/owncloud/ocis/pull/4279)
* Enhancement - Optional events in graph service: [#55555](https://github.com/owncloud/ocis/pull/55555)
* Enhancement - Fix behavior for foobar (in present tense): [#4346](https://github.com/owncloud/ocis/pull/4346)
* Enhancement - Add the "hidden" state to the search index: [#5018](https://github.com/owncloud/ocis/pull/5018)
* Enhancement - Restrict admins from self-removal: [#3713](https://github.com/owncloud/ocis/issues/3713)
* Enhancement - OCS get share now also handle received shares: [#4322](https://github.com/owncloud/ocis/issues/4322)
* Enhancement - Add config option to provide TLS certificate: [#3818](https://github.com/owncloud/ocis/issues/3818)
* Enhancement - Add descriptions for graph-explorer config: [#3759](https://github.com/owncloud/ocis/pull/3759)
* Enhancement - Add /me/changePassword endpoint to GraphAPI: [#3063](https://github.com/owncloud/ocis/issues/3063)
* Enhancement - Allow to setup TLS for grpc services: [#4798](https://github.com/owncloud/ocis/pull/4798)
* Enhancement - Generate signing key and encryption secret: [#3909](https://github.com/owncloud/ocis/issues/3909)
* Enhancement - Update IdP UI: [#3493](https://github.com/owncloud/ocis/issues/3493)
* Enhancement - Logging improvements: [#4815](https://github.com/owncloud/ocis/pull/4815)
* Enhancement - Wrap metadata storage with dedicated reva gateway: [#3602](https://github.com/owncloud/ocis/pull/3602)
* Enhancement - New migrate command for migrating shares and public shares: [#3987](https://github.com/owncloud/ocis/pull/3987)
* Enhancement - Default to tls 1.2: [#4969](https://github.com/owncloud/ocis/pull/4969)
* Enhancement - Add missing unprotected paths: [#4454](https://github.com/owncloud/ocis/pull/4454)
* Enhancement - Secure the nats connection with TLS: [#4781](https://github.com/owncloud/ocis/pull/4781)
* Enhancement - Product field in OCS version: [#2918](https://github.com/owncloud/ocis/pull/2918)
* Enhancement - Automatically orientate photos when generating thumbnails: [#4477](https://github.com/owncloud/ocis/issues/4477)
* Enhancement - Refactor extensions to services: [#3980](https://github.com/owncloud/ocis/pull/3980)
* Enhancement - Refactor the proxy service: [#4401](https://github.com/owncloud/ocis/issues/4401)
* Enhancement - Remove windows from ci & release makefile: [#5026](https://github.com/owncloud/ocis/pull/5026)
* Enhancement - Rename AUTH_BASIC_AUTH_PROVIDER envvar: [#4966](https://github.com/owncloud/ocis/pull/4966)
* Enhancement - Report parent id: [#4757](https://github.com/owncloud/ocis/pull/4757)
* Enhancement - Allow resharing: [#3904](https://github.com/owncloud/ocis/pull/3904)
* Enhancement - Rewrite of the request authentication middleware: [#4374](https://github.com/owncloud/ocis/pull/4374)
* Enhancement - Add initial version of the search extensions: [#3635](https://github.com/owncloud/ocis/pull/3635)
* Enhancement - Prohibit users from setting or listing other user's values: [#4897](https://github.com/owncloud/ocis/pull/4897)
* Enhancement - Add capability for public link single file edit: [#6787](https://github.com/owncloud/web/pull/6787)
* Enhancement - Added `share_jail` and `projects` feature flags in spaces capability: [#3626](https://github.com/owncloud/ocis/pull/3626)
* Enhancement - Use storageID when requesting special items: [#4356](https://github.com/owncloud/ocis/pull/4356)
* Enhancement - Add description tags to the thumbnails config structs: [#3752](https://github.com/owncloud/ocis/pull/3752)
* Enhancement - Make thumbnails service log less noisy: [#3959](https://github.com/owncloud/ocis/pull/3959)
* Enhancement - Add thumbnails support for tiff and bmp files: [#4634](https://github.com/owncloud/ocis/pull/4634)
* Enhancement - Update linkshare capabilities: [#3579](https://github.com/owncloud/ocis/pull/3579)
* Enhancement - Update reva: [#3944](https://github.com/owncloud/ocis/pull/3944)
* Enhancement - Update reva to version 2.7.2: [#4115](https://github.com/owncloud/ocis/pull/4115)
* Enhancement - Update reva to v2.7.4: [#4294](https://github.com/owncloud/ocis/pull/4294)
* Enhancement - Update reva to v2.8.0: [#4444](https://github.com/owncloud/ocis/pull/4444)
* Enhancement - Update reva to version 2.4.1: [#3746](https://github.com/owncloud/ocis/pull/3746)
* Enhancement - Update reva to version 2.5.1: [#3932](https://github.com/owncloud/ocis/pull/3932)
* Enhancement - Update Reva to version 2.10.0: [#4522](https://github.com/owncloud/ocis/pull/4522)
* Enhancement - Update reva to version 2.11.0: [#4588](https://github.com/owncloud/ocis/pull/4588)
* Enhancement - Update reva to v2.3.1: [#3552](https://github.com/owncloud/ocis/pull/3552)
* Enhancement - Update ownCloud Web to v5.5.0-rc.8: [#6854](https://github.com/owncloud/web/pull/6854)
* Enhancement - Update ownCloud Web to v5.5.0-rc.9: [#6854](https://github.com/owncloud/web/pull/6854)
* Enhancement - Update ownCloud Web to v5.5.0-rc.6: [#6854](https://github.com/owncloud/web/pull/6854)
* Enhancement - Update ownCloud Web to v5.7.0-rc.1: [#4005](https://github.com/owncloud/ocis/pull/4005)
* Enhancement - Update ownCloud Web to v6.0.0: [#5153](https://github.com/owncloud/ocis/pull/5153)
* Enhancement - Update ownCloud Web to v5.7.0-rc.4: [#4140](https://github.com/owncloud/ocis/pull/4140)
* Enhancement - Update ownCloud Web to v5.7.0-rc.8: [#4314](https://github.com/owncloud/ocis/pull/4314)
* Enhancement - Update ownCloud Web to v5.7.0-rc.10: [#4439](https://github.com/owncloud/ocis/pull/4439)
* Enhancement - Update ownCloud Web to v5.7.0: [#4508](https://github.com/owncloud/ocis/pull/4508)
* Enhancement - Expand personal drive on the graph user: [#4357](https://github.com/owncloud/ocis/pull/4357)
* Enhancement - Validate space names: [#4955](https://github.com/owncloud/ocis/pull/4955)
* Enhancement - Add descriptions to webdav configuration: [#3755](https://github.com/owncloud/ocis/pull/3755)
* Enhancement - Search service at the old webdav endpoint: [#4118](https://github.com/owncloud/ocis/pull/4118)
* Enhancement - Make it possible to configure a WOPI folderurl: [#4716](https://github.com/owncloud/ocis/pull/4716)

## Details

* Bugfix - Fix configuration of mimetypes for the app registry: [#4411](https://github.com/owncloud/ocis/pull/4411)

   We've fixed the configuration option for mimetypes in the app registry. Previously the
   default config would always be merged over the user provided configuration. Now the default
   mimetype configuration is only used if the user does not provide any mimetype configuration
   (like it is already done in the proxy with the routes configuration).

   https://github.com/owncloud/ocis/pull/4411

* Bugfix - Disable default expiration for public links: [#4445](https://github.com/owncloud/ocis/issues/4445)

   The default expiration for public links was enabled in the capabilities without providing a
   (then required) default amount of days for clients to pick a reasonable expiration date upon
   link creation. This has been fixed by disabling the default expiration for public links in the
   capabilities. With this configuration clients will no longer set a default expiration date
   upon link creation.

   https://github.com/owncloud/ocis/issues/4445
   https://github.com/owncloud/ocis/pull/4475

* Bugfix - Show help for some commands when unconfigured: [#4405](https://github.com/owncloud/ocis/pull/4405)

   We've fixed some commands to show the help also when oCIS is not yet configured. Previously the
   help was not displayed to the user but instead a configuration validation error.

   https://github.com/owncloud/ocis/pull/4405

* Bugfix - Translations on login page: [#7550](https://github.com/owncloud/web/issues/7550)

   We've fixed several translations on the login page. Also, the browser language is now being
   used properly to determine the language.

   https://github.com/owncloud/web/issues/7550
   https://github.com/owncloud/ocis/pull/4504

* Bugfix - Autocreate IDP private key also if file exists but is empty: [#4394](https://github.com/owncloud/ocis/pull/4394)

   We've fixed the behavior for the IDP private key generation so that a private key is also
   generated when the file already exists but is empty.

   https://github.com/owncloud/ocis/pull/4394

* Bugfix - Rename extensions to services (leftover occurrences): [#4407](https://github.com/owncloud/ocis/pull/4407)

   We've already renamed extensions to services in previous PRs and this PR performs this rename
   for leftover occurrences.

   https://github.com/owncloud/ocis/pull/4407

* Bugfix - Fix DN parsing issues and sizelimit handling in libregraph/idm: [#3631](https://github.com/owncloud/ocis/issues/3631)

   We fixed a couple on issues in libregraph/idm related to correctly parsing LDAP DNs for
   usernames contain characters that require escaping.

   Also libregraph/idm was not properly returning "Size limit exceeded" errors when the result
   set exceeded the requested size.

   https://github.com/owncloud/ocis/issues/3631
   https://github.com/owncloud/ocis/issues/4039
   https://github.com/owncloud/ocis/issues/4078

* Bugfix - Lower IDP token lifespans: [#5077](https://github.com/owncloud/ocis/pull/5077)

   We've lowered the IDP token lifespans to more reasonable durations.

   https://github.com/owncloud/ocis/pull/5077

* Bugfix - Remove runtime kill and run commands: [#3740](https://github.com/owncloud/ocis/pull/3740)

   We've removed the kill and run commands from the oCIS runtime. If these dynamic capabilities
   are needed, one should switch to a full fledged supervisor and start oCIS as individual
   services.

   If one wants to start a only a subset of services, this is still possible by setting
   OCIS_RUN_EXTENSIONS.

   https://github.com/owncloud/ocis/pull/3740

* Bugfix - Check permissions when deleting Space: [#3709](https://github.com/owncloud/ocis/pull/3709)

   Check for manager permissions when deleting spaces. Do not allow deleting spaces via dav
   service

   https://github.com/owncloud/ocis/pull/3709

* Bugfix - Do not reindex a space twice at the same time: [#5001](https://github.com/owncloud/ocis/pull/5001)

   We fixed a problem where the search service reindexed a space while another reindex process was
   still in progress.

   https://github.com/owncloud/ocis/pull/5001

* Bugfix - Disable federation capabilities: [#4864](https://github.com/owncloud/ocis/pull/4864)

   We disabled the federation support in the capabilities because it is currently not supported.

   https://github.com/owncloud/ocis/pull/4864

* Bugfix - Decomposedfs increase filelock duration factor: [#5130](https://github.com/owncloud/ocis/pull/5130)

   We made the file lock duration per lock cycle for decomposedfs configurable and increased it to
   make locks work on top of NFS.

   https://github.com/owncloud/ocis/issues/5024
   https://github.com/owncloud/ocis/pull/5130

* Bugfix - Find spaces by their name: [#5044](https://github.com/owncloud/ocis/pull/5044)

   We've fixed finding spaces by their name in the search service.

   https://github.com/owncloud/ocis/issues/4506
   https://github.com/owncloud/ocis/pull/5044

* Bugfix - Logging in on the wrong account when an email address is not unique: [#4039](https://github.com/owncloud/ocis/issues/4039)

   The default configuration to use the same logon attribute for all services. Also, if the
   configured logon attribute is not unique access to ocis is denied.

   https://github.com/owncloud/ocis/issues/4039

* Bugfix - Allow empty environment variables: [#3892](https://github.com/owncloud/ocis/pull/3892)

   We've fixed the behavior for empty environment variables, that previously would not have
   overwritten default values. Therefore it had the same effect like not setting the environment
   variable. We now check if the environment variable is set at all and if so, we also allow to
   override a default value with an empty value.

   https://github.com/owncloud/ocis/pull/3892

* Bugfix - Remove unused transfer secret from app provider: [#3798](https://github.com/owncloud/ocis/pull/3798)

   We've fixed the startup of the app provider by removing the startup dependency on a configured
   transfer secret, which was not used. This only happened if you start the app provider without
   runtime (eg. `ocis app-provider server`) and didn't have configured all oCIS secrets.

   https://github.com/owncloud/ocis/pull/3798

* Bugfix - Fix authentication for autoprovisioned users: [#4616](https://github.com/owncloud/ocis/issues/4616)

   We've fixed an issue in the proxy, which made the first http request of an autoprovisioned user
   fail.

   https://github.com/owncloud/ocis/issues/4616

* Bugfix - Bring back the settings UI in Web: [#4691](https://github.com/owncloud/ocis/pull/4691)

   We've fixed the oC Web configuration in oCIS so that the settings UI will be shown again in Web.

   https://github.com/owncloud/ocis/pull/4691

* Bugfix - Fix cache stat table config: [#4732](https://github.com/owncloud/ocis/pull/4732)

   We have aligned the cache table config for the gateway and the dataprovider to make them
   actually use the same cache instance.

   https://github.com/owncloud/ocis/pull/4732

* Bugfix - Adjust cache related configuration options: [#5087](https://github.com/owncloud/ocis/pull/5087)

   We've adjusted cache related configuration options of the gateway and storage-users service
   to the other services.

   https://github.com/owncloud/ocis/pull/5087

* Bugfix - Make IDP secrets configurable via environment variables: [#3744](https://github.com/owncloud/ocis/pull/3744)

   We've fixed the configuration options of the IDP to make the IDP secrets again configurable via
   environment variables.

   https://github.com/owncloud/ocis/pull/3744

* Bugfix - CSP rules for silent token refresh in iframe: [#4031](https://github.com/owncloud/ocis/pull/4031)

   When renewing the access token silently web needs to be opened in an iframe. This was previously
   blocked by a restrictive iframe CSP rule in the `Secure` middleware and has now been fixed by
   allow `self` for iframes.

   https://github.com/owncloud/web/issues/7030
   https://github.com/owncloud/ocis/pull/4031

* Bugfix - Enable debug server by default: [#3827](https://github.com/owncloud/ocis/pull/3827)

   We've fixed the behavior for the audit, idm, nats and notifications extensions, that did not
   start their debug server by default.

   https://github.com/owncloud/ocis/pull/3827

* Bugfix - Rework default role provisioning: [#3900](https://github.com/owncloud/ocis/issues/3900)

   We fixed a race condition in the default role assignment code that could lead to users loosing
   privileges. When authenticating before the settings service was fully running.

   https://github.com/owncloud/ocis/issues/3900

* Bugfix - Fix search index getting out of sync: [#3851](https://github.com/owncloud/ocis/pull/3851)

   We fixed a problem where the search index got out of sync with child elements of a parent
   containing special characters.

   https://github.com/owncloud/ocis/pull/3851

* Bugfix - Change the default value for PROXY_OIDC_INSECURE to false: [#4601](https://github.com/owncloud/ocis/pull/4601)

   We've changed the default value for PROXY_OIDC_INSECURE to `false`. Previously the default
   values was `true` which is not acceptable since default values need to be secure.

   https://github.com/owncloud/ocis/pull/4601

* Bugfix - Fix sharing jsoncs3 driver options: [#4593](https://github.com/owncloud/ocis/pull/4593)

   We've fixed the environment variable config options of the jsoncs3 driver that previously
   used the same environment variables as the cs3 driver. Now the jsoncs3 driver has it's own
   configuration environment variables.

   If you used the jsoncs3 sharing driver and explicitly set
   `SHARING_PUBLIC_CS3_SYSTEM_USER_ID`, this PR is a breaking change for your deployment. To
   workaround you may set the value you had configured in `SHARING_PUBLIC_CS3_SYSTEM_USER_ID`
   to both `SHARING_PUBLIC_JSONCS3_SYSTEM_USER_ID` and
   `SHARING_PUBLIC_JSONCS3_SYSTEM_USER_IDP`.

   https://github.com/owncloud/ocis/pull/4593

* Bugfix - Inconsistency env var naming for LDAP filter configuration: [#3890](https://github.com/owncloud/ocis/issues/3890)

   There was a naming inconsistency for the environment variables used to define LDAP filters for
   user and groups queries. Some services used `LDAP_USER_FILTER` while others used
   `LDAP_USERFILTER`. This is now changed to use `LDAP_USER_FILTER` and `LDAP_GROUP_FILTER`.

   Note: If your oCIS setup is using an LDAP configuration that has any of the `*_LDAP_USERFILTER`
   or `*_LDAP_GROUPFILTER` environment variables set, please update the configuration to use
   the new unified names `*_LDAP_USER_FILTER` respectively `*_LDAP_GROUP_FILTER` instead.

   https://github.com/owncloud/ocis/issues/3890

* Bugfix - Fix LDAP insecure options: [#3897](https://github.com/owncloud/ocis/pull/3897)

   We've fixed multiple LDAP insecure options:

  * The Graph LDAP insecure option default was set to `true` and now defaults to `false`. This is possible after #3888, since the Graph also now uses the LDAP CAcert by default.
  * The Graph LDAP insecure option was configurable by the environment variable `OCIS_INSECURE`, which was replaced by the dedicated `LDAP_INSECURE` variable. This variable is also used by all other services using LDAP.
  * The IDP insecure option for the user backend now also picks up configuration from `LDAP_INSECURE`.

   https://github.com/owncloud/ocis/pull/3897

* Bugfix - Fix handling of invalid LDAP users and groups: [#4274](https://github.com/owncloud/ocis/issues/4274)

   We fixed an issue where ocis would exit with a panic when LDAP users or groups where missing
   required attributes (e.g. the id)

   https://github.com/owncloud/ocis/issues/4274

* Bugfix - Fix logging levels: [#4102](https://github.com/owncloud/ocis/pull/4102)

   We've fixed the configuration of logging levels. Previously it was not possible to configure a
   service with a more or less verbose log level then all other services when running in the
   supervised / runtime mode `ocis server`.

   For example `OCIS_LOG_LEVEL=error PROXY_LOG_LEVEL=debug ocis server` did not configure
   error logging for all services except the proxy, which should be on debug logging. This is now
   fixed and working properly.

   Also we fixed the format of go-micro logs to always default to error level. Previously this was
   only ensured in the supervised / runtime mode.

   https://github.com/owncloud/ocis/issues/4089
   https://github.com/owncloud/ocis/pull/4102

* Bugfix - Don't run auth-bearer service by default: [#4692](https://github.com/owncloud/ocis/issues/4692)

   We no longer start the auth-bearer service by default. This service is currently unused and not
   required to run ocis. The equivalent functionality to verify OpenID connect tokens and to mint
   reva tokes for OIDC authenticated clients is currently implemented inside the oidc-auth
   middleware of the proxy.

   https://github.com/owncloud/ocis/issues/4692

* Bugfix - Fix notifications service settings: [#4652](https://github.com/owncloud/ocis/pull/4652)

   We've fixed two notifications service setting: - `NOTIFICATIONS_MACHINE_AUTH_API_KEY`
   was previously not picked up (only `OCIS_MACHINE_AUTH_API_KEY` was loaded) - If you used a
   email sender address in the format of the default value of `NOTIFICATIONS_SMTP_SENDER` no
   email could be send.

   https://github.com/owncloud/ocis/pull/4652

* Bugfix - Fix notifications Web UI url: [#4998](https://github.com/owncloud/ocis/pull/4998)

   We've fixed the configuration of the notification service's Web UI url that appears in emails.

   Previously it was only configurable via the global "OCIS_URL" and is now also configurable via
   "NOTIFICATIONS_WEB_UI_URL".

   https://github.com/owncloud/ocis/pull/4998

* Bugfix - Fix `OCIS_RUN_SERVICES`: [#4133](https://github.com/owncloud/ocis/pull/4133)

   `OCIS_RUN_SERVICES` was introduced as successor to `OCIS_RUN_EXTENSIONS` because we
   wanted to call oCIS "core" extensions services. We kept `OCIS_RUN_EXTENSIONS` for backwards
   compatibility reasons.

   It turned out, that setting `OCIS_RUN_SERVICES` has no effect since introduced.
   `OCIS_RUN_EXTENSIONS`. `OCIS_RUN_EXTENSIONS` was working fine all the time.

   We now fixed `OCIS_RUN_SERVICES`, so that you can use it as a equivalent replacement for
   `OCIS_RUN_EXTENSIONS`

   https://github.com/owncloud/ocis/pull/4133

* Bugfix - Fix the OIDC provider cache: [#4600](https://github.com/owncloud/ocis/pull/4600)

   We've fixed the OIDC provider cache. It never had a cache hit before this fix. Under some
   circumstances it could cause a painfully slow OCIS if the IDP well-known endpoint takes some
   time to respond.

   https://github.com/owncloud/ocis/pull/4600

* Bugfix - Fix permissions in REPORT: [#4520](https://github.com/owncloud/ocis/pull/4520)

   The REPORT endpoint wouldn't return any permissions on personal spaces Now it does. Also bumps
   reva

   https://github.com/owncloud/ocis/pull/4520

* Bugfix - Set default name for public link via capabilities: [#3834](https://github.com/owncloud/ocis/pull/3834)

   We have now added a default name for public link shares which is communicated via the
   capabilities.

   https://github.com/owncloud/ocis/issues/1237
   https://github.com/owncloud/ocis/pull/3834

* Bugfix - Remove legacy accounts proxy routes: [#3831](https://github.com/owncloud/ocis/pull/3831)

   We've removed the legacy accounts routes from the proxy default config. There were no longer
   used since the switch to IDM as the default user backend. Also accounts is no longer part of the
   oCIS binary and therefore should not be part of the proxy default route config.

   https://github.com/owncloud/ocis/pull/3831

* Bugfix - Fix unused config option `GRAPH_SPACES_INSECURE`: [#55555](https://github.com/owncloud/ocis/pull/55555)

   We've removed the unused config option `GRAPH_SPACES_INSECURE` from the GRAPH service.

   https://github.com/owncloud/ocis/pull/55555

* Bugfix - Remove unused configuration options: [#3973](https://github.com/owncloud/ocis/pull/3973)

   We've removed multiple unused configuration options:

   - `STORAGE_SYSTEM_DATAPROVIDER_INSECURE`, see also cs3org/reva#2993 -
   `STORAGE_USERS_DATAPROVIDER_INSECURE`, see also cs3org/reva#2993 -
   `STORAGE_SYSTEM_TEMP_FOLDER`, see also cs3org/reva#2993 -
   `STORAGE_USERS_TEMP_FOLDER`, see also cs3org/reva#2993 - `WEB_UI_CONFIG_VERSION`, see
   also owncloud/web#7130 - `GATEWAY_COMMIT_SHARE_TO_STORAGE_REF`, see also
   cs3org/reva#3017

   https://github.com/owncloud/ocis/pull/3973

* Bugfix - Remove static ocs user backend config: [#4077](https://github.com/owncloud/ocis/pull/4077)

   We've remove the `OCS_ACCOUNT_BACKEND_TYPE` configuration option. It was intended to allow
   configuration of different user backends for the ocs service. Right now the ocs service only
   has a "cs3" backend. Therefor it's a static entry and not configurable.

   https://github.com/owncloud/ocis/pull/4077

* Bugfix - Remove unused OCS storage configuration: [#3955](https://github.com/owncloud/ocis/pull/3955)

   We've removed the unused OCS configuration option `OCS_STORAGE_USERS_DRIVER`.

   https://github.com/owncloud/ocis/pull/3955

* Bugfix - Fix the `ocis search` command: [#3796](https://github.com/owncloud/ocis/pull/3796)

   We've fixed the behavior for `ocis search`, which didn't show further help when not all secrets
   have been configured. It also was not possible to start the search service standalone from the
   oCIS binary without configuring all oCIS secrets, even they were not needed by the search
   service.

   https://github.com/owncloud/ocis/pull/3796

* Bugfix - Rename search env variable for the grpc server address: [#3800](https://github.com/owncloud/ocis/pull/3800)

   We've fixed the gprc server address configuration environment variable by renaming it from
   `ACCOUNTS_GRPC_ADDR` to `SEARCH_GRPC_ADDR`

   https://github.com/owncloud/ocis/pull/3800

* Bugfix - Fix search in received shares: [#4308](https://github.com/owncloud/ocis/issues/4308)

   We fixed a problem where items in received shares were not found.

   https://github.com/owncloud/ocis/issues/4308

* Bugfix - Fix search report: [#7557](https://github.com/owncloud/web/issues/7557)

   There were multiple issues with REPORT search responses from webdav. Also we want it to be
   consistent with PROPFIND responses. * the `remote.php` prefix was missing from the href
   (added even though not necessary) * the ids were formatted wrong, they should look different
   for shares and spaces. * the name of the resource was missing * the shareid was missing (for
   shares) * the prop `shareroot` (containing the name of the share root) was missing * the
   permissions prop was empty

   https://github.com/owncloud/web/issues/7557
   https://github.com/owncloud/ocis/pull/4484

* Bugfix - Render webdav permissions as string in search report: [#4575](https://github.com/owncloud/ocis/issues/4575)

   We now correctly render the `oc:permissions` of resources as a string.

   https://github.com/owncloud/ocis/issues/4575
   https://github.com/owncloud/ocis/pull/4579

* Bugfix - Fix make sensitive config values in the proxy's debug server: [#4086](https://github.com/owncloud/ocis/pull/4086)

   We've fixed a security issue of the proxy's debug server config report endpoint. Previously
   sensitive configuration values haven't been masked. We now mask these values.

   https://github.com/owncloud/ocis/pull/4086

* Bugfix - Fix the idm and settings extensions' admin user id configuration option: [#3799](https://github.com/owncloud/ocis/pull/3799)

   We've fixed the admin user id configuration of the settings and idm extensions. The have
   previously only been configurable via the oCIS shared configuration and therefore have been
   undocumented for the extensions. This config option is now part of both extensions'
   configuration and can now also be used when the extensions are compiled standalone.

   https://github.com/owncloud/ocis/pull/3799

* Bugfix - Mail notifications for group shares: [#4714](https://github.com/owncloud/ocis/pull/4714)

   We fixed multiple issues in the notifications service, which broke notification mails new
   shares with groups.

   https://github.com/owncloud/ocis/issues/4703
   https://github.com/owncloud/ocis/issues/4688
   https://github.com/owncloud/ocis/pull/4714

* Bugfix - Substring search for sharees: [#547](https://github.com/owncloud/ocis/issues/547)

   We fixed searching for sharees to be no longer case-sensitive. With this we introduced two new
   settings for the users and groups services: "group_substring_filter_type" for the group
   services and "user_substring_filter_type" for the users service. They allow to set the type
   of LDAP filter that is used for substring user searches. Possible values are: "initial",
   "final" and "any" to do either prefix, suffix or full substring searches. Both settings
   default to "initial".

   Also a new option "search_min_length" was added for the "frontend" service. It allows to
   configure the minimum number of characters to enter before a search for Sharees is started.
   This setting is e.g. evaluated by the web ui via the capabilities endpoint.

   https://github.com/owncloud/ocis/issues/547

* Bugfix - Fix configuration validation for extensions' server commands: [#3911](https://github.com/owncloud/ocis/pull/3911)

   We've fixed the configuration validation for the extensions' server commands. Before this
   fix error messages have occurred when trying to start individual services without certain
   oCIS fullstack configuration values.

   We now no longer do the common oCIS configuration validation for extensions' server commands
   and now rely only on the extensions' validation function.

   https://github.com/owncloud/ocis/pull/3911

* Bugfix - Fix startup error logging: [#4093](https://github.com/owncloud/ocis/pull/4093)

   We've fixed the startup error logging, so that users will the reason for a failed startup even on
   "error" log level. Previously they would only see it on "info" log level. Also in a lot of cases
   the reason for the failed shutdown was omitted.

   https://github.com/owncloud/ocis/pull/4093

* Bugfix - Disable cache for selected static web assets: [#4809](https://github.com/owncloud/ocis/pull/4809)

   We've disabled caching for some static web assets. Files like the web index.html,
   oidc-callback.html or similar contain paths to timestamped resources and should not be
   cached.

   https://github.com/owncloud/ocis/pull/4809

* Bugfix - Fix multiple storage-users env variables: [#3802](https://github.com/owncloud/ocis/pull/3802)

   We've fixed multiple environment variable configuration options for the storage-users
   extension:

  * `STORAGE_USERS_GRPC_ADDR` was used to configure both the address of the http and grpc server. This resulted in a failing startup of the storage-users extension if this config option is set, because the service tries to double-bind the configured port (one time for each of the http and grpc server). You can now configure the grpc server's address with the environment variable `STORAGE_USERS_GRPC_ADDR` and the http server's address with the environment variable `STORAGE_USERS_HTTP_ADDR`
  * `STORAGE_USERS_S3NG_USERS_PROVIDER_ENDPOINT` was used to configure the permissions service endpoint for the S3NG driver and was therefore renamed to `STORAGE_USERS_S3NG_PERMISSIONS_ENDPOINT`
  * It's now possible to configure the permissions service endpoint for all  storage drivers with the environment variable `STORAGE_USERS_PERMISSION_ENDPOINT`, which was previously only used by the S3NG driver.

   https://github.com/owncloud/ocis/pull/3802

* Bugfix - Thumbnails for `/dav/xxx?preview=1` requests: [#3567](https://github.com/owncloud/ocis/pull/3567)

   We've added the thumbnail rendering for `/dav/xxx?preview=1`,
   `/remote.php/webdav/{relative path}?preview=1` and `/webdav/{relative
   path}?preview=1` requests, which was previously not supported because of missing routes. It
   now returns the same thumbnails as for `/remote.php/dav/xxx?preview=1`.

   https://github.com/owncloud/ocis/pull/3567

* Bugfix - Fix unfindable entities from shares/publicshares: [#4651](https://github.com/owncloud/ocis/pull/4651)

   We fixed a problem where directories or empty files weren't findable because they were to the
   search index improperly when created through a share or publicshare.

   https://github.com/owncloud/ocis/issues/4489
   https://github.com/owncloud/ocis/pull/4651

* Bugfix - Fix unrestricted quota on the graphAPI: [#4363](https://github.com/owncloud/ocis/pull/4363)

   Unrestricted quota needs to show 0 on the API. It is not good for clients when the property is
   missing.

   https://github.com/owncloud/ocis/pull/4363

* Bugfix - Fix user autoprovisioning: [#3893](https://github.com/owncloud/ocis/issues/3893)

   We've fixed the autoprovsioning feature that was introduced in beta2. Due to a bug the role
   assignment of the privileged user that is used to create accounts wasn't propagated correctly
   to the `graph` service.

   https://github.com/owncloud/ocis/issues/3893

* Bugfix - Fix version info: [#3953](https://github.com/owncloud/ocis/pull/3953)

   We've fixed the version info that is displayed when you run:

   - `ocis version` - `ocis <extension name> version`

   Since #2918, these commands returned an empty version only.

   https://github.com/owncloud/ocis/pull/3953

* Bugfix - Fix version number in status page: [#3788](https://github.com/owncloud/ocis/issues/3788)

   We needed to undo the version number changes on the status page to keep compatibility for legacy
   clients. We added a new field `productversion` for the actual version of the product.

   https://github.com/owncloud/ocis/issues/3788
   https://github.com/owncloud/ocis/pull/3805

* Bugfix - Fix CORS in frontend service: [#4948](https://github.com/owncloud/ocis/pull/4948)

   We now pass CORS config to the frontend reva service middleware.

   https://github.com/owncloud/ocis/issues/1340
   https://github.com/owncloud/ocis/pull/4948

* Bugfix - Graph service now forwards trace context: [#4582](https://github.com/owncloud/ocis/pull/4582)

   https://github.com/owncloud/ocis/pull/4582

* Bugfix - Fix the webdav URL of drive roots: [#3706](https://github.com/owncloud/ocis/issues/3706)

   Fixed the webdav URL of drive roots in the graph API.

   https://github.com/owncloud/ocis/issues/3706
   https://github.com/owncloud/ocis/pull/3916

* Bugfix - Idp: Check if CA certificate if present: [#3623](https://github.com/owncloud/ocis/issues/3623)

   Upon first start with the default configuration the idm service creates a server certificate,
   that might not be finished before the idp service is starting. Add a check to idp similar to what
   the user, group, and auth-providers implement.

   https://github.com/owncloud/ocis/issues/3623

* Bugfix - Fix graph endpoint: [#3925](https://github.com/owncloud/ocis/issues/3925)

   We have added the memberOf slice to the /users endpoint and the member slice to the /group
   endpoint

   https://github.com/owncloud/ocis/issues/3925

* Bugfix - Initial role assignment with external IDM: [#5045](https://github.com/owncloud/ocis/issues/5045)

   We've the initial user role assignment when using an external LDAP server.

   https://github.com/owncloud/ocis/issues/5045

* Bugfix - Escape DN attribute value: [#4117](https://github.com/owncloud/ocis/pull/4117)

   Escaped the DN attribute value on creating users and groups.

   https://github.com/owncloud/ocis/pull/4117

* Bugfix - Make IDP only wait for certs when using LDAP: [#3965](https://github.com/owncloud/ocis/pull/3965)

   When configuring cs3 as the backend the IDP no longer waits for an LDAP certificate to appear.

   https://github.com/owncloud/ocis/pull/3965

* Bugfix - Make ocdav service behave properly: [#3957](https://github.com/owncloud/ocis/pull/3957)

   The ocdav service now properly passes the tracing config and shuts down when receiving a kill
   signal.

   https://github.com/owncloud/ocis/pull/3957

* Bugfix - Make storage users mount ids unique by default: [#5091](https://github.com/owncloud/ocis/pull/5091)

   The mount ID of the storage users provider needs to be unique by default. We made this value
   configurable and added it to ocis init to be sure that we have a random uuid v4. This is important
   for federated instances.

   > **Warning** >BREAKING Change: In order to make every ocis storage provider ID unique by
   default, we needed to use a random uuidv4 during ocis init. Existing installations need to set
   this value explicitly or ocis will terminate after the upgrade. > To upgrade from 2.0.0-rc.1 to
   2.0.0-rc.2, 2.0.0 or later you need to set `GATEWAY_STORAGE_USERS_MOUNT_ID` and
   `STORAGE_USERS_MOUNT_ID` to the same random uuidv4. > >You can also add >``` >storage_users:
   > mount_id: some-random-uuid >gateway: > storage_registry: > storage_users_mount_id:
   some-random-uuid >``` >to the ocis.yaml file which was created during initialisation >
   >Changing the ID of the storage-users provider will change all >- WebDAV Urls >- FileIDs >-
   SpaceIDs >- Bookmarks >- and will make all existing shares invalid. > >The Android, Web and iOS
   clients will continue to work without interruptions. The Desktop Client sync connections
   need to be deleted and recreated. >Sorry for the inconvenience 😅 > >WORKAROUND - Not
   Recommended: You can avoid this by setting
   >`GATEWAY_STORAGE_USERS_MOUNT_ID=1284d238-aa92-42ce-bdc4-0b0000009157` and
   >`STORAGE_USERS_MOUNT_ID=1284d238-aa92-42ce-bdc4-0b0000009157` >But this will cause
   problems later when two ocis instances want to federate.

   https://github.com/owncloud/ocis/pull/5091

* Bugfix - Return proper errors when ocs/cloud/users is using the cs3 backend: [#3483](https://github.com/owncloud/ocis/issues/3483)

   The ocs API was just exiting with a fatal error on any update request, when configured for the cs3
   backend. Now it returns a proper error.

   https://github.com/owncloud/ocis/issues/3483

* Bugfix - Polish search: [#4094](https://github.com/owncloud/ocis/pull/4094)

   We improved the feedback when providing invalid search queries and added support for limiting
   the number of results returned.

   https://github.com/owncloud/ocis/pull/4094

* Bugfix - Fix the shareroot path in REPORT responses: [#4859](https://github.com/owncloud/ocis/pull/4859)

   Fixed the shareroot path in REPORT responses. Before this change the attribute leaked part of
   the folder tree of the sharer.

   https://github.com/owncloud/ocis/issues/4796
   https://github.com/owncloud/ocis/pull/4859

* Bugfix - Remove the storage-users event configuration: [#4825](https://github.com/owncloud/ocis/pull/4825)

   We've removed the events configuration from the storage-users section because it is not
   needed.

   https://github.com/owncloud/ocis/pull/4825

* Bugfix - Trigger a rescan of spaces in the search index when items have changed: [#4777](https://github.com/owncloud/ocis/pull/4777)

   The search service now scans spaces when items have been changed. This fixes the problem that
   mtime and treesize propagation was not reflected in the search index properly.

   https://github.com/owncloud/ocis/issues/4410
   https://github.com/owncloud/ocis/pull/4777

* Bugfix - Save Katherine: [#3823](https://github.com/owncloud/ocis/issues/3823)

   SpaceManager user katherine was removed with the demo user switch. Now she comes back

   https://github.com/owncloud/ocis/issues/3823
   https://github.com/owncloud/ocis/pull/3824

* Bugfix - Fix permission check in settings service: [#4890](https://github.com/owncloud/ocis/pull/4890)

   Added a check of the stored roles as a fallback if no roles are contained in the context.

   https://github.com/owncloud/ocis/pull/4890

* Bugfix - Fix Thumbnails for IDs without a trailing path: [#3791](https://github.com/owncloud/ocis/pull/3791)

   The routes in the chi router were not matching thumbnail requests without a trailing path.

   https://github.com/owncloud/ocis/pull/3791

* Bugfix - Space Creators can hand over spaces: [#4244](https://github.com/owncloud/ocis/pull/4244)

   Set no owner on non personal spaces to be able to pass the space manager role to a new user.

   https://github.com/owncloud/ocis/pull/4244

* Bugfix - Make tokeninfo endpoint unprotected: [#4715](https://github.com/owncloud/ocis/pull/4715)

   Make the tokeninfo endpoint unprotected as it is supposed to be available to the public.

   https://github.com/owncloud/ocis/pull/4715

* Bugfix - Update reva to version 2.12.0: [#5092](https://github.com/owncloud/ocis/pull/5092)

   Changelog for reva 2.12.0 (2022-11-25)  2 ✘  14:57:56 
   =======================================

  *   Bugfix [cs3org/reva#3436](https://github.com/cs3org/reva/pull/3436): Allow updating to internal link
  *   Bugfix [cs3org/reva#3473](https://github.com/cs3org/reva/pull/3473): Decomposedfs fix revision download
  *   Bugfix [cs3org/reva#3482](https://github.com/cs3org/reva/pull/3482): Decomposedfs propagate sizediff
  *   Bugfix [cs3org/reva#3449](https://github.com/cs3org/reva/pull/3449): Don't leak space information on update drive
  *   Bugfix [cs3org/reva#3470](https://github.com/cs3org/reva/pull/3470): Add missing events for managing spaces
  *   Bugfix [cs3org/reva#3472](https://github.com/cs3org/reva/pull/3472): Fix an oCDAV error message
  *   Bugfix [cs3org/reva#3452](https://github.com/cs3org/reva/pull/3452): Fix access to spaces shared via public link
  *   Bugfix [cs3org/reva#3440](https://github.com/cs3org/reva/pull/3440): Set proper names and paths for space roots
  *   Bugfix [cs3org/reva#3437](https://github.com/cs3org/reva/pull/3437): Refactor delete error handling
  *   Bugfix [cs3org/reva#3432](https://github.com/cs3org/reva/pull/3432): Remove share jail fix
  *   Bugfix [cs3org/reva#3458](https://github.com/cs3org/reva/pull/3458): Set the Oc-Fileid header when copying items
  *   Enhancement [cs3org/reva#3441](https://github.com/cs3org/reva/pull/3441): Cover ocdav with more unit tests
  *   Enhancement [cs3org/reva#3493](https://github.com/cs3org/reva/pull/3493): Configurable filelock duration factor in decomposedfs
  *   Enhancement [cs3org/reva#3397](https://github.com/cs3org/reva/pull/3397): Reduce lock contention issues

   https://github.com/owncloud/ocis/pull/5092
   https://github.com/owncloud/ocis/pull/5131

* Bugfix - URL encode the webdav url in the graph API: [#3597](https://github.com/owncloud/ocis/pull/3597)

   Fixed the webdav URL in the drives responses. Without encoding the URL could be broken by files
   with spaces in the file name.

   https://github.com/owncloud/ocis/issues/3538
   https://github.com/owncloud/ocis/pull/3597

* Bugfix - Store user passwords hashed in idm: [#3778](https://github.com/owncloud/ocis/issues/3778)

   Support for hashing user passwords was added to libregraph/idm. The graph API will now set
   userpasswords using the LDAP Modify Extended Operation (RFC3062). In the default
   configuration passwords will be hashed using the argon2id algorithm.

   https://github.com/owncloud/ocis/issues/3778
   https://github.com/owncloud/ocis/pull/4053

* Bugfix - Fix wopi access to public shares: [#4631](https://github.com/owncloud/ocis/pull/4631)

   I've added a request check to the public share authenticator middleware to allow wopi to access
   public shares.

   https://github.com/owncloud/ocis/issues/4382
   https://github.com/owncloud/ocis/pull/4631

* Change - Update ocis packages and imports to V2: [#3678](https://github.com/owncloud/ocis/pull/3678)

   This needs to be done in preparation for the major version bump in ocis.

   https://github.com/owncloud/ocis/pull/3678

* Change - Build service frontends with pnpm instead of yarn: [#4878](https://github.com/owncloud/ocis/pull/4878)

   We changed the Node.js packager from Yarn to pnpm to make it more consistent with the main Web
   repo. pnpm offers better package isolation and prevents a whole class of errors. This is only
   relevant for developers.

   https://github.com/owncloud/ocis/pull/4878
   https://github.com/owncloud/web/pull/7835

* Change - Load configuration files just from one directory: [#3587](https://github.com/owncloud/ocis/pull/3587)

   We've changed the configuration file loading behavior and are now only loading configuration
   files from ONE single directory. This directory can be set on compile time or via an environment
   variable on startup (`OCIS_CONFIG_DIR`).

   We are using following configuration default paths:

   - Docker images: `/etc/ocis/` - Binary releases: `$HOME/.ocis/config/`

   https://github.com/owncloud/ocis/pull/3587

* Change - Reduce permissions on docker image predeclared volumes: [#3641](https://github.com/owncloud/ocis/pull/3641)

   We've lowered the permissions on the predeclared volumes of the oCIS docker image from 777 to
   750.

   This change doesn't affect you, unless you use the docker image with the non default uid/guid to
   start oCIS (default is 1000:1000).

   https://github.com/owncloud/ocis/pull/3641

* Change - Introduce `ocis init` and remove all default secrets: [#3551](https://github.com/owncloud/ocis/pull/3551)

   We've removed all default secrets and the hardcoded UUID of the user `admin`. This means you
   can't start oCIS any longer without setting these via environment variable or configuration
   file.

   In order to make this easy for you, we introduced a new command: `ocis init`. You can run this
   command before starting oCIS with `ocis server` and it will bootstrap you a configuration file
   for a secure oCIS instance.

   https://github.com/owncloud/ocis/issues/3524
   https://github.com/owncloud/ocis/pull/3551
   https://github.com/owncloud/ocis/pull/3743

* Change - Rename "uploads purge" command to "uploads clean": [#4403](https://github.com/owncloud/ocis/pull/4403)

   We've renamed the storage-users service's "uploads purge" command to "upload clean".

   https://github.com/owncloud/ocis/pull/4403

* Change - Enable private links by default: [#4599](https://github.com/owncloud/ocis/pull/4599/)

   Enable private links by default in the capabilities.

   https://github.com/owncloud/ocis/pull/4599/

* Change - The `glauth` and `accounts` services are removed: [#3685](https://github.com/owncloud/ocis/pull/3685)

   After switching the default configuration to libregraph/idm we could remove the glauth and
   accounts services from the source code (they were already disabled by default with the
   previous release)

   https://github.com/owncloud/ocis/pull/3685

* Change - Reduce drives in graph /me/drives API: [#3629](https://github.com/owncloud/ocis/pull/3629)

   Reduced the drives in the graph `/me/drives` API to only the drives the user has access to. The
   endpoint `/drives` will list all drives when the user has the permission.

   https://github.com/owncloud/ocis/pull/3629

* Change - Switched default configuration to use libregraph/idm: [#3331](https://github.com/owncloud/ocis/pull/3331)

   We switched the default configuration of oCIS to use the "idm" service (based on
   libregraph/idm) as the standard source for user and group information. The accounts and
   glauth services are no longer enabled by default and will be removed with an upcoming release.

   https://github.com/owncloud/ocis/pull/3331
   https://github.com/owncloud/ocis/pull/3633

* Change - Rename MetadataUserID: [#3671](https://github.com/owncloud/ocis/pull/3671)

   MetadataUserID is renamed to SystemUserID including yaml tags and env vars

   https://github.com/owncloud/ocis/pull/3671

* Change - Use new space ID util functions: [#3648](https://github.com/owncloud/ocis/pull/3648)

   Changed code to use the new space ID util functions so that everything works with the new spaces
   ID format.

   https://github.com/owncloud/ocis/pull/3648
   https://github.com/owncloud/ocis/pull/3669

* Change - Prevent access to disabled space: [#3779](https://github.com/owncloud/ocis/pull/3779)

   Previously managers where allowed to edit the space even when it is disabled This is no longer
   possible

   https://github.com/owncloud/ocis/pull/3779

* Change - Rename serviceUser to systemUser: [#3673](https://github.com/owncloud/ocis/pull/3673)

   We renamed serviceUser to systemUser in all configs and vars including yaml-tags and env vars

   https://github.com/owncloud/ocis/pull/3673

* Change - Use the spaceID on the cs3 resource: [#4748](https://github.com/owncloud/ocis/pull/4748)

   We cleaned up the CS3Api to use a proper attribute for the space id.

   https://github.com/owncloud/ocis/pull/4748

* Change - Split MachineAuth from SystemUser: [#3672](https://github.com/owncloud/ocis/pull/3672)

   We now have two different APIKeys: MachineAuth for the machine-auth service and SystemUser
   for the system user used e.g. by settings service

   https://github.com/owncloud/ocis/pull/3672

* Enhancement - Add capability for alias links: [#3983](https://github.com/owncloud/ocis/issues/3983)

   For better UX clients need a way to discover if alias links are supported by the server. We added a
   capability under "files_sharing/public/alias"

   https://github.com/owncloud/ocis/issues/3983
   https://github.com/owncloud/ocis/pull/3991

* Enhancement - Add curl to the oCIS OCI image: [#4751](https://github.com/owncloud/ocis/pull/4751)

   We've added curl to the oCIS OCI image published on Dockerhub. This can be used for eg.
   healthchecks with the services' health endpoint.

   https://github.com/owncloud/ocis/pull/4751

* Enhancement - Add deprecation annotation: [#3917](https://github.com/owncloud/ocis/issues/3917)

   We have added the ability to annotate variables in case of deprecations:

   Example:

   `services/nats/pkg/config/config.go`

   ``` Host string `yaml:"host" env:"NATS_HOST_ADDRESS,NATS_NATS_HOST" desc:"Bind
   address." deprecationVersion:"1.6.2" removalVersion:"1.7.5" deprecationInfo:"the
   name is ugly" deprecationReplacement:"NATS_HOST_ADDRESS"` ```

   https://github.com/owncloud/ocis/issues/3917
   https://github.com/owncloud/ocis/pull/5143

* Enhancement - Add drives field to users endpoint: [#4072](https://github.com/owncloud/ocis/pull/4072)

   We have added `$expand=drives` to the `/users/{id}/` endpoint using the user filter
   implemented in reva.

   https://github.com/owncloud/ocis/pull/4072
   https://github.com/cs3org/reva/pull/3046
   https://github.com/owncloud/ocis/pull/4323

* Enhancement - Add Email templating: [#4564](https://github.com/owncloud/ocis/pull/4564)

   We have added email templating to ocis. Which are send on the SpaceShared and ShareCreated
   event.

   https://github.com/owncloud/ocis/issues/4303
   https://github.com/owncloud/ocis/pull/4564
   https://github.com/cs3org/reva/pull/3252

* Enhancement - Add FRONTEND_ENABLE_RESHARING env variable: [#4023](https://github.com/owncloud/ocis/pull/4023)

   We introduced resharing which was enabled by default, this is now configurable and can be
   enabled by setting the env `FRONTEND_ENABLE_RESHARING` to `true`. By default resharing is
   now disabled.

   https://github.com/owncloud/ocis/pull/4023

* Enhancement - We added e-mail subject templating: [#4799](https://github.com/owncloud/ocis/pull/4799)

   We have added e-mail subject templating.

   https://github.com/owncloud/ocis/pull/4799

* Enhancement - Add number of total matches to the search result: [#4189](https://github.com/owncloud/ocis/issues/4189)

   The search service now returns the number of total matches alongside the results.

   https://github.com/owncloud/ocis/issues/4189

* Enhancement - Add tracing to search: [#5113](https://github.com/owncloud/ocis/pull/5113)

   We added tracing to search and its indexer

   https://github.com/owncloud/ocis/issues/5063
   https://github.com/owncloud/ocis/pull/5113

* Enhancement - Add webURL to space root: [#4588](https://github.com/owncloud/ocis/pull/4588)

   Add the web url to the space root on the graphAPI.

   https://github.com/owncloud/ocis/pull/4588

* Enhancement - Align service naming: [#3606](https://github.com/owncloud/ocis/pull/3606)

   We now reflect the configured service names when listing them in the ocis runtime

   https://github.com/owncloud/ocis/issues/3603
   https://github.com/owncloud/ocis/pull/3606

* Enhancement - Add acting user to the audit log: [#3753](https://github.com/owncloud/ocis/issues/3753)

   Added the acting user to the events in the audit log.

   https://github.com/owncloud/ocis/issues/3753
   https://github.com/owncloud/ocis/pull/3992

* Enhancement - Configurable max lock cycles: [#4965](https://github.com/owncloud/ocis/pull/4965)

   Adds config option for max lock cycles. Also bumps reva

   https://github.com/owncloud/ocis/pull/4965

* Enhancement - Allow to configuring the reva cache store: [#4627](https://github.com/owncloud/ocis/pull/4627)

   We have added the possibility to configure the cache store implementation for the users
   storage.

   https://github.com/owncloud/ocis/pull/4627

* Enhancement - Add audit events for created containers: [#3941](https://github.com/owncloud/ocis/pull/3941)

   Handle the event `ContainerCreated` in the audit service.

   https://github.com/owncloud/ocis/pull/3941

* Enhancement - Add support for REPORT requests to /dav/spaces URLs: [#4661](https://github.com/owncloud/ocis/pull/4661)

   We added support for /dav/spaces REPORT requests which allow for searching specific spaces.

   https://github.com/owncloud/ocis/issues/4034
   https://github.com/owncloud/ocis/pull/4661

* Enhancement - Don't setup demo role assignments on default: [#3661](https://github.com/owncloud/ocis/issues/3661)

   Added a configuration option to explicitly tell the settings service to generate the default
   role assignments.

   https://github.com/owncloud/ocis/issues/3661
   https://github.com/owncloud/ocis/pull/3956

* Enhancement - Introduce "delete-all-spaces" permission: [#4196](https://github.com/owncloud/ocis/issues/4196)

   This is assigned to the Admin role by default and allows to cleanup orphaned spaces (e.g. where
   the owner as been deleted)

   https://github.com/owncloud/ocis/issues/4196

* Enhancement - Deny access to resources: [#4903](https://github.com/owncloud/ocis/pull/4903)

   We added an experimental feature to deny access to a certain resource. This feature is disabled
   by default and considered as EXPERIMENTAL. You can enable it by setting
   FRONTEND_OCS_ENABLE_DENIALS to `true`. It announces an available deny access permission
   via WebDAV on each resource. By convention it is only possible to deny access on folders. The
   clients can check the presence of the feature by the capability `deny_access` in the
   `files_sharing` section.

   https://github.com/owncloud/ocis/pull/4903

* Enhancement - Improve validation of OIDC access tokens: [#3841](https://github.com/owncloud/ocis/issues/3841)

   Previously OIDC access tokes were only validated by requesting the userinfo from the IDP. It is
   now possible to enable additional verification if the IDP issues access tokens in JWT format.
   In that case the oCIS proxy service will now verify the signature of the token using the public
   keys provided by jwks_uri endpoint of the IDP. It will also verify if the issuer claim (iss)
   matches the expected values.

   The new validation is enabled by setting `PROXY_OIDC_ACCESS_TOKEN_VERIFY_METHOD` to
   "jwt". Which is also the default. Setting it to "none" will disable the feature.

   https://github.com/owncloud/ocis/issues/3841
   https://github.com/owncloud/ocis/pull/4227

* Enhancement - Add /app/open-with-web endpoint: [#4376](https://github.com/owncloud/ocis/pull/4376)

   We've added an /app/open-with-web endpoint to the app provider, so that clients that are no
   browser or have only limited browser access can also open apps with the help of a Web URL.

   https://github.com/owncloud/ocis/pull/4376
   https://github.com/cs3org/reva/pull/3143

* Enhancement - Add previewFileMimeTypes to web default config: [#4414](https://github.com/owncloud/ocis/pull/4414)

   We've added previewFileMimeTypes to the web default config, so web can determine which
   preview types are supported by the backend.

   https://github.com/owncloud/ocis/pull/4414

* Enhancement - Added language option to the app provider: [#4399](https://github.com/owncloud/ocis/pull/4399)

   We've added a language option to the app provider which will in the end be passed to the app a user
   opens so that the web ui is displayed in the users language.

   https://github.com/owncloud/ocis/issues/4367
   https://github.com/owncloud/ocis/pull/4399
   https://github.com/cs3org/reva/pull/3156

* Enhancement - Improve error log for "could not get user by claim" error: [#4227](https://github.com/owncloud/ocis/pull/4227)

   We've improved the error log for "could not get user by claim" error where previously only the
   "nil" error has been logged. Now we're logging the message from the transport.

   https://github.com/owncloud/ocis/pull/4227

* Enhancement - Improve login screen design: [#4500](https://github.com/owncloud/ocis/pull/4500)

   We've improved the design of the login screen to match with the general design used in Web.

   https://github.com/owncloud/web/issues/7552
   https://github.com/owncloud/ocis/pull/4500

* Enhancement - Add configuration options for mail authentication and encryption: [#4443](https://github.com/owncloud/ocis/pull/4443)

   We've added configuration options to configure the authentication and encryption for
   sending mails in the notifications service.

   Furthermore there is now a distinguished configuration option for the username to use for
   authentication against the mail server. This allows you to customize the sender address to
   your liking. For example sender addresses like `my oCIS instance <ocis@owncloud.test>` are
   now possible, too.

   https://github.com/owncloud/ocis/pull/4443

* Enhancement - Introduce service registry cache: [#3833](https://github.com/owncloud/ocis/pull/3833)

   We've improved the service registry / service discovery by setting up registry caching (TTL
   20s), so that not every requests has to do a lookup on the registry.

   https://github.com/owncloud/ocis/pull/3833

* Enhancement - Reintroduce user autoprovisioning in proxy: [#3860](https://github.com/owncloud/ocis/pull/3860)

   With the removal of the accounts service autoprovisioning of users upon first login was no
   longer possible. We added this feature back for the cs3 user backend in the proxy. Leveraging
   the libregraph users API for creating the users.

   https://github.com/owncloud/ocis/pull/3860

* Enhancement - Allow to configure applications in Web: [#4578](https://github.com/owncloud/ocis/pull/4578)

   We've added the possibility to configure applications in the Web configuration.

   https://github.com/owncloud/ocis/pull/4578

* Enhancement - Added command to reset administrator password: [#4084](https://github.com/owncloud/ocis/issues/4084)

   The new command `ocis idm resetpassword` allows to reset the administrator password when ocis
   is not running. So it is possible to recover setups where the admin password was lost.

   https://github.com/owncloud/ocis/issues/4084
   https://github.com/owncloud/ocis/pull/4365

* Enhancement - Disable the color logging in docker compose examples: [#871](https://github.com/owncloud/ocis/issues/871)

   Disabled the color logging in the example docker compose deployments. Although colored logs
   are helpful during the development process they may be undesired in other situations like
   production deployments, where the logs aren't consumed by humans directly but instead by a log
   aggregator.

   https://github.com/owncloud/ocis/issues/871
   https://github.com/owncloud/ocis/pull/3935

* Enhancement - Allow providing list of services NOT to start: [#4254](https://github.com/owncloud/ocis/pull/4254)

   Until now if one wanted to use a custom version of a service, one needed to provide
   `OCIS_RUN_SERVICES` which is a list of all services to start. Now one can provide
   `OCIS_EXCLUDE_RUN_SERVICES` which is a list of only services not to start

   https://github.com/owncloud/ocis/pull/4254

* Enhancement - Introduce insecure flag for smtp email notifications: [#4279](https://github.com/owncloud/ocis/pull/4279)

   We've introduced the `NOTIFICATIONS_SMTP_INSECURE` configuration option, that let's you
   skip certificate verification for smtp email servers.

   https://github.com/owncloud/ocis/pull/4279

* Enhancement - Optional events in graph service: [#55555](https://github.com/owncloud/ocis/pull/55555)

   We've changed the graph service so that you also can start it without any event bus. Therefore
   you need to set `GRAPH_EVENTS_ENDPOINT` to an empty string. The graph API will not emit any
   events in this case.

   https://github.com/owncloud/ocis/pull/55555

* Enhancement - Fix behavior for foobar (in present tense): [#4346](https://github.com/owncloud/ocis/pull/4346)

   We've added the configuration option `PROXY_OIDC_REWRITE_WELLKNOWN` to rewrite the
   `/.well-known/openid-configuration` endpoint. If active, it serves the
   `/.well-known/openid-configuration` response of the original IDP configured in
   `OCIS_OIDC_ISSUER` / `PROXY_OIDC_ISSUER`. This is needed so that the Desktop Client,
   Android Client and iOS Client can discover the OIDC identity provider.

   Previously this rewrite needed to be performed with an external proxy as NGINX or Traefik if an
   external IDP was used.

   https://github.com/owncloud/ocis/issues/2819
   https://github.com/owncloud/ocis/issues/3280
   https://github.com/owncloud/ocis/pull/4346

* Enhancement - Add the "hidden" state to the search index: [#5018](https://github.com/owncloud/ocis/pull/5018)

   We changed the search service to store the "hidden" state of entries in the search index. That
   will allow for filtering/searching hidden files in the future.

   https://github.com/owncloud/ocis/pull/5018

* Enhancement - Restrict admins from self-removal: [#3713](https://github.com/owncloud/ocis/issues/3713)

   Admin users are no longer allowed to remove their own account or to edit their own role
   assignments. By this restriction we try to prevent situation where no administrative users is
   available in the system anymore

   https://github.com/owncloud/ocis/issues/3713

* Enhancement - OCS get share now also handle received shares: [#4322](https://github.com/owncloud/ocis/issues/4322)

   Requesting a specific share can now also correctly map the path to the mountpoint if the
   requested share is a received share.

   https://github.com/owncloud/ocis/issues/4322
   https://github.com/owncloud/ocis/pull/4539

* Enhancement - Add config option to provide TLS certificate: [#3818](https://github.com/owncloud/ocis/issues/3818)

   Added a config option to the graph service to provide a TLS certificate to be used to verify the
   LDAP server certificate.

   https://github.com/owncloud/ocis/issues/3818
   https://github.com/owncloud/ocis/pull/3888

* Enhancement - Add descriptions for graph-explorer config: [#3759](https://github.com/owncloud/ocis/pull/3759)

   Added descriptions tags to the graph-explorer config tags so that they will be included in the
   documentation.

   https://github.com/owncloud/ocis/pull/3759

* Enhancement - Add /me/changePassword endpoint to GraphAPI: [#3063](https://github.com/owncloud/ocis/issues/3063)

   When using the builtin user management, allow users to update their own password via the
   graph/v1.0/me/changePassword endpoint.

   https://github.com/owncloud/ocis/issues/3063
   https://github.com/owncloud/ocis/pull/3705

* Enhancement - Allow to setup TLS for grpc services: [#4798](https://github.com/owncloud/ocis/pull/4798)

   We added config options to allow enabling TLS encryption for all reva and go-micro backed grpc
   services.

   https://github.com/owncloud/ocis/pull/4798
   https://github.com/owncloud/ocis/pull/4901

* Enhancement - Generate signing key and encryption secret: [#3909](https://github.com/owncloud/ocis/issues/3909)

   The idp service now automatically generates a signing key and encryption secret when they
   don't exist. This will enable service restarts without invalidating existing sessions.

   https://github.com/owncloud/ocis/issues/3909
   https://github.com/owncloud/ocis/pull/4022

* Enhancement - Update IdP UI: [#3493](https://github.com/owncloud/ocis/issues/3493)

   Updated our fork of the lico IdP UI. This also updated the used npm dependencies. The design
   didn't change.

   https://github.com/owncloud/ocis/issues/3493
   https://github.com/owncloud/ocis/pull/4074

* Enhancement - Logging improvements: [#4815](https://github.com/owncloud/ocis/pull/4815)

   We improved the logging of several http services. If possible and present, we now log the
   `X-Request-Id`.

   https://github.com/owncloud/ocis/pull/4815
   https://github.com/owncloud/ocis/pull/4974

* Enhancement - Wrap metadata storage with dedicated reva gateway: [#3602](https://github.com/owncloud/ocis/pull/3602)

   We wrapped the metadata storage in a minimal reva instance with a dedicated gateway, including
   static storage registry, static auth registry, in memory userprovider, machine
   authprovider and demo permissions service. This allows us to preconfigure the service user
   for the ocis settings service, share and public share providers.

   https://github.com/owncloud/ocis/pull/3602
   https://github.com/owncloud/ocis/pull/3647

* Enhancement - New migrate command for migrating shares and public shares: [#3987](https://github.com/owncloud/ocis/pull/3987)

   We added a new `migrate` subcommand which can be used to migrate shares and public shares
   between different share and publicshare managers.

   https://github.com/owncloud/ocis/pull/3987
   https://github.com/owncloud/ocis/pull/4019

* Enhancement - Default to tls 1.2: [#4969](https://github.com/owncloud/ocis/pull/4969)

   https://github.com/owncloud/ocis/pull/4969

* Enhancement - Add missing unprotected paths: [#4454](https://github.com/owncloud/ocis/pull/4454)

   Added missing unprotected paths for the text-editor, preview, pdf-viewer, draw-io and
   index.html to the authentication middleware.

   https://github.com/owncloud/ocis/pull/4454
   https://github.com/owncloud/ocis/pull/4458

* Enhancement - Secure the nats connection with TLS: [#4781](https://github.com/owncloud/ocis/pull/4781)

   Encrypted the connection to the event broker using TLS. Per default TLS is not enabled but can be
   enabled by setting either `OCIS_EVENTS_ENABLE_TLS=true` or the respective service
   configs:

   - `AUDIT_EVENTS_ENABLE_TLS=true` - `GRAPH_EVENTS_ENABLE_TLS=true` -
   `NATS_EVENTS_ENABLE_TLS=true` - `NOTIFICATIONS_EVENTS_ENABLE_TLS=true` -
   `SEARCH_EVENTS_ENABLE_TLS=true` - `SHARING_EVENTS_ENABLE_TLS=true` -
   `STORAGE_USERS_EVENTS_ENABLE_TLS=true`

   https://github.com/owncloud/ocis/pull/4781
   https://github.com/owncloud/ocis/pull/4800
   https://github.com/owncloud/ocis/pull/4867

* Enhancement - Product field in OCS version: [#2918](https://github.com/owncloud/ocis/pull/2918)

   We've added a new field to the OCS Version, which is supposed to announce the product name. The
   web ui as a client will make use of it to make the backend product and version available (e.g. for
   easier bug reports).

   https://github.com/owncloud/ocis/pull/2918

* Enhancement - Automatically orientate photos when generating thumbnails: [#4477](https://github.com/owncloud/ocis/issues/4477)

   The thumbnailer now makes use of the exif orientation information to automatically orientate
   pictures before generating thumbnails.

   https://github.com/owncloud/ocis/issues/4477
   https://github.com/owncloud/ocis/pull/4513

* Enhancement - Refactor extensions to services: [#3980](https://github.com/owncloud/ocis/pull/3980)

   We have decided to name all extensions, we maintain and provide with ocis, services from here on
   to avoid confusion between external extensions and code we provide and maintain.

   https://github.com/owncloud/ocis/pull/3980

* Enhancement - Refactor the proxy service: [#4401](https://github.com/owncloud/ocis/issues/4401)

   The routes of the proxy service now have a "unprotected" flag. This is used by the
   authentication middleware to determine if the request needs to be blocked when missing
   authentication or not.

   https://github.com/owncloud/ocis/issues/4401
   https://github.com/owncloud/ocis/issues/4497
   https://github.com/owncloud/ocis/pull/4461
   https://github.com/owncloud/ocis/pull/4498
   https://github.com/owncloud/ocis/pull/4514

* Enhancement - Remove windows from ci & release makefile: [#5026](https://github.com/owncloud/ocis/pull/5026)

   We have removed windows from the ci & release makefile

   https://github.com/owncloud/ocis/issues/5011
   https://github.com/owncloud/ocis/pull/5026

* Enhancement - Rename AUTH_BASIC_AUTH_PROVIDER envvar: [#4966](https://github.com/owncloud/ocis/pull/4966)

   Rename the `AUTH_BASIC_AUTH_PROVIDER` envvar to `AUTH_BASIC_AUTH_MANAGER`

   https://github.com/owncloud/ocis/pull/4966
   https://github.com/owncloud/ocis/pull/4981

* Enhancement - Report parent id: [#4757](https://github.com/owncloud/ocis/pull/4757)

   We now index and return the parent id of a resource in search REPORTs.

   https://github.com/owncloud/ocis/issues/4727
   https://github.com/owncloud/ocis/pull/4757

* Enhancement - Allow resharing: [#3904](https://github.com/owncloud/ocis/pull/3904)

   This will allow resharing files

   https://github.com/owncloud/ocis/pull/3904

* Enhancement - Rewrite of the request authentication middleware: [#4374](https://github.com/owncloud/ocis/pull/4374)

   There were some flaws in the authentication middleware which were resolved by this rewrite.
   This rewrite also introduced the need to manually mark certain paths as "unprotected" if
   requests to these paths must not be authenticated.

   https://github.com/owncloud/ocis/pull/4374

* Enhancement - Add initial version of the search extensions: [#3635](https://github.com/owncloud/ocis/pull/3635)

   It is now possible to search for files and directories by their name using the web UI. Therefor
   new search extension indexes files in a persistent local index.

   https://github.com/owncloud/ocis/pull/3635

* Enhancement - Prohibit users from setting or listing other user's values: [#4897](https://github.com/owncloud/ocis/pull/4897)

   Added checks that users can only set and list their own settings.

   https://github.com/owncloud/ocis/pull/4897

* Enhancement - Add capability for public link single file edit: [#6787](https://github.com/owncloud/web/pull/6787)

   It is now possible to share a single file by link with edit permissions. Therefore we need a
   public share capability to enable that feature in the clients. At the same time, we improved the
   WebDAV permissions for public links.

   https://github.com/owncloud/web/pull/6787
   https://github.com/owncloud/ocis/pull/3538

* Enhancement - Added `share_jail` and `projects` feature flags in spaces capability: [#3626](https://github.com/owncloud/ocis/pull/3626)

   We've added feature flags to the `spaces` capability to indicate to clients which features are
   supposed to be shown to users.

   https://github.com/owncloud/ocis/pull/3626

* Enhancement - Use storageID when requesting special items: [#4356](https://github.com/owncloud/ocis/pull/4356)

   We need to use the storageID when requesting the special items of a space to spare a registry
   lookup and improve the performance

   https://github.com/owncloud/ocis/pull/4356

* Enhancement - Add description tags to the thumbnails config structs: [#3752](https://github.com/owncloud/ocis/pull/3752)

   Added description tags to the config structs in the thumbnails service so they will be included
   in the config documentation.

  **Important** If you ran `ocis init` with the `v2.0.0-alpha*` version then you have to manually add the `transfer_secret` to the ocis.yaml.

   Just open the `ocis.yaml` config file and look for the thumbnails section. Then add a random
   `transfer_secret` so that it looks like this:

   ```yaml thumbnails: thumbnail: transfer_secret: <put random value here> ```

   https://github.com/owncloud/ocis/pull/3752

* Enhancement - Make thumbnails service log less noisy: [#3959](https://github.com/owncloud/ocis/pull/3959)

   Reduced the log severity when no thumbnail was found from warn to debug. This reduces the spam in
   the logs.

   https://github.com/owncloud/ocis/pull/3959

* Enhancement - Add thumbnails support for tiff and bmp files: [#4634](https://github.com/owncloud/ocis/pull/4634)

   Support generating thumbnails for tiff and bmp files in the thumbnails service.

   https://github.com/owncloud/ocis/pull/4634

* Enhancement - Update linkshare capabilities: [#3579](https://github.com/owncloud/ocis/pull/3579)

   We have updated the capabilities regarding password enforcement and expiration dates of
   public links. They were previously hardcoded in a way that didn't reflect the actual backend
   functionality anymore.

   https://github.com/owncloud/ocis/pull/3579

* Enhancement - Update reva: [#3944](https://github.com/owncloud/ocis/pull/3944)

   Changelog for reva 2.6.1 (2022-06-27) =======================================

   The following sections list the changes in reva 2.6.1 relevant to reva users. The changes are
   ordered by importance.

   Summary -------

  * Bugfix [cs3org/reva#2998](https://github.com/cs3org/reva/pull/2998): Fix 0-byte-uploads
  * Enhancement [cs3org/reva#3983](https://github.com/cs3org/reva/pull/3983): Add capability for alias links
  * Enhancement [cs3org/reva#3000](https://github.com/cs3org/reva/pull/3000): Make less stat requests
  * Enhancement [cs3org/reva#3003](https://github.com/cs3org/reva/pull/3003): Distinguish GRPC FAILED_PRECONDITION and ABORTED codes
  * Enhancement [cs3org/reva#3005](https://github.com/cs3org/reva/pull/3005): Remove unused HomeMapping variable

   Changelog for reva 2.6.0 (2022-06-21) =======================================

   The following sections list the changes in reva 2.6.0 relevant to reva users. The changes are
   ordered by importance.

  * Bugfix [cs3org/reva#2985](https://github.com/cs3org/reva/pull/2985): Make stat requests route based on storage providerid
  * Bugfix [cs3org/reva#2987](https://github.com/cs3org/reva/pull/2987): Let archiver handle all error codes
  * Bugfix [cs3org/reva#2994](https://github.com/cs3org/reva/pull/2994): Bugfix errors when loading shares
  * Bugfix [cs3org/reva#2996](https://github.com/cs3org/reva/pull/2996): Do not close share dump channels
  * Bugfix [cs3org/reva#2993](https://github.com/cs3org/reva/pull/2993): Remove unused configuration
  * Bugfix [cs3org/reva#2950](https://github.com/cs3org/reva/pull/2950): Bugfix sharing with space ref
  * Bugfix [cs3org/reva#2991](https://github.com/cs3org/reva/pull/2991): Make sharesstorageprovider get accepted share
  * Change [cs3org/reva#2877](https://github.com/cs3org/reva/pull/2877): Enable resharing
  * Change [cs3org/reva#2984](https://github.com/cs3org/reva/pull/2984): Update CS3Apis
  * Enhancement [cs3org/reva#3753](https://github.com/cs3org/reva/pull/3753): Add executant to the events
  * Enhancement [cs3org/reva#2820](https://github.com/cs3org/reva/pull/2820): Instrument GRPC and HTTP requests with OTel
  * Enhancement [cs3org/reva#2975](https://github.com/cs3org/reva/pull/2975): Leverage shares space storageid and type when listing shares
  * Enhancement [cs3org/reva#3882](https://github.com/cs3org/reva/pull/3882): Explicitly return on ocdav move requests with body
  * Enhancement [cs3org/reva#2932](https://github.com/cs3org/reva/pull/2932): Stat accepted shares mountpoints, configure existing share updates
  * Enhancement [cs3org/reva#2944](https://github.com/cs3org/reva/pull/2944): Improve owncloudsql connection management
  * Enhancement [cs3org/reva#2962](https://github.com/cs3org/reva/pull/2962): Per service TracerProvider
  * Enhancement [cs3org/reva#2911](https://github.com/cs3org/reva/pull/2911): Allow for dumping and loading shares
  * Enhancement [cs3org/reva#2938](https://github.com/cs3org/reva/pull/2938): Sharpen tooling

   https://github.com/owncloud/ocis/pull/3944
   https://github.com/owncloud/ocis/pull/3975
   https://github.com/owncloud/ocis/pull/3982
   https://github.com/owncloud/ocis/pull/4000
   https://github.com/owncloud/ocis/pull/4006

* Enhancement - Update reva to version 2.7.2: [#4115](https://github.com/owncloud/ocis/pull/4115)

   Changelog for reva 2.7.2 (2022-07-18) =======================================

  * Bugfix [cs3org/reva#3079](https://github.com/cs3org/reva/pull/3079): Allow empty permissions
  * Bugfix [cs3org/reva#3084](https://github.com/cs3org/reva/pull/3084): Spaces related permissions and providerID cleanup
  * Bugfix [cs3org/reva#3083](https://github.com/cs3org/reva/pull/3083): Add space id to ItemTrashed event

   Changelog for reva 2.7.1 (2022-07-15) =======================================

  * Bugfix [cs3org/reva#3080](https://github.com/cs3org/reva/pull/3080): Make dataproviders return more headers
  * Enhancement [cs3org/reva#3046](https://github.com/cs3org/reva/pull/3046): Add user filter

   Changelog for reva 2.7.0 (2022-07-15) =======================================

  * Bugfix [cs3org/reva#3075](https://github.com/cs3org/reva/pull/3075): Check permissions of the move operation destination
  * Bugfix [cs3org/reva#3036](https://github.com/cs3org/reva/pull/3036):
  * Bugfix revad with EOS docker image
  * Bugfix [cs3org/reva#3037](https://github.com/cs3org/reva/pull/3037): Add uid- and gidNumber to LDAP queries
  * Bugfix [cs3org/reva#4061](https://github.com/cs3org/reva/pull/4061): Forbid resharing with higher permissions
  * Bugfix [cs3org/reva#3017](https://github.com/cs3org/reva/pull/3017): Removed unused gateway config "commit_share_to_storage_ref"
  * Bugfix [cs3org/reva#3031](https://github.com/cs3org/reva/pull/3031): Return proper response code when detecting recursive copy/move operations
  * Bugfix [cs3org/reva#3071](https://github.com/cs3org/reva/pull/3071): Make CS3 sharing drivers parse legacy resource id
  * Bugfix [cs3org/reva#3035](https://github.com/cs3org/reva/pull/3035): Prevent cross space move
  * Bugfix [cs3org/reva#3074](https://github.com/cs3org/reva/pull/3074): Send storage provider and space id to wopi server
  * Bugfix [cs3org/reva#3022](https://github.com/cs3org/reva/pull/3022): Improve the sharing internals
  * Bugfix [cs3org/reva#2977](https://github.com/cs3org/reva/pull/2977): Test valid filename on spaces tus upload
  * Change [cs3org/reva#3006](https://github.com/cs3org/reva/pull/3006): Use spaceID on the cs3api
  * Enhancement [cs3org/reva#3043](https://github.com/cs3org/reva/pull/3043): Introduce LookupCtx for index interface
  * Enhancement [cs3org/reva#3009](https://github.com/cs3org/reva/pull/3009): Prevent recursive copy/move operations
  * Enhancement [cs3org/reva#2977](https://github.com/cs3org/reva/pull/2977): Skip space lookup on space propfind

   https://github.com/owncloud/ocis/pull/4115
   https://github.com/owncloud/ocis/pull/4201
   https://github.com/owncloud/ocis/pull/4203
   https://github.com/owncloud/ocis/pull/4025
   https://github.com/owncloud/ocis/pull/4211

* Enhancement - Update reva to v2.7.4: [#4294](https://github.com/owncloud/ocis/pull/4294)

   Updated reva to version 2.7.4 This update includes:

  *  Bugfix [cs3org/reva#3141](https://github.com/cs3org/reva/pull/3141): Check ListGrants permission when listing shares

   Updated reva to version 2.7.3 This update includes:

  *  Bugfix [cs3org/reva#3109](https://github.com/cs3org/reva/pull/3109): Bugfix missing check in MustCheckNodePermissions
  *  Bugfix [cs3org/reva#3086](https://github.com/cs3org/reva/pull/3086): Bugfix crash in ldap authprovider
  *  Bugfix [cs3org/reva#3094](https://github.com/cs3org/reva/pull/3094): Allow removing password from public links
  *  Bugfix [cs3org/reva#3096](https://github.com/cs3org/reva/pull/3096): Bugfix user filter
  *  Bugfix [cs3org/reva#3091](https://github.com/cs3org/reva/pull/3091): Project spaces need no real owner
  *  Bugfix [cs3org/reva#3088](https://github.com/cs3org/reva/pull/3088): Use correct sublogger
  *  Enhancement [cs3org/reva#3123](https://github.com/cs3org/reva/pull/3123): Allow stating links that have no permissions
  *  Enhancement [cs3org/reva#3087](https://github.com/cs3org/reva/pull/3087): Allow to set LDAP substring filter type
  *  Enhancement [cs3org/reva#3098](https://github.com/cs3org/reva/pull/3098): App provider http endpoint uses Form instead of Query
  *  Enhancement [cs3org/reva#3133](https://github.com/cs3org/reva/pull/3133): Admins can set quota on all spaces
  *  Enhancement [cs3org/reva#3117](https://github.com/cs3org/reva/pull/3117): Update go-ldap to v3.4.4
  *  Enhancement [cs3org/reva#3095](https://github.com/cs3org/reva/pull/3095): Upload expiration and cleanup

   Https://github.com/owncloud/ocis/pull/4272
   https://github.com/cs3org/reva/pull/3096 https://github.com/cs3org/reva/pull/4315

   https://github.com/owncloud/ocis/pull/4294
   https://github.com/owncloud/ocis/pull/4330
   https://github.com/owncloud/ocis/pull/4369

* Enhancement - Update reva to v2.8.0: [#4444](https://github.com/owncloud/ocis/pull/4444)

   Updated reva to version 2.8.0. This update includes:

  * Bugfix [cs3org/reva#3158](https://github.com/cs3org/reva/pull/3158): Add name to the propfind response
  * Bugfix [cs3org/reva#3157](https://github.com/cs3org/reva/pull/3157): Fix locking response codes
  * Bugfix [cs3org/reva#3152](https://github.com/cs3org/reva/pull/3152): Disable caching of not found stat responses
  * Bugfix [cs3org/reva#4251](https://github.com/cs3org/reva/pull/4251): Disable caching
  * Enhancement [cs3org/reva#3154](https://github.com/cs3org/reva/pull/3154): Dataproviders now return file metadata
  * Enhancement [cs3org/reva#3143](https://github.com/cs3org/reva/pull/3143): Add /app/open-with-web endpoint
  * Enhancement [cs3org/reva#3156](https://github.com/cs3org/reva/pull/3156): Added language option to the app provider
  * Enhancement [cs3org/reva#3148](https://github.com/cs3org/reva/pull/3148): Add new jsoncs3 share manager

   https://github.com/owncloud/ocis/pull/4444

* Enhancement - Update reva to version 2.4.1: [#3746](https://github.com/owncloud/ocis/pull/3746)

   Changelog for reva 2.4.1 (2022-05-24) =======================================

   The following sections list the changes in reva 2.4.1 relevant to reva users. The changes are
   ordered by importance.

   Summary -------

  * Bugfix [cs3org/reva#2891](https://github.com/cs3org/reva/pull/2891): Add missing http status code

   Changelog for reva 2.4.0 (2022-05-24) =======================================

   The following sections list the changes in reva 2.4.0 relevant to reva users. The changes are
   ordered by importance.

   Summary -------

  * Bugfix [cs3org/reva#2854](https://github.com/cs3org/reva/pull/2854): Handle non uuid space and nodeid in decomposedfs
  * Bugfix [cs3org/reva#2853](https://github.com/cs3org/reva/pull/2853): Filter CS3 share manager listing
  * Bugfix [cs3org/reva#2868](https://github.com/cs3org/reva/pull/2868): Actually remove blobs when purging
  * Bugfix [cs3org/reva#2882](https://github.com/cs3org/reva/pull/2882): Fix FileUploaded event being emitted too early
  * Bugfix [cs3org/reva#2848](https://github.com/cs3org/reva/pull/2848): Fix storage id in the references in the ItemTrashed events
  * Bugfix [cs3org/reva#2852](https://github.com/cs3org/reva/pull/2852): Fix rcbox dependency on reva 1.18
  * Bugfix [cs3org/reva#3505](https://github.com/cs3org/reva/pull/3505): Fix creating a new file with wopi
  * Bugfix [cs3org/reva#2885](https://github.com/cs3org/reva/pull/2885): Move stat out of usershareprovider
  * Bugfix [cs3org/reva#2883](https://github.com/cs3org/reva/pull/2883): Fix role consideration when updating a share
  * Bugfix [cs3org/reva#2864](https://github.com/cs3org/reva/pull/2864): Fix Grant Space IDs
  * Bugfix [cs3org/reva#2870](https://github.com/cs3org/reva/pull/2870): Update quota calculation
  * Bugfix [cs3org/reva#2876](https://github.com/cs3org/reva/pull/2876): Fix version number in status page
  * Bugfix [cs3org/reva#2829](https://github.com/cs3org/reva/pull/2829): Don't include versions in quota
  * Change [cs3org/reva#2856](https://github.com/cs3org/reva/pull/2856): Do not allow to edit disabled spaces
  * Enhancement [cs3org/reva#3741](https://github.com/cs3org/reva/pull/3741): Add download endpoint to ocdav versions API
  * Enhancement [cs3org/reva#2884](https://github.com/cs3org/reva/pull/2884): Show mounted shares in virtual share jail root
  * Enhancement [cs3org/reva#2792](https://github.com/cs3org/reva/pull/2792): Use storageproviderid for spaces routing

   https://github.com/owncloud/ocis/pull/3746
   https://github.com/owncloud/ocis/pull/3771
   https://github.com/owncloud/ocis/pull/3778
   https://github.com/owncloud/ocis/pull/3842
   https://github.com/owncloud/ocis/pull/3854
   https://github.com/owncloud/ocis/pull/3858
   https://github.com/owncloud/ocis/pull/3867

* Enhancement - Update reva to version 2.5.1: [#3932](https://github.com/owncloud/ocis/pull/3932)

   Changelog for reva 2.5.1 (2022-06-08) =======================================

   The following sections list the changes in reva 2.5.1 relevant to reva users. The changes are
   ordered by importance.

   Summary -------

  * Bugfix [cs3org/reva#2931](https://github.com/cs3org/reva/pull/2931): Allow listing share jail space
  * Bugfix [cs3org/reva#2918](https://github.com/cs3org/reva/pull/2918): Fix propfinds with depth 0

   Changelog for reva 2.5.0 (2022-06-07) =======================================

   The following sections list the changes in reva 2.5.0 relevant to reva users. The changes are
   ordered by importance.

   Summary -------

  * Bugfix [cs3org/reva#2909](https://github.com/cs3org/reva/pull/2909): The decomposedfs now checks the GetPath permission
  * Bugfix [cs3org/reva#2899](https://github.com/cs3org/reva/pull/2899): Empty meta requests should return body
  * Bugfix [cs3org/reva#2928](https://github.com/cs3org/reva/pull/2928): Fix mkcol response code
  * Bugfix [cs3org/reva#2907](https://github.com/cs3org/reva/pull/2907): Correct share jail child aggregation
  * Bugfix [cs3org/reva#2895](https://github.com/cs3org/reva/pull/2895): Fix unlimited quota in spaces
  * Bugfix [cs3org/reva#2905](https://github.com/cs3org/reva/pull/2905): Check user permissions before updating/removing public shares
  * Bugfix [cs3org/reva#2904](https://github.com/cs3org/reva/pull/2904): Share jail now works properly when accessed as a space
  * Bugfix [cs3org/reva#2903](https://github.com/cs3org/reva/pull/2903): User owncloudsql now uses the correct userid
  * Change [cs3org/reva#2920](https://github.com/cs3org/reva/pull/2920): Clean up the propfind code
  * Change [cs3org/reva#2913](https://github.com/cs3org/reva/pull/2913): Rename ocs parameter "space_ref"
  * Enhancement [cs3org/reva#2919](https://github.com/cs3org/reva/pull/2919): EOS Spaces implementation
  * Enhancement [cs3org/reva#2888](https://github.com/cs3org/reva/pull/2888): Introduce spaces field mask
  * Enhancement [cs3org/reva#2922](https://github.com/cs3org/reva/pull/2922): Refactor webdav error handling

   https://github.com/owncloud/ocis/pull/3932
   https://github.com/owncloud/ocis/pull/3928
   https://github.com/owncloud/ocis/pull/3922

* Enhancement - Update Reva to version 2.10.0: [#4522](https://github.com/owncloud/ocis/pull/4522)

   Changelog for reva 2.10.0 (2022-09-09) =======================================

  * Bugfix [cs3org/reva#3210](https://github.com/cs3org/reva/pull/3210): Jsoncs3 mtime fix
  * Enhancement [cs3org/reva#3213](https://github.com/cs3org/reva/pull/3213): Allow for dumping the public shares from the cs3 publicshare manager
  * Enhancement [cs3org/reva#3199](https://github.com/cs3org/reva/pull/3199): Add support for cs3 storage backends to the json publicshare manager

   Changelog for reva 2.9.0 (2022-09-08) =======================================

  * Bugfix [cs3org/reva#3206](https://github.com/cs3org/reva/pull/3206): Add spaceid when listing share jail mount points
  * Bugfix [cs3org/reva#3194](https://github.com/cs3org/reva/pull/3194): Adds the rootinfo to storage spaces
  * Bugfix [cs3org/reva#3201](https://github.com/cs3org/reva/pull/3201): Fix shareid on PROPFIND
  * Bugfix [cs3org/reva#3176](https://github.com/cs3org/reva/pull/3176): Forbid duplicate shares
  * Bugfix [cs3org/reva#3208](https://github.com/cs3org/reva/pull/3208): Prevent panic in time conversion
  * Bugfix [cs3org/reva#3207](https://github.com/cs3org/reva/pull/3207): Align ocs status code for permission error on publiclink update
  * Enhancement [cs3org/reva#3193](https://github.com/cs3org/reva/pull/3193): Add shareid to PROPFIND
  * Enhancement [cs3org/reva#3180](https://github.com/cs3org/reva/pull/3180): Add canDeleteAllHomeSpaces permission
  * Enhancement [cs3org/reva#3203](https://github.com/cs3org/reva/pull/3203): Added "delete-all-spaces" permission
  * Enhancement [cs3org/reva#3200](https://github.com/cs3org/reva/pull/3200): OCS get share now also handle received shares
  * Enhancement [cs3org/reva#3185](https://github.com/cs3org/reva/pull/3185): Improve ldap authprovider's error reporting
  * Enhancement [cs3org/reva#3179](https://github.com/cs3org/reva/pull/3179): Improve tokeninfo endpoint
  * Enhancement [cs3org/reva#3171](https://github.com/cs3org/reva/pull/3171): Cs3 to jsoncs3 share manager migration
  * Enhancement [cs3org/reva#3204](https://github.com/cs3org/reva/pull/3204): Make the function flockFile private
  * Enhancement [cs3org/reva#3192](https://github.com/cs3org/reva/pull/3192): Enable space members to update shares

   https://github.com/owncloud/ocis/pull/4522
   https://github.com/owncloud/ocis/pull/4534
   https://github.com/owncloud/ocis/pull/4548
   https://github.com/owncloud/ocis/pull/4558

* Enhancement - Update reva to version 2.11.0: [#4588](https://github.com/owncloud/ocis/pull/4588)

   Changelog for reva 2.11.0 (2022-11-03) =======================================

  *   Bugfix  [cs3org/reva#3282](https://github.com/cs3org/reva/pull/3282):  Use Displayname in wopi apps
  *   Bugfix  [cs3org/reva#3430](https://github.com/cs3org/reva/pull/3430):  Add missing error check in decomposedfs
  *   Bugfix  [cs3org/reva#3298](https://github.com/cs3org/reva/pull/3298):  Make date only expiry dates valid for the whole day
  *   Bugfix  [cs3org/reva#3394](https://github.com/cs3org/reva/pull/3394):  Avoid AppProvider panic
  *   Bugfix  [cs3org/reva#3267](https://github.com/cs3org/reva/pull/3267):  Reduced default cache sizes for smaller memory footprint
  *   Bugfix  [cs3org/reva#3338](https://github.com/cs3org/reva/pull/3338):  Fix malformed uid string in cache
  *   Bugfix  [cs3org/reva#3255](https://github.com/cs3org/reva/pull/3255):  Properly escape oc:name in propfind response
  *   Bugfix  [cs3org/reva#3324](https://github.com/cs3org/reva/pull/3324):  Correct base URL for download URL and href when listing file public links
  *   Bugfix  [cs3org/reva#3278](https://github.com/cs3org/reva/pull/3278):  Fix public share view mode during app open
  *   Bugfix  [cs3org/reva#3377](https://github.com/cs3org/reva/pull/3377):  Fix possible race conditions
  *   Bugfix  [cs3org/reva#3274](https://github.com/cs3org/reva/pull/3274):  Fix "uploader" role permissions
  *   Bugfix  [cs3org/reva#3241](https://github.com/cs3org/reva/pull/3241):  Fix uploading empty files into shares
  *   Bugfix  [cs3org/reva#3251](https://github.com/cs3org/reva/pull/3251):  Make listing xattrs more robust
  *   Bugfix  [cs3org/reva#3287](https://github.com/cs3org/reva/pull/3287):  Return OCS forbidden error when a share already exists
  *   Bugfix  [cs3org/reva#3218](https://github.com/cs3org/reva/pull/3218):  Improve performance when listing received shares
  *   Bugfix  [cs3org/reva#3251](https://github.com/cs3org/reva/pull/3251):  Lock source on move
  *   Bugfix  [cs3org/reva#3238](https://github.com/cs3org/reva/pull/3238):  Return relative used quota amount as a percent value
  *   Bugfix  [cs3org/reva#3279](https://github.com/cs3org/reva/pull/3279):  Polish OCS error responses
  *   Bugfix  [cs3org/reva#3307](https://github.com/cs3org/reva/pull/3307):  Refresh lock in decomposedFS needs to overwrite
  *   Bugfix  [cs3org/reva#3368](https://github.com/cs3org/reva/pull/3368):  Return 404 when no permission to space
  *   Bugfix  [cs3org/reva#3341](https://github.com/cs3org/reva/pull/3341):  Validate s3ng downloads
  *   Bugfix  [cs3org/reva#3284](https://github.com/cs3org/reva/pull/3284):  Prevent nil pointer when requesting user
  *   Bugfix  [cs3org/reva#3257](https://github.com/cs3org/reva/pull/3257):  Fix wopi access to publicly shared files
  *   Change  [cs3org/reva#3267](https://github.com/cs3org/reva/pull/3267):  Decomposedfs no longer stores the idp
  *   Change  [cs3org/reva#3381](https://github.com/cs3org/reva/pull/3381):  Changed Name of the Shares Jail
  *   Enhancement  [cs3org/reva#3381](https://github.com/cs3org/reva/pull/3381):  Add capability for sharing by role
  *   Enhancement  [cs3org/reva#3320](https://github.com/cs3org/reva/pull/3320):  Add the parentID to the ocs and dav responses
  *   Enhancement  [cs3org/reva#3239](https://github.com/cs3org/reva/pull/3239):  Add privatelink to PROPFIND response
  *   Enhancement  [cs3org/reva#3340](https://github.com/cs3org/reva/pull/3340):  Add SpaceOwner to some event
  *   Enhancement  [cs3org/reva#3252](https://github.com/cs3org/reva/pull/3252):  Add SpaceShared event
  *   Enhancement  [cs3org/reva#3297](https://github.com/cs3org/reva/pull/3297):  Update dependencies
  *   Enhancement  [cs3org/reva#3429](https://github.com/cs3org/reva/pull/3429):  Make max lock cycles configurable
  *   Enhancement  [cs3org/reva#3011](https://github.com/cs3org/reva/pull/3011):  Expose capability to deny access in OCS API
  *   Enhancement  [cs3org/reva#3224](https://github.com/cs3org/reva/pull/3224):  Make the jsoncs3 share manager cache ttl configurable
  *   Enhancement  [cs3org/reva#3290](https://github.com/cs3org/reva/pull/3290):  Harden file system accesses
  *   Enhancement  [cs3org/reva#3332](https://github.com/cs3org/reva/pull/3332):  Allow to enable TLS for grpc service
  *   Enhancement  [cs3org/reva#3223](https://github.com/cs3org/reva/pull/3223):  Improve CreateShare grpc error reporting
  *   Enhancement  [cs3org/reva#3376](https://github.com/cs3org/reva/pull/3376):  Improve logging
  *   Enhancement  [cs3org/reva#3250](https://github.com/cs3org/reva/pull/3250):  Allow sharing the gateway caches
  *   Enhancement  [cs3org/reva#3240](https://github.com/cs3org/reva/pull/3240):  We now only encode &, < and > in PROPFIND PCDATA
  *   Enhancement  [cs3org/reva#3334](https://github.com/cs3org/reva/pull/3334):  Secure the nats connection with TLS
  *   Enhancement  [cs3org/reva#3300](https://github.com/cs3org/reva/pull/3300):  Do not leak existence of resources
  *   Enhancement  [cs3org/reva#3233](https://github.com/cs3org/reva/pull/3233):  Allow to override default broker for go-micro base ocdav service
  *   Enhancement  [cs3org/reva#3258](https://github.com/cs3org/reva/pull/3258):  Allow ocdav to share the registry instance with other services
  *   Enhancement  [cs3org/reva#3225](https://github.com/cs3org/reva/pull/3225):  Render file parent id for ocs shares
  *   Enhancement  [cs3org/reva#3222](https://github.com/cs3org/reva/pull/3222):  Support Prefer: return=minimal in PROPFIND
  *   Enhancement  [cs3org/reva#3395](https://github.com/cs3org/reva/pull/3395):  Reduce lock contention issues
  *   Enhancement  [cs3org/reva#3286](https://github.com/cs3org/reva/pull/3286):  Make Refresh Lock operation WOPI compliant
  *   Enhancement  [cs3org/reva#3229](https://github.com/cs3org/reva/pull/3229):  Request counting middleware
  *   Enhancement  [cs3org/reva#3312](https://github.com/cs3org/reva/pull/3312):  Implemented new share filters
  *   Enhancement  [cs3org/reva#3308](https://github.com/cs3org/reva/pull/3308):  Update the ttlcache library
  *   Enhancement  [cs3org/reva#3291](https://github.com/cs3org/reva/pull/3291):  The wopi app driver supports more options

   https://github.com/owncloud/ocis/pull/4588
   https://github.com/owncloud/ocis/pull/4716
   https://github.com/owncloud/ocis/pull/4719
   https://github.com/owncloud/ocis/pull/4750
   https://github.com/owncloud/ocis/pull/4833
   https://github.com/owncloud/ocis/pull/4867
   https://github.com/owncloud/ocis/pull/4903
   https://github.com/owncloud/ocis/pull/4908
   https://github.com/owncloud/ocis/pull/4915
   https://github.com/owncloud/ocis/pull/4964

* Enhancement - Update reva to v2.3.1: [#3552](https://github.com/owncloud/ocis/pull/3552)

   Updated reva to version 2.3.1. This update includes

  * Bugfix [cs3org/reva#2827](https://github.com/cs3org/reva/pull/2827): Check permissions when deleting spaces
  * Bugfix [cs3org/reva#2830](https://github.com/cs3org/reva/pull/2830): Correctly render response when accepting merged shares
  * Bugfix [cs3org/reva#2831](https://github.com/cs3org/reva/pull/2831): Fix uploads to owncloudsql storage when no mtime is provided
  * Enhancement [cs3org/reva#2833](https://github.com/cs3org/reva/pull/2833): Make status.php values configurable
  * Enhancement [cs3org/reva#2832](https://github.com/cs3org/reva/pull/2832): Add version option for ocdav go-micro service

   Updated reva to version 2.3.0. This update includes:

  * Bugfix [cs3org/reva#2693](https://github.com/cs3org/reva/pull/2693): Support editnew actions from MS Office
  * Bugfix [cs3org/reva#2588](https://github.com/cs3org/reva/pull/2588): Dockerfile.revad-ceph to use the right base image
  * Bugfix [cs3org/reva#2499](https://github.com/cs3org/reva/pull/2499): Removed check DenyGrant in resource permission
  * Bugfix [cs3org/reva#2285](https://github.com/cs3org/reva/pull/2285): Accept new userid idp format
  * Bugfix [cs3org/reva#2802](https://github.com/cs3org/reva/pull/2802): Bugfix the resource id handling for space shares
  * Bugfix [cs3org/reva#2800](https://github.com/cs3org/reva/pull/2800): Bugfix spaceid parsing in spaces trashbin API
  * Bugfix [cs3org/reva#2608](https://github.com/cs3org/reva/pull/2608): Respect the tracing_service_name config variable
  * Bugfix [cs3org/reva#2742](https://github.com/cs3org/reva/pull/2742): Use exact match in login filter
  * Bugfix [cs3org/reva#2759](https://github.com/cs3org/reva/pull/2759): Made uid, gid claims parsing more robust in OIDC auth provider
  * Bugfix [cs3org/reva#2788](https://github.com/cs3org/reva/pull/2788): Return the correct file IDs on public link resources
  * Bugfix [cs3org/reva#2322](https://github.com/cs3org/reva/pull/2322): Use RFC3339 for parsing dates
  * Bugfix [cs3org/reva#2784](https://github.com/cs3org/reva/pull/2784): Disable storageprovider cache for the share jail
  * Bugfix [cs3org/reva#2555](https://github.com/cs3org/reva/pull/2555): Bugfix site accounts endpoints
  * Bugfix [cs3org/reva#2675](https://github.com/cs3org/reva/pull/2675): Updates Makefile according to latest go standards
  * Bugfix [cs3org/reva#2572](https://github.com/cs3org/reva/pull/2572): Wait for nats server on middleware start
  * Change [cs3org/reva#2735](https://github.com/cs3org/reva/pull/2735): Avoid user enumeration
  * Change [cs3org/reva#2737](https://github.com/cs3org/reva/pull/2737): Bump go-cs3api
  * Change [cs3org/reva#2763](https://github.com/cs3org/reva/pull/2763): Change the oCIS and S3NG  storage driver blob store layout
  * Change [cs3org/reva#2596](https://github.com/cs3org/reva/pull/2596): Remove hash from public link urls
  * Change [cs3org/reva#2785](https://github.com/cs3org/reva/pull/2785): Implement workaround for chi.RegisterMethod
  * Change [cs3org/reva#2559](https://github.com/cs3org/reva/pull/2559): Do not encode webDAV ids to base64
  * Change [cs3org/reva#2740](https://github.com/cs3org/reva/pull/2740): Rename oc10 share manager driver
  * Change [cs3org/reva#2561](https://github.com/cs3org/reva/pull/2561): Merge oidcmapping auth manager into oidc
  * Enhancement [cs3org/reva#2698](https://github.com/cs3org/reva/pull/2698): Make capabilities endpoint public, authenticate users is present
  * Enhancement [cs3org/reva#2515](https://github.com/cs3org/reva/pull/2515): Enabling tracing by default if not explicitly disabled
  * Enhancement [cs3org/reva#2686](https://github.com/cs3org/reva/pull/2686): Features for favorites xattrs in EOS, cache for scope expansion
  * Enhancement [cs3org/reva#2494](https://github.com/cs3org/reva/pull/2494): Use sys ACLs for file permissions
  * Enhancement [cs3org/reva#2522](https://github.com/cs3org/reva/pull/2522): Introduce events
  * Enhancement [cs3org/reva#2811](https://github.com/cs3org/reva/pull/2811): Add event for created directories
  * Enhancement [cs3org/reva#2798](https://github.com/cs3org/reva/pull/2798): Add additional fields to events to enable search
  * Enhancement [cs3org/reva#2790](https://github.com/cs3org/reva/pull/2790): Fake providerids so API stays stable after beta
  * Enhancement [cs3org/reva#2685](https://github.com/cs3org/reva/pull/2685): Enable federated account access
  * Enhancement [cs3org/reva#1787](https://github.com/cs3org/reva/pull/1787): Add support for HTTP TPC
  * Enhancement [cs3org/reva#2799](https://github.com/cs3org/reva/pull/2799): Add flag to enable unrestricted listing of spaces
  * Enhancement [cs3org/reva#2560](https://github.com/cs3org/reva/pull/2560): Mentix PromSD extensions
  * Enhancement [cs3org/reva#2741](https://github.com/cs3org/reva/pull/2741): Meta path for user
  * Enhancement [cs3org/reva#2613](https://github.com/cs3org/reva/pull/2613): Externalize custom mime types configuration for storage providers
  * Enhancement [cs3org/reva#2163](https://github.com/cs3org/reva/pull/2163): Nextcloud-based share manager for pkg/ocm/share
  * Enhancement [cs3org/reva#2696](https://github.com/cs3org/reva/pull/2696): Preferences driver refactor and cbox sql implementation
  * Enhancement [cs3org/reva#2052](https://github.com/cs3org/reva/pull/2052): New CS3API datatx methods
  * Enhancement [cs3org/reva#2743](https://github.com/cs3org/reva/pull/2743): Add capability for public link single file edit
  * Enhancement [cs3org/reva#2738](https://github.com/cs3org/reva/pull/2738): Site accounts site-global settings
  * Enhancement [cs3org/reva#2672](https://github.com/cs3org/reva/pull/2672): Further Site Accounts improvements
  * Enhancement [cs3org/reva#2549](https://github.com/cs3org/reva/pull/2549): Site accounts improvements
  * Enhancement [cs3org/reva#2795](https://github.com/cs3org/reva/pull/2795): Add feature flags "projects" and "share_jail" to spaces capability
  * Enhancement [cs3org/reva#2514](https://github.com/cs3org/reva/pull/2514): Reuse ocs role objects in other drivers
  * Enhancement [cs3org/reva#2781](https://github.com/cs3org/reva/pull/2781): In memory user provider
  * Enhancement [cs3org/reva#2752](https://github.com/cs3org/reva/pull/2752): Refactor the rest user and group provider drivers

   https://github.com/owncloud/ocis/issues/3621
   https://github.com/owncloud/ocis/pull/3552
   https://github.com/owncloud/ocis/pull/3570
   https://github.com/owncloud/ocis/pull/3601
   https://github.com/owncloud/ocis/pull/3602
   https://github.com/owncloud/ocis/pull/3605
   https://github.com/owncloud/ocis/pull/3611
   https://github.com/owncloud/ocis/pull/3637
   https://github.com/owncloud/ocis/pull/3652
   https://github.com/owncloud/ocis/pull/3681

* Enhancement - Update ownCloud Web to v5.5.0-rc.8: [#6854](https://github.com/owncloud/web/pull/6854)

   Tags: web

   We updated ownCloud Web to v5.5.0-rc.8. Please refer to the changelog (linked) for details on
   the web release.

   https://github.com/owncloud/web/pull/6854
   https://github.com/owncloud/ocis/pull/3844
   https://github.com/owncloud/ocis/pull/3862
   https://github.com/owncloud/web/releases/tag/v5.5.0-rc.8

* Enhancement - Update ownCloud Web to v5.5.0-rc.9: [#6854](https://github.com/owncloud/web/pull/6854)

   Tags: web

   We updated ownCloud Web to v5.5.0-rc.9. Please refer to the changelog (linked) for details on
   the web release.

   Summary -------

  * Bugfix [owncloud/web#6939](https://github.com/owncloud/web/pull/6939): Not logged out if backend is ownCloud 10
  * Bugfix [owncloud/web#7061](https://github.com/owncloud/web/pull/7061): Prevent rename button from getting covered
  * Bugfix [owncloud/web#7032](https://github.com/owncloud/web/pull/7032): Show message when upload size exceeds quota
  * Bugfix [owncloud/web#7036](https://github.com/owncloud/web/pull/7036): Drag and drop upload when a file is selected
  * Enhancement [owncloud/web#7022](https://github.com/owncloud/web/pull/7022): Add config option for hoverable quick actions
  * Enhancement [owncloud/web#6555](https://github.com/owncloud/web/issues/6555): Consistent dropdown menus
  * Enhancement [owncloud/web#6994](https://github.com/owncloud/web/pull/6994): Copy/Move conflict dialog
  * Enhancement [owncloud/web#6750](https://github.com/owncloud/web/pull/6750): Make contexthelpers opt-out
  * Enhancement [owncloud/web#7038](https://github.com/owncloud/web/issues/7038): Rendering of share-indicators in ResourceTable
  * Enhancement [owncloud/web#6776](https://github.com/owncloud/web/issues/6776): Prevent the resource name in the sidebar from being truncated
  * Enhancement [owncloud/web#7067](https://github.com/owncloud/web/pull/7067): Upload progress & overlay improvements

   https://github.com/owncloud/web/pull/6854
   https://github.com/owncloud/ocis/pull/3927
   https://github.com/owncloud/web/releases/tag/v5.5.0-rc.9

* Enhancement - Update ownCloud Web to v5.5.0-rc.6: [#6854](https://github.com/owncloud/web/pull/6854)

   Tags: web

   We updated ownCloud Web to v5.5.0-rc.6. Please refer to the changelog (linked) for details on
   the web release.

   https://github.com/owncloud/web/pull/6854
   https://github.com/owncloud/ocis/pull/3664
   https://github.com/owncloud/ocis/pull/3680
   https://github.com/owncloud/ocis/pull/3727
   https://github.com/owncloud/ocis/pull/3747
   https://github.com/owncloud/ocis/pull/3797
   https://github.com/owncloud/web/releases/tag/v5.5.0-rc.6

* Enhancement - Update ownCloud Web to v5.7.0-rc.1: [#4005](https://github.com/owncloud/ocis/pull/4005)

   Tags: web

   We updated ownCloud Web to v5.7.0-rc.1. Please refer to the changelog (linked) for details on
   the web release.

  * Enhancement [owncloud/web#7119](https://github.com/owncloud/web/pull/7119): Copy/Move conflict dialog
  * Enhancement [owncloud/web#7122](https://github.com/owncloud/web/pull/7122): Enable Drag&Drop and keyboard shortcuts for all views
  * Enhancement [owncloud/web#7053](https://github.com/owncloud/web/pull/7053): Personal space id in URL
  * Enhancement [owncloud/web#6933](https://github.com/owncloud/web/pull/6933): Customize additional mimeTypes for preview app
  * Enhancement [owncloud/web#7078](https://github.com/owncloud/web/pull/7078): Add Hotkeys to ResourceTable
  * Enhancement [owncloud/web#7120](https://github.com/owncloud/web/pull/7120): Use tus chunksize from backend
  * Enhancement [owncloud/web#6749](https://github.com/owncloud/web/pull/6749): Update ODS to v13.2.0-rc.1
  * Enhancement [owncloud/web#7111](https://github.com/owncloud/web/pull/7111): Upload data during creation
  * Enhancement [owncloud/web#7109](https://github.com/owncloud/web/pull/7109): Clickable folder links in upload overlay
  * Enhancement [owncloud/web#7123](https://github.com/owncloud/web/pull/7123): Indeterminate progress bar in upload overlay
  * Enhancement [owncloud/web#7088](https://github.com/owncloud/web/pull/7088): Upload time estimation
  * Enhancement [owncloud/web#7125](https://github.com/owncloud/web/pull/7125): Wording improvements
  * Enhancement [owncloud/web#7140](https://github.com/owncloud/web/pull/7140): Separate direct and indirect link shares in sidebar
  * Bugfix [owncloud/web#7156](https://github.com/owncloud/web/pull/7156): Folder link targets
  * Bugfix [owncloud/web#7108](https://github.com/owncloud/web/pull/7108): Reload of an updated space-image and/or -readme
  * Bugfix [owncloud/web#6846](https://github.com/owncloud/web/pull/6846): Upload meta data serialization
  * Bugfix [owncloud/web#7100](https://github.com/owncloud/web/pull/7100): Complete-state of the upload overlay
  * Bugfix [owncloud/web#7104](https://github.com/owncloud/web/pull/7104): Parent folder name on public links
  * Bugfix [owncloud/web#7173](https://github.com/owncloud/web/pull/7173): Re-introduce dynamic app name in document title
  * Bugfix [owncloud/web#7166](https://github.com/owncloud/web/pull/7166): External apps fixes

   https://github.com/owncloud/ocis/pull/4005
   https://github.com/owncloud/web/pull/7158
   https://github.com/owncloud/ocis/pull/3990
   https://github.com/owncloud/web/pull/6854
   https://github.com/owncloud/web/releases/tag/v5.7.0-rc.1

* Enhancement - Update ownCloud Web to v6.0.0: [#5153](https://github.com/owncloud/ocis/pull/5153)

   Tags: web

   We updated ownCloud Web to v6.0.0. Please refer to the changelog (linked) for details on the web
   release.

   ### Breaking changes * BREAKING CHANGE for users in
   [owncloud/web#6648](https://github.com/owncloud/web/issues/6648): breaks existing
   bookmarks - they won't resolve anymore. * BREAKING CHANGE for developers in
   [owncloud/web#6648](https://github.com/owncloud/web/issues/6648): the appDefaults
   composables from web-pkg now work with drive aliases, concatenated with relative item paths,
   instead of webdav paths. If you use the appDefaults composables in your application it's
   likely that your code needs to be adapted.

   ### Changes * Bugfix
   [owncloud/web#7419](https://github.com/owncloud/web/issues/7419): Add language
   param opening external app * Bugfix
   [owncloud/web#7731](https://github.com/owncloud/web/pull/7731): "Copy
   Quicklink"-translations * Bugfix
   [owncloud/web#7830](https://github.com/owncloud/web/pull/7830): "Cut" and "Copy"
   actions for current folder * Bugfix
   [owncloud/web#7652](https://github.com/owncloud/web/pull/7652): Disable copy/move
   overwrite on self * Bugfix
   [owncloud/web#7739](https://github.com/owncloud/web/pull/7739): Disable shares
   loading on public and trash locations * Bugfix
   [owncloud/web#7740](https://github.com/owncloud/web/pull/7740): Disappearing
   quicklink in sidebar * Bugfix
   [owncloud/web#7946](https://github.com/owncloud/web/issues/7946): Prevent shares
   from disappearing after sharing with groups * Bugfix
   [owncloud/web#7820](https://github.com/owncloud/web/pull/7820): Edit new created
   user in user management * Bugfix
   [owncloud/web#7936](https://github.com/owncloud/web/pull/7936): Editing text files
   on public pages * Bugfix
   [owncloud/web#7861](https://github.com/owncloud/web/pull/7861): Handle non 2xx
   external app responses * Bugfix
   [owncloud/web#7734](https://github.com/owncloud/web/pull/7734): File name
   reactivity * Bugfix [owncloud/web#7975](https://github.com/owncloud/web/pull/7975):
   Prevent file upload when folder creation failed * Bugfix
   [owncloud/web#7724](https://github.com/owncloud/web/pull/7724): Folder conflict
   dialog * Bugfix [owncloud/web#7603](https://github.com/owncloud/web/issues/7603):
   Hide search bar in public link context * Bugfix
   [owncloud/web#7889](https://github.com/owncloud/web/pull/7889): Hide share
   indicators on public page * Bugfix
   [owncloud/web#7903](https://github.com/owncloud/web/issues/7903): "Keep
   both"-conflict option * Bugfix
   [owncloud/web#7697](https://github.com/owncloud/web/issues/7697): Link indicator on
   "Shared with me"-page * Bugfix
   [owncloud/web#8007](https://github.com/owncloud/web/pull/8007): Missing password
   form on public drop page * Bugfix
   [owncloud/web#7652](https://github.com/owncloud/web/pull/7652): Inhibit move files
   between spaces * Bugfix
   [owncloud/web#7985](https://github.com/owncloud/web/pull/7985): Prevent retrying
   uploads with status code 5xx * Bugfix
   [owncloud/web#7811](https://github.com/owncloud/web/pull/7811): Do not load files
   from cache in public links * Bugfix
   [owncloud/web#7941](https://github.com/owncloud/web/pull/7941): Add origin check to
   Draw.io events * Bugfix
   [owncloud/web#7916](https://github.com/owncloud/web/pull/7916): Prefer alias links
   over private links * Bugfix
   [owncloud/web#7640](https://github.com/owncloud/web/pull/7640): "Private
   link"-button alignment * Bugfix
   [owncloud/web#8006](https://github.com/owncloud/web/pull/8006): Public link loading
   on role change * Bugfix
   [owncloud/web#7962](https://github.com/owncloud/web/issues/7962): Quota check when
   replacing files * Bugfix
   [owncloud/web#7748](https://github.com/owncloud/web/pull/7748): Reload file list
   after last share removal * Bugfix
   [owncloud/web#7699](https://github.com/owncloud/web/issues/7699): Remove the "close
   sidebar"-calls on delete * Bugfix
   [owncloud/web#7504](https://github.com/owncloud/web/pull/7504): Resolve upload
   existing folder * Bugfix
   [owncloud/web#7771](https://github.com/owncloud/web/pull/7771): Routing for
   re-shares * Bugfix [owncloud/web#7675](https://github.com/owncloud/web/pull/7675):
   Search bar on small screens * Bugfix
   [owncloud/web#7662](https://github.com/owncloud/web/pull/7662): Sidebar for
   received shares in search file list * Bugfix
   [owncloud/web#7873](https://github.com/owncloud/web/pull/7873): Share editing after
   selecting a space * Bugfix
   [owncloud/web#7657](https://github.com/owncloud/web/issues/7657): Share
   permissions for re-shares * Bugfix
   [owncloud/web#7506](https://github.com/owncloud/web/issues/7506): Shares loading *
   Bugfix [owncloud/web#7632](https://github.com/owncloud/web/pull/7632): Sidebar
   toggle icon * Bugfix
   [owncloud/web#7781](https://github.com/owncloud/web/issues/7781): Sidebar without
   highlighted resource * Bugfix
   [owncloud/web#7756](https://github.com/owncloud/web/pull/7756): Try to obtain
   refresh token before the error case * Bugfix
   [owncloud/web#7768](https://github.com/owncloud/web/pull/7768): Hide actions in
   space trash bins * Bugfix
   [owncloud/web#7651](https://github.com/owncloud/web/pull/7651): Spaces on "Shared
   via link"-page * Bugfix
   [owncloud/web#7521](https://github.com/owncloud/web/issues/7521): Spaces
   reactivity on update * Bugfix
   [owncloud/web#7960](https://github.com/owncloud/web/issues/7960): Display error
   messages in text editor * Bugfix
   [owncloud/web#8030](https://github.com/owncloud/web/pull/8030): Saving a file
   multiple times with the text editor * Bugfix
   [owncloud/web#7778](https://github.com/owncloud/web/issues/7778): Trash bin sidebar
   * Bugfix [owncloud/web#7956](https://github.com/owncloud/web/issues/7956):
   Introduce "upload finalizing"-state in upload overlay * Bugfix
   [owncloud/web#7630](https://github.com/owncloud/web/pull/7630): Upload modify time *
   Bugfix [owncloud/web#8011](https://github.com/owncloud/web/issues/8011): Prevent
   unnecessary request when saving a user * Bugfix
   [owncloud/web#7989](https://github.com/owncloud/web/pull/7989): Versions on the
   "Shared with me"-page * Change
   [owncloud/web#6648](https://github.com/owncloud/web/issues/6648): Drive aliases in
   URLs * Change [owncloud/web#7935](https://github.com/owncloud/web/pull/7935): Remove
   mediaSource and v-image-source * Enhancement
   [owncloud/web#7635](https://github.com/owncloud/web/pull/7635): Add restore
   conflict dialog * Enhancement
   [owncloud/web#7901](https://github.com/owncloud/web/pull/7901): Add search field for
   space members * Enhancement
   [owncloud/web#4675](https://github.com/owncloud/web/issues/4675): Add
   `X-Request-ID` header to all outgoing requests * Enhancement
   [owncloud/web#7904](https://github.com/owncloud/web/pull/7904): Batch actions for
   two or more items only * Enhancement
   [owncloud/web#7892](https://github.com/owncloud/web/pull/7892): Respect the new
   sharing denials capability (experimental) * Enhancement
   [owncloud/web#7709](https://github.com/owncloud/web/pull/7709): Edit custom
   permissions wording * Enhancement
   [owncloud/web#7373](https://github.com/owncloud/web/issues/7373): Align dark mode
   colors with given design * Enhancement
   [owncloud/web#7190](https://github.com/owncloud/web/pull/7190): Deny subfolders
   inside share * Enhancement
   [owncloud/web#7684](https://github.com/owncloud/web/pull/7684): Design polishing *
   Enhancement [owncloud/web#7865](https://github.com/owncloud/web/pull/7865):
   Disable share renaming * Enhancement
   [owncloud/web#7725](https://github.com/owncloud/web/pull/7725): Enable renaming on
   received shares * Enhancement
   [owncloud/web#7747](https://github.com/owncloud/web/pull/7747): Friendlier logout
   screen * Enhancement
   [owncloud/web#6247](https://github.com/owncloud/web/issues/6247): Id based routing *
   Enhancement [owncloud/web#7803](https://github.com/owncloud/web/issues/7803):
   Internal link on unaccepted share * Enhancement
   [owncloud/web#7304](https://github.com/owncloud/web/issues/7304): Resolve internal
   links * Enhancement [owncloud/web#7569](https://github.com/owncloud/web/pull/7569):
   Make keybindings global * Enhancement
   [owncloud/web#7894](https://github.com/owncloud/web/pull/7894): Optimize email
   validation in the user management app * Enhancement
   [owncloud/web#7707](https://github.com/owncloud/web/issues/7707): Resolve private
   links * Enhancement
   [owncloud/web#7234](https://github.com/owncloud/web/issues/7234): Auth context in
   route meta props * Enhancement
   [owncloud/web#7821](https://github.com/owncloud/web/pull/7821): Improve search
   experience * Enhancement
   [owncloud/web#7801](https://github.com/owncloud/web/pull/7801): Make search results
   sortable * Enhancement
   [owncloud/web#8028](https://github.com/owncloud/web/pull/8028): Update ODS to
   v14.0.1 * Enhancement
   [owncloud/web#7890](https://github.com/owncloud/web/pull/7890): Validate space
   names * Enhancement [owncloud/web#7430](https://github.com/owncloud/web/pull/7430):
   Webdav support in web-client package * Enhancement
   [owncloud/web#7900](https://github.com/owncloud/web/issues/7900): XHR upload
   timeout

   https://github.com/owncloud/ocis/pull/5153
   https://github.com/owncloud/web/releases/tag/v6.0.0

* Enhancement - Update ownCloud Web to v5.7.0-rc.4: [#4140](https://github.com/owncloud/ocis/pull/4140)

   Tags: web

   We updated ownCloud Web to v5.7.0-rc.4. Please refer to the changelog (linked) for details on
   the web release.

  * Bugfix [owncloud/web#7230](https://github.com/owncloud/web/pull/7230): Context menu misplaced when triggered by keyboard navigation
  * Bugfix [owncloud/web#7214](https://github.com/owncloud/web/pull/7214): Prevent error when pasting with empty clipboard
  * Bugfix [owncloud/web#7173](https://github.com/owncloud/web/pull/7173): Re-introduce dynamic app name in document title
  * Bugfix [owncloud/web#7166](https://github.com/owncloud/web/pull/7166): External apps fixes
  * Bugfix [owncloud/web#7248](https://github.com/owncloud/web/pull/7248): Hide empty trash bin modal on error
  * Bugfix [owncloud/web#4677](https://github.com/owncloud/web/issues/4677): Logout deleted user on page reload
  * Bugfix [owncloud/web#7216](https://github.com/owncloud/web/pull/7216): Filename hovers over the image in the preview app
  * Bugfix [owncloud/web#7228](https://github.com/owncloud/web/pull/7228): Shared with others page apps not working with oc10 as backend
  * Bugfix [owncloud/web#7197](https://github.com/owncloud/web/pull/7197): Create space and access user management permission
  * Bugfix [owncloud/web#6921](https://github.com/owncloud/web/pull/6921): Space sidebar sharing indicators
  * Bugfix [owncloud/web#7030](https://github.com/owncloud/web/issues/7030): Access token renewal
  * Enhancement [owncloud/web#7217](https://github.com/owncloud/web/pull/7217): Add app top bar component
  * Enhancement [owncloud/web#7153](https://github.com/owncloud/web/pull/7153): Add Keyboard navigation/selection
  * Enhancement [owncloud/web#7030](https://github.com/owncloud/web/issues/7030): Loading context blocks application bootstrap
  * Enhancement [owncloud/web#7206](https://github.com/owncloud/web/pull/7206): Add change own password dialog to the account info page
  * Enhancement [owncloud/web#7086](https://github.com/owncloud/web/pull/7086): Re-sharing for ocis
  * Enhancement [owncloud/web#7201](https://github.com/owncloud/web/pull/7201): Added a toolbar to pdf-viewer app
  * Enhancement [owncloud/web#7139](https://github.com/owncloud/web/pull/7139): Reposition notifications
  * Enhancement [owncloud/web#7030](https://github.com/owncloud/web/issues/7030): Resolve bookmarked public links with password protection
  * Enhancement [owncloud/web#7038](https://github.com/owncloud/web/issues/7038): Improve performance of share indicators
  * Enhancement [owncloud/web#6661](https://github.com/owncloud/web/issues/6661): Option to block file extensions from text-editor app
  * Enhancement [owncloud/web#7139](https://github.com/owncloud/web/pull/7139): Update ODS to v14.0.0-alpha.4
  * Enhancement [owncloud/web#7176](https://github.com/owncloud/web/pull/7176): Introduce group assignments

   https://github.com/owncloud/ocis/pull/4140
   https://github.com/owncloud/web/releases/tag/v5.7.0-rc.4

* Enhancement - Update ownCloud Web to v5.7.0-rc.8: [#4314](https://github.com/owncloud/ocis/pull/4314)

   Tags: web

   We updated ownCloud Web to v5.7.0-rc.9. Please refer to the changelog (linked) for details on
   the web release.

  * Bugfix [owncloud/web#7080](https://github.com/owncloud/web/issues/7080): Add Droparea again
  * Bugfix [owncloud/web#7357](https://github.com/owncloud/web/pull/7357): Batch deleting multiple files
  * Bugfix [owncloud/web#7379](https://github.com/owncloud/web/pull/7379): Decline share not possible
  * Bugfix [owncloud/web#7322](https://github.com/owncloud/web/pull/7322): Files pagination scroll to top
  * Bugfix [owncloud/web#7348](https://github.com/owncloud/web/pull/7348): Left sidebar active navigation item has wrong cursor
  * Bugfix [owncloud/web#7355](https://github.com/owncloud/web/pull/7355): Link indicator on "Shared via link"-page
  * Bugfix [owncloud/web#7325](https://github.com/owncloud/web/pull/7325): Loading state in views
  * Bugfix [owncloud/web#7344](https://github.com/owncloud/web/pull/7344): Missing file icon in details panel
  * Bugfix [owncloud/web#7321](https://github.com/owncloud/web/pull/7321): Missing scroll bar in user management app
  * Bugfix [owncloud/web#7334](https://github.com/owncloud/web/pull/7334): No redirect after disabling space
  * Bugfix [owncloud/web#3071](https://github.com/owncloud/web/issues/3071): Don't leak oidc callback url into browser history
  * Bugfix [owncloud/web#7379](https://github.com/owncloud/web/pull/7379): Open file on shared space resource not possible
  * Bugfix [owncloud/web#7268](https://github.com/owncloud/web/issues/7268): Personal shares leaked into project space
  * Bugfix [owncloud/web#7359](https://github.com/owncloud/web/pull/7359): Fix infinite loading spinner on invalid preview links
  * Bugfix [owncloud/web#7272](https://github.com/owncloud/web/issues/7272): Print backend version
  * Bugfix [owncloud/web#7424](https://github.com/owncloud/web/pull/7424): Quicklinks not shown
  * Bugfix [owncloud/web#7379](https://github.com/owncloud/web/pull/7379): Rename shared space resource not possible
  * Bugfix [owncloud/web#7210](https://github.com/owncloud/web/pull/7210): Repair navigation highlighter
  * Bugfix [owncloud/web#7393](https://github.com/owncloud/web/pull/7393): Selected item bottom glue
  * Bugfix [owncloud/web#7308](https://github.com/owncloud/web/pull/7308): "Shared with others" and "Shared via Link" resource links not working
  * Bugfix [owncloud/web#7400](https://github.com/owncloud/web/issues/7400): Respect space quota permission
  * Bugfix [owncloud/web#7349](https://github.com/owncloud/web/pull/7349): Missing quick actions in spaces file list
  * Bugfix [owncloud/web#7396](https://github.com/owncloud/web/pull/7396): Add storage ID when navigating to a shared parent directory
  * Bugfix [owncloud/web#7394](https://github.com/owncloud/web/pull/7394): Suppress active panel error log
  * Bugfix [owncloud/web#7038](https://github.com/owncloud/web/issues/7038): File list render performance
  * Bugfix [owncloud/web#7240](https://github.com/owncloud/web/issues/7240): Access token renewal during upload
  * Bugfix [owncloud/web#7376](https://github.com/owncloud/web/pull/7376): Tooltips not shown on disabled create and upload button
  * Bugfix [owncloud/web#7297](https://github.com/owncloud/web/pull/7297): Upload overlay progress bar spacing
  * Bugfix [owncloud/web#7332](https://github.com/owncloud/web/pull/7332): Users list not loading if user has no role
  * Bugfix [owncloud/web#7313](https://github.com/owncloud/web/pull/7313): Versions of shared files not visible
  * Enhancement [owncloud/web#7404](https://github.com/owncloud/web/pull/7404): Adjust helper texts
  * Enhancement [owncloud/web#7350](https://github.com/owncloud/web/pull/7350): Change file loading mechanism in `preview` app
  * Enhancement [owncloud/web#7356](https://github.com/owncloud/web/pull/7356): Declined shares are now easily accessible
  * Enhancement [owncloud/web#7365](https://github.com/owncloud/web/pull/7365): Drop menu styling in right sidebar
  * Enhancement [owncloud/web#7252](https://github.com/owncloud/web/pull/7252): Redesign shared with list
  * Enhancement [owncloud/web#7371](https://github.com/owncloud/web/pull/7371): Use fixed width for the right sidebar
  * Enhancement [owncloud/web#7267](https://github.com/owncloud/web/pull/7267): Search all files announce limit
  * Enhancement [owncloud/web#7364](https://github.com/owncloud/web/pull/7364): Sharing panel show label instead of description for links
  * Enhancement [owncloud/web#7355](https://github.com/owncloud/web/pull/7355): Update ODS to v14.0.0-alpha.12
  * Enhancement [owncloud/web#7375](https://github.com/owncloud/web/pull/7375): User management app saved dialog

   https://github.com/owncloud/ocis/pull/4314
   https://github.com/owncloud/web/releases/tag/v5.7.0-rc.8

* Enhancement - Update ownCloud Web to v5.7.0-rc.10: [#4439](https://github.com/owncloud/ocis/pull/4439)

   Tags: web

   We updated ownCloud Web to v5.7.0-rc.10. Please refer to the changelog (linked) for details on
   the web release.

  * Bugfix [owncloud/web#7443](https://github.com/owncloud/web/pull/7443): Datetime formatting
  * Bugfix [owncloud/web#7437](https://github.com/owncloud/web/pull/7437): Default to user context
  * Bugfix [owncloud/web#7473](https://github.com/owncloud/web/pull/7473): Dragging a file causes no selection
  * Bugfix [owncloud/web#7469](https://github.com/owncloud/web/pull/7469): File size not updated while restoring file version
  * Bugfix [owncloud/web#7443](https://github.com/owncloud/web/pull/7443): File size formatting
  * Bugfix [owncloud/web#7474](https://github.com/owncloud/web/pull/7474): Load only supported thumbnails (configurable)
  * Bugfix [owncloud/web#7309](https://github.com/owncloud/web/pull/7309): SidebarNavItem icon flickering
  * Bugfix [owncloud/web#7425](https://github.com/owncloud/web/pull/7425): Open Folder in project space context menu
  * Bugfix [owncloud/web#7486](https://github.com/owncloud/web/issues/7486): Prevent unnecessary PROPFIND request during upload
  * Bugfix [owncloud/web#7415](https://github.com/owncloud/web/pull/7415): Re-fetch quota
  * Bugfix [owncloud/web#7478](https://github.com/owncloud/web/issues/7478): "Shared via"-indicator for links
  * Bugfix [owncloud/web#7480](https://github.com/owncloud/web/issues/7480): Missing space image in sidebar
  * Bugfix [owncloud/web#7436](https://github.com/owncloud/web/issues/7436): Hide share actions for space viewers/editors
  * Bugfix [owncloud/web#7445](https://github.com/owncloud/web/pull/7445): User management app close side bar throws error
  * Enhancement [owncloud/web#7309](https://github.com/owncloud/web/pull/7309): Keyboard shortcut indicators in ContextMenu
  * Enhancement [owncloud/web#7309](https://github.com/owncloud/web/pull/7309): Lowlight cut resources
  * Enhancement [owncloud/web#7133](https://github.com/owncloud/web/pull/7133): Permissionless (internal) link shares
  * Enhancement [owncloud/web#7309](https://github.com/owncloud/web/pull/7309): Replace locationpicker with clipboard actions
  * Enhancement [owncloud/web#7363](https://github.com/owncloud/web/pull/7363): Streamline UI sizings
  * Enhancement [owncloud/web#7355](https://github.com/owncloud/web/pull/7355): Update ODS to v14.0.0-alpha.16
  * Enhancement [owncloud/web#7476](https://github.com/owncloud/web/pull/7476): Users table on small screen
  * Enhancement [owncloud/web#7182](https://github.com/owncloud/web/pull/7182): User management app edit quota

   https://github.com/owncloud/ocis/pull/4439
   https://github.com/owncloud/web/releases/tag/v5.7.0-rc.10

* Enhancement - Update ownCloud Web to v5.7.0: [#4508](https://github.com/owncloud/ocis/pull/4508)

   Tags: web

   We updated ownCloud Web to v5.7.0. Please refer to the changelog (linked) for details on the web
   release.

  * Bugfix [owncloud/web#7522](https://github.com/owncloud/web/pull/7522): Allow uploads outside of user's home despite quota being exceeded
  * Bugfix [owncloud/web#7622](https://github.com/owncloud/web/issues/7622): Expiration date picker with long language codes
  * Bugfix [owncloud/web#7516](https://github.com/owncloud/web/pull/7516): File name in text editor
  * Bugfix [owncloud/web#7498](https://github.com/owncloud/web/issues/7498): Fix right sidebar content on small screens
  * Bugfix [owncloud/web#7455](https://github.com/owncloud/web/issues/7455): Improve keyboard shortcuts copy/cut files
  * Bugfix [owncloud/web#7510](https://github.com/owncloud/web/issues/7510): Paste action (keyboard) not working in project spaces
  * Bugfix [owncloud/web#7526](https://github.com/owncloud/web/issues/7526): Left sidebar when switching apps
  * Bugfix [owncloud/web#7582](https://github.com/owncloud/web/issues/7582): Merge share with group and group member into one
  * Bugfix [owncloud/web#7534](https://github.com/owncloud/web/issues/7534): Redirect after removing self from space members
  * Bugfix [owncloud/web#7560](https://github.com/owncloud/web/pull/7560): Search share representation
  * Bugfix [owncloud/web#7519](https://github.com/owncloud/web/issues/7519): Sidebar for current folder
  * Bugfix [owncloud/web#7453](https://github.com/owncloud/web/issues/7453): Stuck After Session Expired
  * Bugfix [owncloud/web#7595](https://github.com/owncloud/web/pull/7595): Typo when reading public links capabilities
  * Enhancement [owncloud/web#7570](https://github.com/owncloud/web/pull/7570): Adjust spacing of the files list options menu
  * Enhancement [owncloud/web#7540](https://github.com/owncloud/web/issues/7540): Left sidebar hover effect
  * Enhancement [owncloud/web#7555](https://github.com/owncloud/web/pull/7555): Propose unique file name while creating a new file
  * Enhancement [owncloud/web#7038](https://github.com/owncloud/web/issues/7038): Reduce pagination options
  * Enhancement [owncloud/web#6173](https://github.com/owncloud/web/pull/6173): Remember the UI that was last selected via the application switcher
  * Enhancement [owncloud/web#7584](https://github.com/owncloud/web/pull/7584): Remove clickOutside directive
  * Enhancement [owncloud/web#7485](https://github.com/owncloud/web/pull/7485): Add resource name to the WebDAV properties
  * Enhancement [owncloud/web#7559](https://github.com/owncloud/web/pull/7559): Don't open right sidebar from private links
  * Enhancement [owncloud/web#7586](https://github.com/owncloud/web/pull/7586): Search improvements
  * Enhancement [owncloud/web#7605](https://github.com/owncloud/web/pull/7605): Simplify mime type checking
  * Enhancement [owncloud/web#7626](https://github.com/owncloud/web/pull/7626): Update ODS to v14.0.0-alpha.18
  * Enhancement [owncloud/web#7177](https://github.com/owncloud/web/issues/7177): Update Uppy to v3.0.1
  * Enhancement [owncloud/web#7182](https://github.com/owncloud/web/pull/7182): User management app edit quota

   https://github.com/owncloud/ocis/pull/4508
   https://github.com/owncloud/ocis/pull/4547
   https://github.com/owncloud/ocis/pull/4550
   https://github.com/owncloud/web/releases/tag/v5.7.0

* Enhancement - Expand personal drive on the graph user: [#4357](https://github.com/owncloud/ocis/pull/4357)

   We can now list the personal drive on the users endpoint via the graph API. A user can add an
   `$expand=drive` query to list the personal drive of the requested user.

   https://github.com/owncloud/ocis/pull/4357

* Enhancement - Validate space names: [#4955](https://github.com/owncloud/ocis/pull/4955)

   We now return `BAD REQUEST` when space names are - too long (max 255 characters) - containing
   evil characters (`/`, `\`, `.`, `\\`, `:`, `?`, `*`, `"`, `>`, `<`, `|`)

   Additionally leading and trailing spaces will be removed silently.

   https://github.com/owncloud/ocis/pull/4955

* Enhancement - Add descriptions to webdav configuration: [#3755](https://github.com/owncloud/ocis/pull/3755)

   Added descriptions to webdav config structs to include them in the config documentation.

   https://github.com/owncloud/ocis/pull/3755

* Enhancement - Search service at the old webdav endpoint: [#4118](https://github.com/owncloud/ocis/pull/4118)

   We made the search service available for legacy clients at the old webdav endpoint.

   https://github.com/owncloud/ocis/pull/4118

* Enhancement - Make it possible to configure a WOPI folderurl: [#4716](https://github.com/owncloud/ocis/pull/4716)

   The wopi folder URL is used to jump back from an application to the containing folder in the files
   list.

   https://github.com/owncloud/ocis/pull/4716
# Changelog for [1.20.0] (2022-04-13)

The following sections list the changes for 1.20.0.

[1.20.0]: https://github.com/owncloud/ocis/compare/v1.19.1...v1.20.0

## Summary

* Bugfix - Add `owncloudsql` driver to authprovider config: [#3435](https://github.com/owncloud/ocis/pull/3435)
* Bugfix - Corrected documentation: [#3439](https://github.com/owncloud/ocis/pull/3439)
* Bugfix - Ensure the same data on /ocs/v?.php/config like oC10: [#3113](https://github.com/owncloud/ocis/pull/3113)
* Bugfix - Use the default server download protocol if spaces are not supported: [#3386](https://github.com/owncloud/ocis/pull/3386)
* Change - Fix keys with underscores in the config files: [#3412](https://github.com/owncloud/ocis/pull/3412)
* Change - Don't create demo users by default: [#3474](https://github.com/owncloud/ocis/pull/3474)
* Enhancement - Alias links: [#3454](https://github.com/owncloud/ocis/pull/3454)
* Enhancement - Replace deprecated String.prototype.substr(): [#3448](https://github.com/owncloud/ocis/pull/3448)
* Enhancement - Add sorting to GraphAPI users and groups: [#3360](https://github.com/owncloud/ocis/issues/3360)
* Enhancement - Unify LDAP config settings across services: [#3476](https://github.com/owncloud/ocis/pull/3476)
* Enhancement - Make config dir configurable: [#3440](https://github.com/owncloud/ocis/pull/3440)
* Enhancement - Use embeddable ocdav go micro service: [#3397](https://github.com/owncloud/ocis/pull/3397)
* Enhancement - Update reva to v2.2.0: [#3397](https://github.com/owncloud/ocis/pull/3397)
* Enhancement - Update ownCloud Web to v5.4.0: [#6709](https://github.com/owncloud/web/pull/6709)
* Enhancement - Implement audit events for user and groups: [#3467](https://github.com/owncloud/ocis/pull/3467)

## Details

* Bugfix - Add `owncloudsql` driver to authprovider config: [#3435](https://github.com/owncloud/ocis/pull/3435)

   https://github.com/owncloud/ocis/pull/3435

* Bugfix - Corrected documentation: [#3439](https://github.com/owncloud/ocis/pull/3439)

   - ocis-pkg log File Option

   https://github.com/owncloud/ocis/pull/3439

* Bugfix - Ensure the same data on /ocs/v?.php/config like oC10: [#3113](https://github.com/owncloud/ocis/pull/3113)

   We've fixed the returned values on the /ocs/v?.php/config endpoints, so that they now return
   the same values as an oC10 would do.

   https://github.com/owncloud/ocis/pull/3113

* Bugfix - Use the default server download protocol if spaces are not supported: [#3386](https://github.com/owncloud/ocis/pull/3386)

   https://github.com/owncloud/ocis/pull/3386

* Change - Fix keys with underscores in the config files: [#3412](https://github.com/owncloud/ocis/pull/3412)

   We've fixed some config keys in configuration files that previously didn't contain
   underscores but should.

   Please check the documentation on https://owncloud.dev for latest configuration
   documentation.

   https://github.com/owncloud/ocis/pull/3412

* Change - Don't create demo users by default: [#3474](https://github.com/owncloud/ocis/pull/3474)

   As we are coming closer to the first beta, we need to disable the creation of the demo users by
   default.

   https://github.com/owncloud/ocis/issues/3181
   https://github.com/owncloud/ocis/pull/3474

* Enhancement - Alias links: [#3454](https://github.com/owncloud/ocis/pull/3454)

   Bumps reva and configures ocs token endpoint to be unprotected

   https://github.com/owncloud/ocis/pull/3454

* Enhancement - Replace deprecated String.prototype.substr(): [#3448](https://github.com/owncloud/ocis/pull/3448)

   We've replaced all occurrences of the deprecated String.prototype.substr() function with
   String.prototype.slice() which works similarly but isn't deprecated.

   https://github.com/owncloud/ocis/pull/3448

* Enhancement - Add sorting to GraphAPI users and groups: [#3360](https://github.com/owncloud/ocis/issues/3360)

   The GraphAPI endpoints for users and groups support ordering now. User can be ordered by
   displayName, onPremisesSamAccountName and mail. Groups can be ordered by displayName.

   Example: https://localhost:9200/graph/v1.0/groups?$orderby=displayName asc

   https://github.com/owncloud/ocis/issues/3360

* Enhancement - Unify LDAP config settings across services: [#3476](https://github.com/owncloud/ocis/pull/3476)

   The storage services where updated to adapt for the recent changes of the LDAP settings in reva.

   Also we allow now to use a new set of top-level LDAP environment variables that are shared
   between all LDAP-using services in ocis (graph, idp, storage-auth-basic,
   storage-userprovider, storage-groupprovider, idm). This should simplify the most LDAP
   based configurations considerably.

   Here is a list of the new environment variables: LDAP_URI LDAP_INSECURE LDAP_CACERT
   LDAP_BIND_DN LDAP_BIND_PASSWORD LDAP_LOGIN_ATTRIBUTES LDAP_USER_BASE_DN
   LDAP_USER_SCOPE LDAP_USER_FILTER LDAP_USER_OBJECTCLASS LDAP_USER_SCHEMA_MAIL
   LDAP_USER_SCHEMA_DISPLAY_NAME LDAP_USER_SCHEMA_USERNAME LDAP_USER_SCHEMA_ID
   LDAP_USER_SCHEMA_ID_IS_OCTETSTRING LDAP_GROUP_BASE_DN LDAP_GROUP_SCOPE
   LDAP_GROUP_FILTER LDAP_GROUP_OBJECTCLASS LDAP_GROUP_SCHEMA_GROUPNAME
   LDAP_GROUP_SCHEMA_ID LDAP_GROUP_SCHEMA_ID_IS_OCTETSTRING

   Where need these can be overwritten by service specific variables. E.g. it is possible to use
   STORAGE_LDAP_URI to override the top-level LDAP_URI variable.

   https://github.com/owncloud/ocis/issues/3150
   https://github.com/owncloud/ocis/pull/3476

* Enhancement - Make config dir configurable: [#3440](https://github.com/owncloud/ocis/pull/3440)

   We have added an `OCIS_CONFIG_DIR` environment variable the will take precedence over the
   default `/etc/ocis`, `~/.ocis` and `.config` locations. When it is set the default locations
   will be ignored and only the configuration files in that directory will be read.

   https://github.com/owncloud/ocis/pull/3440

* Enhancement - Use embeddable ocdav go micro service: [#3397](https://github.com/owncloud/ocis/pull/3397)

   We now use the reva `pgk/micro/ocdav` package that implements a go micro compatible version of
   the ocdav service.

   https://github.com/owncloud/ocis/pull/3397

* Enhancement - Update reva to v2.2.0: [#3397](https://github.com/owncloud/ocis/pull/3397)

   Updated reva to version 2.2.0. This update includes:

  * Bugfix [cs3org/reva#3373](https://github.com/cs3org/reva/pull/3373):  Fix the permissions attribute in propfind responses
  * Bugfix [cs3org/reva#2721](https://github.com/cs3org/reva/pull/2721):  Fix locking and public link scope checker to make the WOPI server work
  * Bugfix [cs3org/reva#2668](https://github.com/cs3org/reva/pull/2668):  Minor cleanup
  * Bugfix [cs3org/reva#2692](https://github.com/cs3org/reva/pull/2692):  Ensure that the host in the ocs config endpoint has no protocol
  * Bugfix [cs3org/reva#2709](https://github.com/cs3org/reva/pull/2709):  Decomposed FS: return precondition failed if already locked
  * Change [cs3org/reva#2687](https://github.com/cs3org/reva/pull/2687):  Allow link with no or edit permission
  * Change [cs3org/reva#2658](https://github.com/cs3org/reva/pull/2658):  Small clean up of the ocdav code
  * Change [cs3org/reva#2691](https://github.com/cs3org/reva/pull/2691):  Decomposed FS: return a reference to the parent
  * Enhancement [cs3org/reva#2708](https://github.com/cs3org/reva/pull/2708):  Rework LDAP configuration of user and group providers
  * Enhancement [cs3org/reva#2665](https://github.com/cs3org/reva/pull/2665):  Add embeddable ocdav go micro service
  * Enhancement [cs3org/reva#2715](https://github.com/cs3org/reva/pull/2715):  Introduced quicklinks
  * Enhancement [cs3org/reva#3370](https://github.com/cs3org/reva/pull/3370):  Enable all spaces members to list public shares
  * Enhancement [cs3org/reva#3370](https://github.com/cs3org/reva/pull/3370):  Enable space members to list shares inside the space
  * Enhancement [cs3org/reva#2717](https://github.com/cs3org/reva/pull/2717):  Add definitions for user and group events

   https://github.com/owncloud/ocis/pull/3397
   https://github.com/owncloud/ocis/pull/3430
   https://github.com/owncloud/ocis/pull/3476
   https://github.com/owncloud/ocis/pull/3482
   https://github.com/owncloud/ocis/pull/3497
   https://github.com/owncloud/ocis/pull/3513
   https://github.com/owncloud/ocis/pull/3514

* Enhancement - Update ownCloud Web to v5.4.0: [#6709](https://github.com/owncloud/web/pull/6709)

   Tags: web

   We updated ownCloud Web to v5.4.0. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/web/pull/6709
   https://github.com/owncloud/ocis/pull/3437
   https://github.com/owncloud/ocis/pull/3487
   https://github.com/owncloud/ocis/pull/3509
   https://github.com/owncloud/web/releases/tag/v5.4.0

* Enhancement - Implement audit events for user and groups: [#3467](https://github.com/owncloud/ocis/pull/3467)

   Added audit events for users and groups. This will log: * User creation * User deletion * User
   property change (currently only email) * Group creation * Group deletion * Group member add *
   Group member remove

   https://github.com/owncloud/ocis/pull/3467
# Changelog for [1.19.1] (2022-03-29)

The following sections list the changes for 1.19.1.

[1.19.1]: https://github.com/owncloud/ocis/compare/v1.19.0...v1.19.1

## Summary

* Bugfix - Return correct special item urls: [#3419](https://github.com/owncloud/ocis/pull/3419)

## Details

* Bugfix - Return correct special item urls: [#3419](https://github.com/owncloud/ocis/pull/3419)

   URLs for Special items (space image, readme) were broken.

   https://github.com/owncloud/ocis/pull/3419
# Changelog for [1.19.0] (2022-03-29)

The following sections list the changes for 1.19.0.

[1.19.0]: https://github.com/owncloud/ocis/compare/v1.18.0...v1.19.0

## Summary

* Bugfix - Network configuration in individual_services example: [#3238](https://github.com/owncloud/ocis/pull/3238)
* Bugfix - Improve gif thumbnails: [#3305](https://github.com/owncloud/ocis/pull/3305)
* Bugfix - Fix error handling in GraphAPI GetUsers call: [#3357](https://github.com/owncloud/ocis/pull/3357)
* Bugfix - Fix request validation on GraphAPI User updates: [#3167](https://github.com/owncloud/ocis/issues/3167)
* Bugfix - Replace public mountpoint fileid with grant fileid: [#3349](https://github.com/owncloud/ocis/pull/3349)
* Change - Add remote item to mountpoint and fix spaceID: [#3365](https://github.com/owncloud/ocis/pull/3365)
* Change - Switch NATS backend: [#3192](https://github.com/owncloud/ocis/pull/3192)
* Change - Drop json config file support: [#3366](https://github.com/owncloud/ocis/pull/3366)
* Change - Settings service now stores its data via metadata service: [#3232](https://github.com/owncloud/ocis/pull/3232)
* Enhancement - Audit logger will now log file events: [#3332](https://github.com/owncloud/ocis/pull/3332)
* Enhancement - Add password reset link to login page: [#3329](https://github.com/owncloud/ocis/pull/3329)
* Enhancement - Log sharing events in audit service: [#3301](https://github.com/owncloud/ocis/pull/3301)
* Enhancement - Add space aliases: [#3283](https://github.com/owncloud/ocis/pull/3283)
* Enhancement - Include etags in drives listing: [#3267](https://github.com/owncloud/ocis/pull/3267)
* Enhancement - Improve thumbnails API: [#3272](https://github.com/owncloud/ocis/pull/3272)
* Enhancement - Update reva to v2.1.0: [#3330](https://github.com/owncloud/ocis/pull/3330)
* Enhancement - Update ownCloud Web to v5.3.0: [#6561](https://github.com/owncloud/web/pull/6561)

## Details

* Bugfix - Network configuration in individual_services example: [#3238](https://github.com/owncloud/ocis/pull/3238)

   Tidy up the deployments/examples/ocis_individual_services example so that the
   instructions work.

   https://github.com/owncloud/ocis/pull/3238

* Bugfix - Improve gif thumbnails: [#3305](https://github.com/owncloud/ocis/pull/3305)

   Improved the gif thumbnail generation for gifs with different disposal strategies.

   https://github.com/owncloud/ocis/pull/3305

* Bugfix - Fix error handling in GraphAPI GetUsers call: [#3357](https://github.com/owncloud/ocis/pull/3357)

   A missing return statement caused GetUsers to return misleading results when the identity
   backend returned an error.

   https://github.com/owncloud/ocis/pull/3357

* Bugfix - Fix request validation on GraphAPI User updates: [#3167](https://github.com/owncloud/ocis/issues/3167)

   Fix PATCH on graph/v1.0/users when no 'mail' attribute is present in the request body

   https://github.com/owncloud/ocis/issues/3167

* Bugfix - Replace public mountpoint fileid with grant fileid: [#3349](https://github.com/owncloud/ocis/pull/3349)

   We now show the same resource id for resources when accessing them via a public links as when
   using a logged in user. This allows the web ui to start a WOPI session with the correct resource
   id.

   https://github.com/owncloud/ocis/pull/3349

* Change - Add remote item to mountpoint and fix spaceID: [#3365](https://github.com/owncloud/ocis/pull/3365)

   A mountpoint represents the mounted share on the share receivers side. The original resource
   is located where the grant has been set. This item is now shown as libregraph remoteItem on the
   mountpoint. While adding this, we fixed the spaceID for mountpoints.

   https://github.com/owncloud/ocis/pull/3365

* Change - Switch NATS backend: [#3192](https://github.com/owncloud/ocis/pull/3192)

   We've switched the NATS backend from Streaming to JetStream, since NATS Streaming is
   depreciated.

   https://github.com/owncloud/ocis/pull/3192
   https://github.com/cs3org/reva/pull/2574

* Change - Drop json config file support: [#3366](https://github.com/owncloud/ocis/pull/3366)

   We've remove the support to configure oCIS and it's service with a json file. From now on we only
   support yaml configuration files, since they have the possibility to add comments.

   https://github.com/owncloud/ocis/pull/3366

* Change - Settings service now stores its data via metadata service: [#3232](https://github.com/owncloud/ocis/pull/3232)

   Instead of writing files to disk it will use metadata service to do so

   https://github.com/owncloud/ocis/pull/3232

* Enhancement - Audit logger will now log file events: [#3332](https://github.com/owncloud/ocis/pull/3332)

   See full list of supported events in `audit/pkg/types/types.go`

   https://github.com/owncloud/ocis/pull/3332

* Enhancement - Add password reset link to login page: [#3329](https://github.com/owncloud/ocis/pull/3329)

   Added a configurable password reset link to the login page. It can be set via
   `IDP_PASSWORD_RESET_URI`. If the option is not set the link will not be shown.

   https://github.com/owncloud/ocis/pull/3329

* Enhancement - Log sharing events in audit service: [#3301](https://github.com/owncloud/ocis/pull/3301)

   Contains sharing related events. See full list in audit/pkg/types/events.go

   https://github.com/owncloud/ocis/pull/3301

* Enhancement - Add space aliases: [#3283](https://github.com/owncloud/ocis/pull/3283)

   Space aliases can be used to resolve spaceIDs in a client.

   https://github.com/owncloud/ocis/pull/3283

* Enhancement - Include etags in drives listing: [#3267](https://github.com/owncloud/ocis/pull/3267)

   Added etags in the response of list drives.

   https://github.com/owncloud/ocis/pull/3267

* Enhancement - Improve thumbnails API: [#3272](https://github.com/owncloud/ocis/pull/3272)

   Changed the thumbnails API to no longer transfer images via GRPC. GRPC has a limited message
   size and isn't very efficient with large binary data. The new API transports the images over
   HTTP.

   https://github.com/owncloud/ocis/pull/3272

* Enhancement - Update reva to v2.1.0: [#3330](https://github.com/owncloud/ocis/pull/3330)

   Updated reva to version 2.1.0. This update includes:

  * Fix [cs3org/reva#2636](https://github.com/cs3org/reva/pull/2636): Delay reconnect log for events
  * Fix [cs3org/reva#2645](https://github.com/cs3org/reva/pull/2645): Avoid warning about missing .flock files
  * Fix [cs3org/reva#2625](https://github.com/cs3org/reva/pull/2625): Fix locking on public links and the decomposed filesystem
  * Fix [cs3org/reva#2643](https://github.com/cs3org/reva/pull/2643): Emit linkaccessfailed event when share is nil
  * Fix [cs3org/reva#2646](https://github.com/cs3org/reva/pull/2646): Replace public mountpoint fileid with grant fileid in ocdav
  * Fix [cs3org/reva#2612](https://github.com/cs3org/reva/pull/2612): Adjust the scope handling to support the spaces architecture
  * Fix [cs3org/reva#2621](https://github.com/cs3org/reva/pull/2621): Send events only if response code is `OK`
  * Chg [cs3org/reva#2574](https://github.com/cs3org/reva/pull/2574): Switch NATS backend
  * Chg [cs3org/reva#2667](https://github.com/cs3org/reva/pull/2667): Allow LDAP groups to have no gidNumber
  * Chg [cs3org/reva#3233](https://github.com/cs3org/reva/pull/3233): Improve quota handling
  * Chg [cs3org/reva#2600](https://github.com/cs3org/reva/pull/2600): Use the cs3 share api to manage spaces
  * Enh [cs3org/reva#2644](https://github.com/cs3org/reva/pull/2644): Add new public share manager
  * Enh [cs3org/reva#2626](https://github.com/cs3org/reva/pull/2626): Add new share manager
  * Enh [cs3org/reva#2624](https://github.com/cs3org/reva/pull/2624): Add etags to virtual spaces
  * Enh [cs3org/reva#2639](https://github.com/cs3org/reva/pull/2639): File Events
  * Enh [cs3org/reva#2627](https://github.com/cs3org/reva/pull/2627): Add events for sharing action
  * Enh [cs3org/reva#2664](https://github.com/cs3org/reva/pull/2664): Add grantID to mountpoint
  * Enh [cs3org/reva#2622](https://github.com/cs3org/reva/pull/2622): Allow listing shares in spaces via the OCS API
  * Enh [cs3org/reva#2623](https://github.com/cs3org/reva/pull/2623): Add space aliases
  * Enh [cs3org/reva#2647](https://github.com/cs3org/reva/pull/2647): Add space specific events
  * Enh [cs3org/reva#3345](https://github.com/cs3org/reva/pull/3345): Add the spaceid to propfind responses
  * Enh [cs3org/reva#2616](https://github.com/cs3org/reva/pull/2616): Add etag to spaces response
  * Enh [cs3org/reva#2628](https://github.com/cs3org/reva/pull/2628): Add spaces aware trash-bin API

   https://github.com/owncloud/ocis/pull/3330
   https://github.com/owncloud/ocis/pull/3405
   https://github.com/owncloud/ocis/pull/3416

* Enhancement - Update ownCloud Web to v5.3.0: [#6561](https://github.com/owncloud/web/pull/6561)

   Tags: web

   We updated ownCloud Web to v5.3.0. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/web/pull/6561
   https://github.com/owncloud/ocis/pull/3291
   https://github.com/owncloud/ocis/pull/3375
   https://github.com/owncloud/web/releases/tag/v5.3.0
# Changelog for [1.18.0] (2022-03-03)

The following sections list the changes for 1.18.0.

[1.18.0]: https://github.com/owncloud/ocis/compare/v1.17.0...v1.18.0

## Summary

* Bugfix - Capabilities for password protected public links: [#3229](https://github.com/owncloud/ocis/pull/3229)
* Bugfix - Make events settings configurable: [#3214](https://github.com/owncloud/ocis/pull/3214)
* Bugfix - Align storage metadata GPRC bind port with other variable names: [#3169](https://github.com/owncloud/ocis/pull/3169)
* Change - Unify file IDs: [#3185](https://github.com/owncloud/ocis/pull/3185)
* Enhancement - Add sorting to list Spaces: [#3200](https://github.com/owncloud/ocis/issues/3200)
* Enhancement - Change NATS port: [#3210](https://github.com/owncloud/ocis/pull/3210)
* Enhancement - Re-Enabling web cache control: [#3109](https://github.com/owncloud/ocis/pull/3109)
* Enhancement - Add SPA conform fileserver for web: [#3109](https://github.com/owncloud/ocis/pull/3109)
* Enhancement - Implement notifications service: [#3217](https://github.com/owncloud/ocis/pull/3217)
* Enhancement - Thumbnails in spaces: [#3219](https://github.com/owncloud/ocis/pull/3219)
* Enhancement - Update reva to v2.0.0: [#3231](https://github.com/owncloud/ocis/pull/3231)
* Enhancement - Update ownCloud Web to v5.2.0: [#6506](https://github.com/owncloud/web/pull/6506)

## Details

* Bugfix - Capabilities for password protected public links: [#3229](https://github.com/owncloud/ocis/pull/3229)

   Allow password protected public links to request capabilities.

   https://github.com/owncloud/web/issues/5863
   https://github.com/owncloud/ocis/pull/3229
   https://github.com/owncloud/web/pull/6471

* Bugfix - Make events settings configurable: [#3214](https://github.com/owncloud/ocis/pull/3214)

   We've fixed the hardcoded events settings to be configurable.

   https://github.com/owncloud/ocis/pull/3214

* Bugfix - Align storage metadata GPRC bind port with other variable names: [#3169](https://github.com/owncloud/ocis/pull/3169)

   Changed STORAGE_METADATA_GRPC_PROVIDER_ADDR to STORAGE_METADATA_GRPC_ADDR so it aligns
   with standard environment variable naming conventions used in oCIS.

   https://github.com/owncloud/ocis/pull/3169

* Change - Unify file IDs: [#3185](https://github.com/owncloud/ocis/pull/3185)

   We changed the file IDs to be consistent across all our APIs (WebDAV, LibreGraph, OCS). We
   removed the base64 encoding. Now they are formatted like <storageID>!<opaqueID>. They are
   using a reserved character ``!`` as a URL safe separator.

   https://github.com/owncloud/ocis/pull/3185

* Enhancement - Add sorting to list Spaces: [#3200](https://github.com/owncloud/ocis/issues/3200)

   We added the OData query param "orderBy" for listing spaces. We can now order by Space Name and
   LastModifiedDateTime.

   Example 1:
   https://localhost:9200/graph/v1.0/me/drives/?$orderby=lastModifiedDateTime desc
   Example 2: https://localhost:9200/graph/v1.0/me/drives/?$orderby=name asc

   https://github.com/owncloud/ocis/issues/3200
   https://github.com/owncloud/ocis/pull/3201
   https://github.com/owncloud/ocis/pull/3218

* Enhancement - Change NATS port: [#3210](https://github.com/owncloud/ocis/pull/3210)

   Currently only a certain range of ports is allowed for ocis application. Use a supported port
   for nats server

   https://github.com/owncloud/ocis/pull/3210

* Enhancement - Re-Enabling web cache control: [#3109](https://github.com/owncloud/ocis/pull/3109)

   We've re-enable browser caching headers (`Expires` and `Last-Modified`) for the web
   service, this was disabled due to a problem in the fileserver used before. Since we're now using
   our own fileserver implementation this works again and is enabled by default.

   https://github.com/owncloud/ocis/pull/3109

* Enhancement - Add SPA conform fileserver for web: [#3109](https://github.com/owncloud/ocis/pull/3109)

   We've added an SPA conform fileserver to the web service. It enables web to use vue's history
   mode and behaves like nginx try_files.

   https://github.com/owncloud/ocis/pull/3109

* Enhancement - Implement notifications service: [#3217](https://github.com/owncloud/ocis/pull/3217)

   Implemented the minimal version of the notifications service to be able to notify a user when
   they received a share.

   https://github.com/owncloud/ocis/pull/3217

* Enhancement - Thumbnails in spaces: [#3219](https://github.com/owncloud/ocis/pull/3219)

   Added support for thumbnails in spaces.

   https://github.com/owncloud/ocis/pull/3219
   https://github.com/owncloud/ocis/pull/3235

* Enhancement - Update reva to v2.0.0: [#3231](https://github.com/owncloud/ocis/pull/3231)

   We updated reva to the version 2.0.0.

  * Fix [cs3org/reva#2457](https://github.com/cs3org/reva/pull/2457) :  Do not swallow error
  * Fix [cs3org/reva#2422](https://github.com/cs3org/reva/pull/2422) :  Handle non existing spaces correctly
  * Fix [cs3org/reva#2327](https://github.com/cs3org/reva/pull/2327) :  Enable changelog on edge branch
  * Fix [cs3org/reva#2370](https://github.com/cs3org/reva/pull/2370) :  Fixes for apps in public shares, project spaces for EOS driver
  * Fix [cs3org/reva#2464](https://github.com/cs3org/reva/pull/2464) :  Pass spacegrants when adding member to space
  * Fix [cs3org/reva#2430](https://github.com/cs3org/reva/pull/2430) :  Fix aggregated child folder id
  * Fix [cs3org/reva#2348](https://github.com/cs3org/reva/pull/2348) :  Make archiver handle spaces protocol
  * Fix [cs3org/reva#2452](https://github.com/cs3org/reva/pull/2452) :  Fix create space error message
  * Fix [cs3org/reva#2445](https://github.com/cs3org/reva/pull/2445) :  Don't handle ids containing "/" in decomposedfs
  * Fix [cs3org/reva#2285](https://github.com/cs3org/reva/pull/2285) :  Accept new userid idp format
  * Fix [cs3org/reva#2503](https://github.com/cs3org/reva/pull/2503) :  Remove the protection from /v?.php/config endpoints
  * Fix [cs3org/reva#2462](https://github.com/cs3org/reva/pull/2462) :  Public shares path needs to be set
  * Fix [cs3org/reva#2427](https://github.com/cs3org/reva/pull/2427) :  Fix registry caching
  * Fix [cs3org/reva#2298](https://github.com/cs3org/reva/pull/2298) :  Remove share refs from trashbin
  * Fix [cs3org/reva#2433](https://github.com/cs3org/reva/pull/2433) :  Fix shares provider filter
  * Fix [cs3org/reva#2351](https://github.com/cs3org/reva/pull/2351) :  Fix Statcache removing
  * Fix [cs3org/reva#2374](https://github.com/cs3org/reva/pull/2374) :  Fix webdav copy of zero byte files
  * Fix [cs3org/reva#2336](https://github.com/cs3org/reva/pull/2336) :  Handle sending all permissions when creating public links
  * Fix [cs3org/reva#2440](https://github.com/cs3org/reva/pull/2440) :  Add ArbitraryMetadataKeys to statcache key
  * Fix [cs3org/reva#2582](https://github.com/cs3org/reva/pull/2582) :  Keep lock structs in a local map protected by a mutex
  * Fix [cs3org/reva#2372](https://github.com/cs3org/reva/pull/2372) :  Make owncloudsql work with the spaces registry
  * Fix [cs3org/reva#2416](https://github.com/cs3org/reva/pull/2416) :  The registry now returns complete space structs
  * Fix [cs3org/reva#3066](https://github.com/cs3org/reva/pull/3066) :  Fix propfind listing for files
  * Fix [cs3org/reva#2428](https://github.com/cs3org/reva/pull/2428) :  Remove unused home provider from config
  * Fix [cs3org/reva#2334](https://github.com/cs3org/reva/pull/2334) :  Revert fix decomposedfs upload
  * Fix [cs3org/reva#2415](https://github.com/cs3org/reva/pull/2415) :  Services should never return transport level errors
  * Fix [cs3org/reva#2419](https://github.com/cs3org/reva/pull/2419) :  List project spaces for share recipients
  * Fix [cs3org/reva#2501](https://github.com/cs3org/reva/pull/2501) :  Fix spaces stat
  * Fix [cs3org/reva#2432](https://github.com/cs3org/reva/pull/2432) :  Use space reference when listing containers
  * Fix [cs3org/reva#2572](https://github.com/cs3org/reva/pull/2572) :  Wait for nats server on middleware start
  * Fix [cs3org/reva#2454](https://github.com/cs3org/reva/pull/2454) :  Fix webdav paths in PROPFINDS
  * Chg [cs3org/reva#2329](https://github.com/cs3org/reva/pull/2329) :  Activate the statcache
  * Chg [cs3org/reva#2596](https://github.com/cs3org/reva/pull/2596) :  Remove hash from public link urls
  * Chg [cs3org/reva#2495](https://github.com/cs3org/reva/pull/2495) :  Remove the ownCloud storage driver
  * Chg [cs3org/reva#2527](https://github.com/cs3org/reva/pull/2527) :  Store space attributes in decomposedFS
  * Chg [cs3org/reva#2581](https://github.com/cs3org/reva/pull/2581) :  Update hard-coded status values
  * Chg [cs3org/reva#2524](https://github.com/cs3org/reva/pull/2524) :  Use description during space creation
  * Chg [cs3org/reva#2554](https://github.com/cs3org/reva/pull/2554) :  Shard nodes per space in decomposedfs
  * Chg [cs3org/reva#2576](https://github.com/cs3org/reva/pull/2576) :  Harden xattrs errors
  * Chg [cs3org/reva#2436](https://github.com/cs3org/reva/pull/2436) :  Replace template in GroupFilter for UserProvider with a simple string
  * Chg [cs3org/reva#2429](https://github.com/cs3org/reva/pull/2429) :  Make archiver id based
  * Chg [cs3org/reva#2340](https://github.com/cs3org/reva/pull/2340) :  Allow multiple space configurations per provider
  * Chg [cs3org/reva#2396](https://github.com/cs3org/reva/pull/2396) :  The ocdav handler is now spaces aware
  * Chg [cs3org/reva#2349](https://github.com/cs3org/reva/pull/2349) :  Require `ListRecycle` when listing trashbin
  * Chg [cs3org/reva#2353](https://github.com/cs3org/reva/pull/2353) :  Reduce log output
  * Chg [cs3org/reva#2542](https://github.com/cs3org/reva/pull/2542) :  Do not encode webDAV ids to base64
  * Chg [cs3org/reva#2519](https://github.com/cs3org/reva/pull/2519) :  Remove the auto creation of the .space folder
  * Chg [cs3org/reva#2394](https://github.com/cs3org/reva/pull/2394) :  Remove logic from gateway
  * Chg [cs3org/reva#2023](https://github.com/cs3org/reva/pull/2023) :  Add a sharestorageprovider
  * Chg [cs3org/reva#2234](https://github.com/cs3org/reva/pull/2234) :  Add a spaces registry
  * Chg [cs3org/reva#2339](https://github.com/cs3org/reva/pull/2339) :  Fix static registry regressions
  * Chg [cs3org/reva#2370](https://github.com/cs3org/reva/pull/2370) :  Fix static registry regressions
  * Chg [cs3org/reva#2354](https://github.com/cs3org/reva/pull/2354) :  Return not found when updating non existent space
  * Chg [cs3org/reva#2589](https://github.com/cs3org/reva/pull/2589) :  Remove deprecated linter modules
  * Chg [cs3org/reva#2016](https://github.com/cs3org/reva/pull/2016) :  Move wrapping and unwrapping of paths to the storage gateway
  * Enh [cs3org/reva#2591](https://github.com/cs3org/reva/pull/2591) :  Set up App Locks with basic locks
  * Enh [cs3org/reva#1209](https://github.com/cs3org/reva/pull/1209) :  Reva CephFS module v0.2.1
  * Enh [cs3org/reva#2511](https://github.com/cs3org/reva/pull/2511) :  Error handling cleanup in decomposed FS
  * Enh [cs3org/reva#2516](https://github.com/cs3org/reva/pull/2516) :  Cleaned up some code
  * Enh [cs3org/reva#2512](https://github.com/cs3org/reva/pull/2512) :  Consolidate xattr setter and getter
  * Enh [cs3org/reva#2341](https://github.com/cs3org/reva/pull/2341) :  Use CS3 permissions API
  * Enh [cs3org/reva#2343](https://github.com/cs3org/reva/pull/2343) :  Allow multiple space type fileters on decomposedfs
  * Enh [cs3org/reva#2460](https://github.com/cs3org/reva/pull/2460) :  Add locking support to decomposedfs
  * Enh [cs3org/reva#2540](https://github.com/cs3org/reva/pull/2540) :  Refactored the xattrs package in the decomposedfs
  * Enh [cs3org/reva#2463](https://github.com/cs3org/reva/pull/2463) :  Do not log whole nodes
  * Enh [cs3org/reva#2350](https://github.com/cs3org/reva/pull/2350) :  Add file locking methods to the storage and filesystem interfaces
  * Enh [cs3org/reva#2379](https://github.com/cs3org/reva/pull/2379) :  Add new file url of the app provider to the ocs capabilities
  * Enh [cs3org/reva#2369](https://github.com/cs3org/reva/pull/2369) :  Implement TouchFile from the CS3apis
  * Enh [cs3org/reva#2385](https://github.com/cs3org/reva/pull/2385) :  Allow to create new files with the app provider on public links
  * Enh [cs3org/reva#2397](https://github.com/cs3org/reva/pull/2397) :  Product field in OCS version
  * Enh [cs3org/reva#2393](https://github.com/cs3org/reva/pull/2393) :  Update tus/tusd to version 1.8.0
  * Enh [cs3org/reva#2522](https://github.com/cs3org/reva/pull/2522) :  Introduce events
  * Enh [cs3org/reva#2528](https://github.com/cs3org/reva/pull/2528) :  Use an exclusive write lock when writing multiple attributes
  * Enh [cs3org/reva#2595](https://github.com/cs3org/reva/pull/2595) :  Add integration test for the groupprovider
  * Enh [cs3org/reva#2439](https://github.com/cs3org/reva/pull/2439) :  Ignore handled errors when creating spaces
  * Enh [cs3org/reva#2500](https://github.com/cs3org/reva/pull/2500) :  Invalidate listproviders cache
  * Enh [cs3org/reva#2345](https://github.com/cs3org/reva/pull/2345) :  Don't assume that the LDAP groupid in reva matches the name
  * Enh [cs3org/reva#2525](https://github.com/cs3org/reva/pull/2525) :  Allow using AD UUID as userId values
  * Enh [cs3org/reva#2584](https://github.com/cs3org/reva/pull/2584) :  Allow running userprovider integration tests for the LDAP driver
  * Enh [cs3org/reva#2585](https://github.com/cs3org/reva/pull/2585) :  Add metadata storage layer and indexer
  * Enh [cs3org/reva#2163](https://github.com/cs3org/reva/pull/2163) :  Nextcloud-based share manager for pkg/ocm/share
  * Enh [cs3org/reva#2278](https://github.com/cs3org/reva/pull/2278) :  OIDC driver changes for lightweight users
  * Enh [cs3org/reva#2315](https://github.com/cs3org/reva/pull/2315) :  Add new attributes to public link propfinds
  * Enh [cs3org/reva#2431](https://github.com/cs3org/reva/pull/2431) :  Delete shares when purging spaces
  * Enh [cs3org/reva#2434](https://github.com/cs3org/reva/pull/2434) :  Refactor ocdav into smaller chunks
  * Enh [cs3org/reva#2524](https://github.com/cs3org/reva/pull/2524) :  Add checks when removing space members
  * Enh [cs3org/reva#2457](https://github.com/cs3org/reva/pull/2457) :  Restore spaces that were previously deleted
  * Enh [cs3org/reva#2498](https://github.com/cs3org/reva/pull/2498) :  Include grants in list storage spaces response
  * Enh [cs3org/reva#2344](https://github.com/cs3org/reva/pull/2344) :  Allow listing all storage spaces
  * Enh [cs3org/reva#2547](https://github.com/cs3org/reva/pull/2547) :  Add an if-match check to the storage provider
  * Enh [cs3org/reva#2486](https://github.com/cs3org/reva/pull/2486) :  Update cs3apis to include lock api changes
  * Enh [cs3org/reva#2526](https://github.com/cs3org/reva/pull/2526) :  Upgrade ginkgo to v2

   https://github.com/owncloud/ocis/pull/3231
   https://github.com/owncloud/ocis/pull/3258

* Enhancement - Update ownCloud Web to v5.2.0: [#6506](https://github.com/owncloud/web/pull/6506)

   Tags: web

   We updated ownCloud Web to v5.2.0. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/web/pull/6506
   https://github.com/owncloud/ocis/pull/3202
   https://github.com/owncloud/web/releases/tag/v5.2.0
# Changelog for [1.17.0] (2022-02-16)

The following sections list the changes for 1.17.0.

[1.17.0]: https://github.com/owncloud/ocis/compare/v1.16.0...v1.17.0

## Summary

* Bugfix - Add `ocis storage-auth-machine` subcommand: [#2910](https://github.com/owncloud/ocis/pull/2910)
* Bugfix - Use same jwt secret for accounts as for metadata storage: [#3081](https://github.com/owncloud/ocis/pull/3081)
* Bugfix - Make the default grpc client use the registry settings: [#3041](https://github.com/owncloud/ocis/pull/3041)
* Bugfix - Remove group memberships when deleting a user: [#3027](https://github.com/owncloud/ocis/issues/3027)
* Bugfix - Fix retry handling for LDAP connections: [#2974](https://github.com/owncloud/ocis/issues/2974)
* Bugfix - Fix the default tracing provider: [#2952](https://github.com/owncloud/ocis/pull/2952)
* Bugfix - Fix configuration for space membership endpoint: [#2893](https://github.com/owncloud/ocis/pull/2893)
* Change - Change log level default from debug to error: [#3071](https://github.com/owncloud/ocis/pull/3071)
* Change - Remove the ownCloud storage driver: [#3072](https://github.com/owncloud/ocis/pull/3072)
* Change - Unify configuration and commands: [#2818](https://github.com/owncloud/ocis/pull/2818)
* Change - Functionality to restore spaces: [#3092](https://github.com/owncloud/ocis/pull/3092)
* Change - Extended Space Properties: [#3141](https://github.com/owncloud/ocis/pull/3141)
* Change - Update the graph api: [#2885](https://github.com/owncloud/ocis/pull/2885)
* Change - Update libre-graph-api to v0.3.0: [#2858](https://github.com/owncloud/ocis/pull/2858)
* Change - Return not found when updating non existent space: [#2869](https://github.com/cs3org/reva/pull/2869)
* Enhancement - Provide Description when creating a space: [#3167](https://github.com/owncloud/ocis/pull/3167)
* Enhancement - Add graph endpoint to delete and purge spaces: [#2979](https://github.com/owncloud/ocis/pull/2979)
* Enhancement - Add permissions to graph drives: [#3095](https://github.com/owncloud/ocis/pull/3095)
* Enhancement - Add new file url of the app provider to the ocs capabilities: [#2884](https://github.com/owncloud/ocis/pull/2884)
* Enhancement - Add spaces capability: [#2931](https://github.com/owncloud/ocis/pull/2931)
* Enhancement - Consul as supported service registry: [#3133](https://github.com/owncloud/ocis/pull/3133)
* Enhancement - Introduce User and Group Management capabilities on GraphAPI: [#2947](https://github.com/owncloud/ocis/pull/2947)
* Enhancement - Support signature auth in the public share auth middleware: [#2831](https://github.com/owncloud/ocis/pull/2831)
* Enhancement - Update REVA to v1.16.1-0.20220112085026-07451f6cd806: [#2953](https://github.com/owncloud/ocis/pull/2953)
* Enhancement - Add endpoint to retrieve a single space: [#2978](https://github.com/owncloud/ocis/pull/2978)
* Enhancement - Add filter by driveType and id to /me/drives: [#2946](https://github.com/owncloud/ocis/pull/2946)
* Enhancement - Update REVA to v1.16.1-0.20220215130802-df1264deff58: [#2878](https://github.com/owncloud/ocis/pull/2878)
* Enhancement - Update ownCloud Web to v5.0.0: [#2895](https://github.com/owncloud/ocis/pull/2895)

## Details

* Bugfix - Add `ocis storage-auth-machine` subcommand: [#2910](https://github.com/owncloud/ocis/pull/2910)

   We added the ocis subcommand to start the machine auth provider.

   https://github.com/owncloud/ocis/pull/2910

* Bugfix - Use same jwt secret for accounts as for metadata storage: [#3081](https://github.com/owncloud/ocis/pull/3081)

   We've the metadata storage uses the same jwt secret as all other REVA services. Therefore the
   accounts service needs to use the same secret.

   Secrets are documented here:
   https://owncloud.dev/ocis/deployment/#change-default-secrets

   https://github.com/owncloud/ocis/pull/3081

* Bugfix - Make the default grpc client use the registry settings: [#3041](https://github.com/owncloud/ocis/pull/3041)

   We've fixed the default grpc client to use the registry settings. Previously it always used
   mdns.

   https://github.com/owncloud/ocis/pull/3041

* Bugfix - Remove group memberships when deleting a user: [#3027](https://github.com/owncloud/ocis/issues/3027)

   The LDAP backend in the graph API now takes care of removing a user's group membership when
   deleting the user.

   https://github.com/owncloud/ocis/issues/3027

* Bugfix - Fix retry handling for LDAP connections: [#2974](https://github.com/owncloud/ocis/issues/2974)

   We've fixed the handling of network issues (e.g. connection loss) during LDAP Write
   Operations to correctly retry the request.

   https://github.com/owncloud/ocis/issues/2974

* Bugfix - Fix the default tracing provider: [#2952](https://github.com/owncloud/ocis/pull/2952)

   We've fixed the default tracing provider which was no longer configured after
   [owncloud/ocis#2818](https://github.com/owncloud/ocis/pull/2818).

   https://github.com/owncloud/ocis/pull/2952
   https://github.com/owncloud/ocis/pull/2818

* Bugfix - Fix configuration for space membership endpoint: [#2893](https://github.com/owncloud/ocis/pull/2893)

   Added a missing config value to the ocs config related to the space membership endpoint.

   https://github.com/owncloud/ocis/pull/2893

* Change - Change log level default from debug to error: [#3071](https://github.com/owncloud/ocis/pull/3071)

   We've changed the default log level for all services from "info" to "error".

   https://github.com/owncloud/ocis/pull/3071

* Change - Remove the ownCloud storage driver: [#3072](https://github.com/owncloud/ocis/pull/3072)

   We've removed the ownCloud storage driver because it was no longer maintained after the
   ownCloud SQL storage driver was added.

   If you have been using the ownCloud storage driver, please switch to the ownCloud SQL storage
   driver which brings you more features and is under active maintenance.

   https://github.com/owncloud/ocis/pull/3072

* Change - Unify configuration and commands: [#2818](https://github.com/owncloud/ocis/pull/2818)

   We've unified the configuration and commands of all non storage services. This also includes
   the change, that environment variables are now defined on the config struct as tags instead in a
   separate mapping.

   https://github.com/owncloud/ocis/pull/2818

* Change - Functionality to restore spaces: [#3092](https://github.com/owncloud/ocis/pull/3092)

   Disabled spaces can now be restored via the graph api. An information was added to the root item
   of each space when it is deleted

   https://github.com/owncloud/ocis/pull/3092

* Change - Extended Space Properties: [#3141](https://github.com/owncloud/ocis/pull/3141)

   We can now set and modify short description, space image and space readme. Only managers can set
   the short description. Editors can change the space image and readme id.

   https://github.com/owncloud/ocis/pull/3141

* Change - Update the graph api: [#2885](https://github.com/owncloud/ocis/pull/2885)

   GraphApi has been updated to version 0.4.1 and the existing dependency was removed

   https://github.com/owncloud/ocis/pull/2885

* Change - Update libre-graph-api to v0.3.0: [#2858](https://github.com/owncloud/ocis/pull/2858)

   This updates the libre-graph-api to use the latest spec and types.

   https://github.com/owncloud/ocis/pull/2858

* Change - Return not found when updating non existent space: [#2869](https://github.com/cs3org/reva/pull/2869)

   If a spaceid of a space which is updated doesn't exist, handle it as a not found error.

   https://github.com/cs3org/reva/pull/2869

* Enhancement - Provide Description when creating a space: [#3167](https://github.com/owncloud/ocis/pull/3167)

   We added the possibility to send a short description when creating a space.

   https://github.com/owncloud/ocis/pull/3167

* Enhancement - Add graph endpoint to delete and purge spaces: [#2979](https://github.com/owncloud/ocis/pull/2979)

   Added a new graph endpoint to delete and purge spaces.

   https://github.com/owncloud/ocis/pull/2979
   https://github.com/owncloud/ocis/pull/3000

* Enhancement - Add permissions to graph drives: [#3095](https://github.com/owncloud/ocis/pull/3095)

   Added permissions to graph drives when listing drives.

   https://github.com/owncloud/ocis/pull/3095

* Enhancement - Add new file url of the app provider to the ocs capabilities: [#2884](https://github.com/owncloud/ocis/pull/2884)

   We've added the new file capability of the app provider to the ocs capabilities, so that clients
   can discover this url analogous to the app list and file open urls.

   https://github.com/owncloud/ocis/pull/2884
   https://github.com/owncloud/ocis/pull/2907
   https://github.com/cs3org/reva/pull/2379
   https://github.com/owncloud/web/pull/5890#issuecomment-993905242

* Enhancement - Add spaces capability: [#2931](https://github.com/owncloud/ocis/pull/2931)

   We've added the spaces capability with version 0.0.1 and enabled defaulting to true.

   https://github.com/owncloud/ocis/pull/2931
   https://github.com/cs3org/reva/pull/2015
   https://github.com/owncloud/ocis/pull/2965

* Enhancement - Consul as supported service registry: [#3133](https://github.com/owncloud/ocis/pull/3133)

   We have added Consul as an supported service registry. You can now use it to let oCIS services
   discover each other.

   https://github.com/owncloud/ocis/pull/3133

* Enhancement - Introduce User and Group Management capabilities on GraphAPI: [#2947](https://github.com/owncloud/ocis/pull/2947)

   The GraphAPI LDAP Backend is now able to add/modify and delete Users and Groups

   https://github.com/owncloud/ocis/pull/2947
   https://github.com/owncloud/ocis/pull/2996

* Enhancement - Support signature auth in the public share auth middleware: [#2831](https://github.com/owncloud/ocis/pull/2831)

   Enabled public share requests to be authenticated using the public share signature.

   https://github.com/owncloud/ocis/pull/2831

* Enhancement - Update REVA to v1.16.1-0.20220112085026-07451f6cd806: [#2953](https://github.com/owncloud/ocis/pull/2953)

   Update REVA to v1.16.1-0.20220112085026-07451f6cd806

   https://github.com/owncloud/ocis/pull/2953

* Enhancement - Add endpoint to retrieve a single space: [#2978](https://github.com/owncloud/ocis/pull/2978)

   We added the endpoint ``/drives/{driveID}`` to get a single space by id from the server.

   https://github.com/owncloud/ocis/pull/2978

* Enhancement - Add filter by driveType and id to /me/drives: [#2946](https://github.com/owncloud/ocis/pull/2946)

   We added two possible filter terms (driveType, id) to the /me/drives endpoint on the graph api.
   These can be used with the odata query parameter "$filter". We only support the "eq" operator
   for now.

   https://github.com/owncloud/ocis/pull/2946

* Enhancement - Update REVA to v1.16.1-0.20220215130802-df1264deff58: [#2878](https://github.com/owncloud/ocis/pull/2878)

   Updated REVA to v1.16.1-0.20220215130802-df1264deff58 This update includes:

  * Enh [cs3org/reva#2524](https://github.com/cs3org/reva/pull/2524): Remove space members
  * Fix [cs3org/reva#2541](https://github.com/cs3org/reva/pull/2541): fix xattr error types, remove error wrapper
  * Chg [cs3org/reva#2540](https://github.com/cs3org/reva/pull/2540): decomposedfs: refactor xattrs package errors
  * Enh [cs3org/reva#2533](https://github.com/cs3org/reva/pull/2533): Use space description on creation
  * Enh [cs3org/reva#2527](https://github.com/cs3org/reva/pull/2527): Add space props
  * Enh [cs3org/reva#2522](https://github.com/cs3org/reva/pull/2522): Events
  * Chg [cs3org/reva#2512](https://github.com/cs3org/reva/pull/2512): Consolidate all metadata Get's and Set's to central functions.
  * Chg [cs3org/reva#2511](https://github.com/cs3org/reva/pull/2511): Some error cleanup steps in the decomposed FS
  * Enh [cs3org/reva#2460](https://github.com/cs3org/reva/pull/2460): decomposedfs: add locking support
  * Chg [cs3org/reva#2519](https://github.com/cs3org/reva/pull/2519): remove creation of .space folder
  * Fix [cs3org/reva#2506](https://github.com/cs3org/reva/pull/2506): fix propfind listing for files
  * Chg [cs3org/reva#2503](https://github.com/cs3org/reva/pull/2503): unprotected ocs config endpoint
  * Enh [cs3org/reva#2458](https://github.com/cs3org/reva/pull/2458): Restoring Spaces
  * Enh [cs3org/reva#2498](https://github.com/cs3org/reva/pull/2498): add grants to list-spaces
  * Fix [cs3org/reva#2500](https://github.com/cs3org/reva/pull/2500): invalidate cache when modifying or deleting a space
  * Fix [cs3org/reva#2501](https://github.com/cs3org/reva/pull/2501): fix spaces stat requests
  * Enh [cs3org/reva#2472](https://github.com/cs3org/reva/pull/2472): Make owncloudsql spaces aware
  * Enh [cs3org/reva#2464](https://github.com/cs3org/reva/pull/2464): Space grants
  * Fix [cs3org/reva#2463](https://github.com/cs3org/reva/pull/2463): Do not log nodes
  * Enh [cs3org/reva#2437](https://github.com/cs3org/reva/pull/2437): Make gateway dumb again
  * Enh [cs3org/reva#2459](https://github.com/cs3org/reva/pull/2459): prevent purging of enabled spaces
  * Fix [cs3org/reva#2457](https://github.com/cs3org/reva/pull/2457): decomposedfs: do not swallow errors when creating nodes
  * Fix [cs3org/reva#2454](https://github.com/cs3org/reva/pull/2454): fix path construction in webdav propfind
  * Fix [cs3org/reva#2452](https://github.com/cs3org/reva/pull/2452): fix create space error message
  * Enh [cs3org/reva#2431](https://github.com/cs3org/reva/pull/2431): Purge spaces
  * Fix [cs3org/reva#2445](https://github.com/cs3org/reva/pull/2445): Fix publiclinks and decomposedfs
  * Chg [cs3org/reva#2439](https://github.com/cs3org/reva/pull/2439): ignore handled errors when creating spaces
  * Enh [cs3org/reva#2436](https://github.com/cs3org/reva/pull/2436): Adjust "groupfilter" to be able to search by member name
  * Fix [cs3org/reva#2434](https://github.com/cs3org/reva/pull/2434): Start splitting up ocdav
  * Fix [cs3org/reva#2433](https://github.com/cs3org/reva/pull/2433): fix shares provider filter
  * Chg [cs3org/reva#2432](https://github.com/cs3org/reva/pull/2432): use space reference when listing containers
  * Fix [cs3org/reva#2430](https://github.com/cs3org/reva/pull/2430): fix aggregated child folder id
  * Enh [cs3org/reva#2429](https://github.com/cs3org/reva/pull/2429): make archiver id based
  * Fix [cs3org/reva#2427](https://github.com/cs3org/reva/pull/2427): fix registry caching
  * Fix [cs3org/reva#2422](https://github.com/cs3org/reva/pull/2422): handle space does not exist
  * Fix [cs3org/reva#2419](https://github.com/cs3org/reva/pull/2419): Spaces fixes
  * Chg [cs3org/reva#2415](https://github.com/cs3org/reva/pull/2415): services should never return transport level errors
  * Chg [cs3org/reva#2396](https://github.com/cs3org/reva/pull/2396): Ocdav spaces aware
  * Fix [cs3org/reva#2348](https://github.com/cs3org/reva/pull/2348): fix-archiver
  * Chg [cs3org/reva#2344](https://github.com/cs3org/reva/pull/2344): allow listing all storage spaces
  * Chg [cs3org/reva#2345](https://github.com/cs3org/reva/pull/2345): Switch LDAP test to use entryUUID as unique id for groups
  * Chg [cs3org/reva#2343](https://github.com/cs3org/reva/pull/2343): allow multiple space type filters on decomposedfs
  * Enh [cs3org/reva#2329](https://github.com/cs3org/reva/pull/2329): Activate Statcache
  * Enh [cs3org/reva#2340](https://github.com/cs3org/reva/pull/2340): Space registry multiple spaces per provider
  * Chg [cs3org/reva#2336](https://github.com/cs3org/reva/pull/2336): handle sending all permissions when creating public links
  * Fix [cs3org/reva#2330](https://github.com/cs3org/reva/pull/2330): fix decomposedfs upload
  * Enh [cs3org/reva#2234](https://github.com/cs3org/reva/pull/2234): Spaces registry
  * Enh [cs3org/reva#2217](https://github.com/cs3org/reva/pull/2217): New OIDC ESCAPE auth driver.
  * Enh [cs3org/reva#2250](https://github.com/cs3org/reva/pull/2250): Implement space membership endpoints
  * Fix [cs3org/reva#1941](https://github.com/cs3org/reva/pull/1941): fix tus with transfer token only
  * Fix [cs3org/reva#2309](https://github.com/cs3org/reva/pull/2309): Bugfix: Remove early finish for zero byte file uploads
  * Fix [cs3org/reva#2303](https://github.com/cs3org/reva/pull/2303): Fix content disposition
  * Fix [cs3org/reva#2314](https://github.com/cs3org/reva/pull/2314): OIDC: fallback to "email" if IDP doesn't provide "preferred_username" claim
  * Enh [cs3org/reva#2256](https://github.com/cs3org/reva/pull/2256): Return user type in the response of the ocs GET user call
  * Enh [cs3org/reva#2310](https://github.com/cs3org/reva/pull/2310): Implement setting arbitrary metadata for the public storage provider
  * Fix [cs3org/reva#2305](https://github.com/cs3org/reva/pull/2305): Make sure /app/new takes target as absolute path
  * Fix [cs3org/reva#2297](https://github.com/cs3org/reva/pull/2297): Fix public link paths for file shares

   https://github.com/owncloud/ocis/pull/2878
   https://github.com/owncloud/ocis/pull/2901
   https://github.com/owncloud/ocis/pull/2997
   https://github.com/owncloud/ocis/pull/3116
   https://github.com/owncloud/ocis/pull/3130
   https://github.com/owncloud/ocis/pull/3175
   https://github.com/owncloud/ocis/pull/3182

* Enhancement - Update ownCloud Web to v5.0.0: [#2895](https://github.com/owncloud/ocis/pull/2895)

   Tags: web

   We updated ownCloud Web to v5.0.0. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/ocis/pull/2895
   https://github.com/owncloud/ocis/pull/3157
   https://github.com/owncloud/web/releases/tag/v4.8.0
   https://github.com/owncloud/web/releases/tag/v5.0.0
# Changelog for [1.16.0] (2021-12-10)

The following sections list the changes for 1.16.0.

[1.16.0]: https://github.com/owncloud/ocis/compare/v1.15.0...v1.16.0

## Summary

* Bugfix - Fix claim selector based routing for basic auth: [#2779](https://github.com/owncloud/ocis/pull/2779)
* Bugfix - Disallow creation of a group with empty name via the OCS api: [#2825](https://github.com/owncloud/ocis/pull/2825)
* Bugfix - Fix using s3ng as the metadata storage backend: [#2807](https://github.com/owncloud/ocis/pull/2807)
* Bugfix - Use the CS3api up- and download workflow for the accounts service: [#2837](https://github.com/owncloud/ocis/pull/2837)
* Change - Rename `APP_PROVIDER_BASIC_*` environment variables: [#2812](https://github.com/owncloud/ocis/pull/2812)
* Change - Restructure Configuration Parsing: [#2708](https://github.com/owncloud/ocis/pull/2708)
* Change - OIDC: fallback if IDP doesn't provide "preferred_username" claim: [#2644](https://github.com/owncloud/ocis/issues/2644)
* Enhancement - Cleanup ocis-pkg config: [#2813](https://github.com/owncloud/ocis/pull/2813)
* Enhancement - Correct shutdown of services under runtime: [#2843](https://github.com/owncloud/ocis/pull/2843)
* Enhancement - Update REVA to v1.17.0: [#2849](https://github.com/owncloud/ocis/pull/2849)
* Enhancement - Update ownCloud Web to v4.6.1: [#2846](https://github.com/owncloud/ocis/pull/2846)

## Details

* Bugfix - Fix claim selector based routing for basic auth: [#2779](https://github.com/owncloud/ocis/pull/2779)

   We've fixed the claim selector based routing for requests using basic auth. Previously
   requests using basic auth have always been routed to the DefaultPolicy when using the claim
   selector despite the set cookie because the basic auth middleware fakes some OIDC claims.

   Now the cookie is checked before routing to the DefaultPolicy and therefore set cookie will
   also be respected for requests using basic auth.

   https://github.com/owncloud/ocis/pull/2779

* Bugfix - Disallow creation of a group with empty name via the OCS api: [#2825](https://github.com/owncloud/ocis/pull/2825)

   We've fixed the behavior for group creation on the OCS api, where it was possible to create a
   group with an empty name. This was is not possible on oC10 and is therefore also forbidden on oCIS
   to keep compatibility. This PR forbids the creation and also ensures the correct status code
   for both OCS v1 and OCS v2 apis.

   https://github.com/owncloud/ocis/issues/2823
   https://github.com/owncloud/ocis/pull/2825

* Bugfix - Fix using s3ng as the metadata storage backend: [#2807](https://github.com/owncloud/ocis/pull/2807)

   It is now possible to use s3ng as the metadata storage backend.

   https://github.com/owncloud/ocis/issues/2668
   https://github.com/owncloud/ocis/pull/2807

* Bugfix - Use the CS3api up- and download workflow for the accounts service: [#2837](https://github.com/owncloud/ocis/pull/2837)

   We've fixed the interaction of the accounts service with the metadata storage after bypassing
   the InitiateUpload and InitiateDownload have been removed from various storage drivers. The
   accounts service now uses the proper CS3apis workflow for up- and downloads.

   https://github.com/owncloud/ocis/pull/2837
   https://github.com/cs3org/reva/pull/2309

* Change - Rename `APP_PROVIDER_BASIC_*` environment variables: [#2812](https://github.com/owncloud/ocis/pull/2812)

   We've renamed the `APP_PROVIDER_BASIC_*` to `APP_PROVIDER_*` since the `_BASIC_` part is a
   copy and paste error. Now all app provider environment variables are consistently starting
   with `APP_PROVIDER_*`.

   https://github.com/owncloud/ocis/pull/2812
   https://github.com/owncloud/ocis/pull/2811

* Change - Restructure Configuration Parsing: [#2708](https://github.com/owncloud/ocis/pull/2708)

   Tags: ocis

   CLI flags are no longer needed for subcommands, as we rely solely on env variables and config
   files. This greatly simplifies configuration and deployment.

   https://github.com/owncloud/ocis/pull/2708

* Change - OIDC: fallback if IDP doesn't provide "preferred_username" claim: [#2644](https://github.com/owncloud/ocis/issues/2644)

   Some IDPs don't add the "preferred_username" claim. Fallback to the "email" claim in that case

   https://github.com/owncloud/ocis/issues/2644

* Enhancement - Cleanup ocis-pkg config: [#2813](https://github.com/owncloud/ocis/pull/2813)

   Certain values were of no use when configuring the ocis runtime.

   https://github.com/owncloud/ocis/pull/2813

* Enhancement - Correct shutdown of services under runtime: [#2843](https://github.com/owncloud/ocis/pull/2843)

   Supervised goroutines now shut themselves down on context cancellation propagation.

   https://github.com/owncloud/ocis/pull/2843

* Enhancement - Update REVA to v1.17.0: [#2849](https://github.com/owncloud/ocis/pull/2849)

   Updated REVA to v1.17.0 This update includes:

  * Fix [cs3org/reva#2305](https://github.com/cs3org/reva/pull/2305): Make sure /app/new takes `target` as absolute path
  * Fix [cs3org/reva#2303](https://github.com/cs3org/reva/pull/2303): Fix content disposition header for public links files
  * Fix [cs3org/reva#2316](https://github.com/cs3org/reva/pull/2316): Fix the share types in propfinds
  * Fix [cs3org/reva#2803](https://github.com/cs3org/reva/pull/2310): Fix app provider for editor public links
  * Fix [cs3org/reva#2298](https://github.com/cs3org/reva/pull/2298): Remove share refs from trashbin
  * Fix [cs3org/reva#2309](https://github.com/cs3org/reva/pull/2309): Remove early finish for zero byte file uploads
  * Fix [cs3org/reva#1941](https://github.com/cs3org/reva/pull/1941): Fix TUS uploads with transfer token only
  * Chg [cs3org/reva#2210](https://github.com/cs3org/reva/pull/2210): Fix app provider new file creation and improved error codes
  * Enh [cs3org/reva#2217](https://github.com/cs3org/reva/pull/2217): OIDC auth driver for ESCAPE IAM
  * Enh [cs3org/reva#2256](https://github.com/cs3org/reva/pull/2256): Return user type in the response of the ocs GET user call
  * Enh [cs3org/reva#2315](https://github.com/cs3org/reva/pull/2315): Add new attributes to public link propfinds
  * Enh [cs3org/reva#2740](https://github.com/cs3org/reva/pull/2250): Implement space membership endpoints
  * Enh [cs3org/reva#2252](https://github.com/cs3org/reva/pull/2252): Add the xattr sys.acl to SysACL (eosgrpc)
  * Enh [cs3org/reva#2314](https://github.com/cs3org/reva/pull/2314): OIDC: fallback if IDP doesn't provide "preferred_username" claim

   https://github.com/owncloud/ocis/pull/2849
   https://github.com/owncloud/ocis/pull/2835
   https://github.com/owncloud/ocis/pull/2837

* Enhancement - Update ownCloud Web to v4.6.1: [#2846](https://github.com/owncloud/ocis/pull/2846)

   Tags: web

   We updated ownCloud Web to v4.6.1. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/ocis/pull/2846
   https://github.com/owncloud/web/releases/tag/v4.6.1
# Changelog for [1.15.0] (2021-11-19)

The following sections list the changes for 1.15.0.

[1.15.0]: https://github.com/owncloud/ocis/compare/v1.14.0...v1.15.0

## Summary

* Bugfix - Don't allow empty password: [#197](https://github.com/owncloud/product/issues/197)
* Bugfix - Fix basic auth config: [#2719](https://github.com/owncloud/ocis/pull/2719)
* Bugfix - Fix basic auth with custom user claim: [#2755](https://github.com/owncloud/ocis/pull/2755)
* Bugfix - Fix oCIS startup ony systems with IPv6: [#2698](https://github.com/owncloud/ocis/pull/2698)
* Bugfix - Fix opening images in media viewer for some usernames: [#2738](https://github.com/owncloud/ocis/pull/2738)
* Bugfix - Fix error logging when there is no thumbnail for a file: [#2702](https://github.com/owncloud/ocis/pull/2702)
* Bugfix - Don't announce resharing via capabilities: [#2690](https://github.com/owncloud/ocis/pull/2690)
* Change - Make all insecure options configurable and change the default to false: [#2700](https://github.com/owncloud/ocis/issues/2700)
* Change - Update ownCloud Web to v4.5.0: [#2780](https://github.com/owncloud/ocis/pull/2780)
* Enhancement - Add API to list all spaces: [#2692](https://github.com/owncloud/ocis/pull/2692)
* Enhancement - Update REVA to v1.16.0: [#2737](https://github.com/owncloud/ocis/pull/2737)

## Details

* Bugfix - Don't allow empty password: [#197](https://github.com/owncloud/product/issues/197)

   It was allowed to create users with empty or spaces-only password. This is fixed

   https://github.com/owncloud/product/issues/197

* Bugfix - Fix basic auth config: [#2719](https://github.com/owncloud/ocis/pull/2719)

   Users could authenticate using basic auth even though `PROXY_ENABLE_BASIC_AUTH` was set to
   false.

   https://github.com/owncloud/ocis/issues/2466
   https://github.com/owncloud/ocis/pull/2719

* Bugfix - Fix basic auth with custom user claim: [#2755](https://github.com/owncloud/ocis/pull/2755)

   We've fixed authentication with basic if oCIS is configured to use a non-standard claim as user
   claim (`PROXY_USER_OIDC_CLAIM`). Prior to this bugfix the authentication always failed and
   is now working.

   https://github.com/owncloud/ocis/pull/2755

* Bugfix - Fix oCIS startup ony systems with IPv6: [#2698](https://github.com/owncloud/ocis/pull/2698)

   We've fixed failing startup of oCIS on systems with IPv6 addresses.

   https://github.com/owncloud/ocis/issues/2300
   https://github.com/owncloud/ocis/pull/2698

* Bugfix - Fix opening images in media viewer for some usernames: [#2738](https://github.com/owncloud/ocis/pull/2738)

   We've fixed the opening of images in the media viewer for user names containing special
   characters (eg. `@`) which will be URL-escaped. Before this fix users could not see the image in
   the media viewer. Now the user name is correctly escaped and the user can view the image in the
   media viewer.

   https://github.com/owncloud/ocis/pull/2738

* Bugfix - Fix error logging when there is no thumbnail for a file: [#2702](https://github.com/owncloud/ocis/pull/2702)

   We've fixed the behavior of the logging when there is no thumbnail for a file (because the
   filetype is not supported for thumbnail generation). Previously the WebDAV service always
   issues an error log in this case. Now, we don't log this event any more.

   https://github.com/owncloud/ocis/pull/2702

* Bugfix - Don't announce resharing via capabilities: [#2690](https://github.com/owncloud/ocis/pull/2690)

   OCIS / Reva is not capable of resharing, yet. We've set the resharing capability to false, so
   that clients have a chance to react accordingly.

   https://github.com/owncloud/ocis/pull/2690

* Change - Make all insecure options configurable and change the default to false: [#2700](https://github.com/owncloud/ocis/issues/2700)

   We had several hard-coded 'insecure' flags. These options are now configurable and default to
   false. Also we changed all other 'insecure' flags with a previous default of true to false.

   In development environments using self signed certs (the default) you now need to set these
   flags:

   ``` PROXY_OIDC_INSECURE=true STORAGE_FRONTEND_APPPROVIDER_INSECURE=true
   STORAGE_FRONTEND_ARCHIVER_INSECURE=true STORAGE_FRONTEND_OCDAV_INSECURE=true
   STORAGE_HOME_DATAPROVIDER_INSECURE=true
   STORAGE_METADATA_DATAPROVIDER_INSECURE=true STORAGE_OIDC_INSECURE=true
   STORAGE_USERS_DATAPROVIDER_INSECURE=true THUMBNAILS_CS3SOURCE_INSECURE=true
   THUMBNAILS_WEBDAVSOURCE_INSECURE=true ```

   As an alternative you also can set a single flag, which configures all options together:

   ``` OCIS_INSECURE=true ```

   https://github.com/owncloud/ocis/issues/2700
   https://github.com/owncloud/ocis/pull/2745

* Change - Update ownCloud Web to v4.5.0: [#2780](https://github.com/owncloud/ocis/pull/2780)

   Tags: web

   We updated ownCloud Web to v4.5.0. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/ocis/pull/2780
   https://github.com/owncloud/web/releases/tag/v4.5.0

* Enhancement - Add API to list all spaces: [#2692](https://github.com/owncloud/ocis/pull/2692)

   Added a graph endpoint to enable users with the `list-all-spaces` permission to list all
   spaces.

   https://github.com/owncloud/ocis/pull/2692

* Enhancement - Update REVA to v1.16.0: [#2737](https://github.com/owncloud/ocis/pull/2737)

   Updated REVA to v1.16.0 This update includes:

  * Fix [cs3org/reva#2245](https://github.com/cs3org/reva/pull/2245): Don't announce search-files capability
  * Fix [cs3org/reva#2247](https://github.com/cs3org/reva/pull/2247): Merge user ACLs from EOS to sys ACLs
  * Fix [cs3org/reva#2279](https://github.com/cs3org/reva/pull/2279): Return the inode of the version folder for files when listing in EOS
  * Fix [cs3org/reva#2294](https://github.com/cs3org/reva/pull/2294): Fix HTTP return code when path is invalid
  * Fix [cs3org/reva#2231](https://github.com/cs3org/reva/pull/2231): Fix share permission on a single file in sql share driver (cbox pkg)
  * Fix [cs3org/reva#2230](https://github.com/cs3org/reva/pull/2230): Fix open by default app and expose default app
  * Fix [cs3org/reva#2265](https://github.com/cs3org/reva/pull/2265): Fix nil pointer exception when resolving members of a group (rest driver)
  * Fix [cs3org/reva#1214](https://github.com/cs3org/reva/pull/1214): Fix restoring versions
  * Fix [cs3org/reva#2254](https://github.com/cs3org/reva/pull/2254): Fix spaces propfind
  * Fix [cs3org/reva#2260](https://github.com/cs3org/reva/pull/2260): Fix unset quota xattr on darwin
  * Fix [cs3org/reva#5776](https://github.com/cs3org/reva/pull/5776): Enforce permissions in public share apps
  * Fix [cs3org/reva#2767](https://github.com/cs3org/reva/pull/2767): Fix status code for WebDAV mkcol requests where an ancestor is missing
  * Fix [cs3org/reva#2287](https://github.com/cs3org/reva/pull/2287): Add public link access via mount-ID:token/relative-path to the scope
  * Fix [cs3org/reva#2244](https://github.com/cs3org/reva/pull/2244): Fix the permissions response for shared files in the cbox sql driver
  * Enh [cs3org/reva#2219](https://github.com/cs3org/reva/pull/2219): Add virtual view tests
  * Enh [cs3org/reva#2230](https://github.com/cs3org/reva/pull/2230): Add priority to app providers
  * Enh [cs3org/reva#2258](https://github.com/cs3org/reva/pull/2258): Improved error messages from the AppProviders
  * Enh [cs3org/reva#2119](https://github.com/cs3org/reva/pull/2119): Add authprovider owncloudsql
  * Enh [cs3org/reva#2211](https://github.com/cs3org/reva/pull/2211): Enhance the cbox share sql driver to store accepted group shares
  * Enh [cs3org/reva#2212](https://github.com/cs3org/reva/pull/2212): Filter root path according to the agent that makes the request
  * Enh [cs3org/reva#2237](https://github.com/cs3org/reva/pull/2237): Skip get user call in eosfs in case previous ones also failed
  * Enh [cs3org/reva#2266](https://github.com/cs3org/reva/pull/2266): Callback for the EOS UID cache to retry fetch for failed keys
  * Enh [cs3org/reva#2215](https://github.com/cs3org/reva/pull/2215): Aggregate resource info properties for virtual views
  * Enh [cs3org/reva#2271](https://github.com/cs3org/reva/pull/2271): Revamp the favorite manager and add the cbox sql driver
  * Enh [cs3org/reva#2248](https://github.com/cs3org/reva/pull/2248): Cache whether a user home was created or not
  * Enh [cs3org/reva#2282](https://github.com/cs3org/reva/pull/2282): Return a proper NOT_FOUND error when a user or group is not found
  * Enh [cs3org/reva#2268](https://github.com/cs3org/reva/pull/2268): Add the reverseproxy http service
  * Enh [cs3org/reva#2207](https://github.com/cs3org/reva/pull/2207): Enable users to list all spaces
  * Enh [cs3org/reva#2286](https://github.com/cs3org/reva/pull/2286): Add trace ID to middleware loggers
  * Enh [cs3org/reva#2251](https://github.com/cs3org/reva/pull/2251): Mentix service inference
  * Enh [cs3org/reva#2218](https://github.com/cs3org/reva/pull/2218): Allow filtering of mime types supported by app providers
  * Enh [cs3org/reva#2213](https://github.com/cs3org/reva/pull/2213): Add public link share type to propfind response
  * Enh [cs3org/reva#2253](https://github.com/cs3org/reva/pull/2253): Support the file editor role for public links
  * Enh [cs3org/reva#2208](https://github.com/cs3org/reva/pull/2208): Reduce redundant stat calls when statting by resource ID
  * Enh [cs3org/reva#2235](https://github.com/cs3org/reva/pull/2235): Specify a list of allowed folders/files to be archived
  * Enh [cs3org/reva#2267](https://github.com/cs3org/reva/pull/2267): Restrict the paths where share creation is allowed
  * Enh [cs3org/reva#2252](https://github.com/cs3org/reva/pull/2252): Add the xattr sys.acl to SysACL (eosgrpc)
  * Enh [cs3org/reva#2239](https://github.com/cs3org/reva/pull/2239): Update toml configs

   https://github.com/owncloud/ocis/pull/2737
   https://github.com/owncloud/ocis/pull/2726
   https://github.com/owncloud/ocis/pull/2790
   https://github.com/owncloud/ocis/pull/2797
# Changelog for [1.14.0] (2021-10-27)

The following sections list the changes for 1.14.0.

[1.14.0]: https://github.com/owncloud/ocis/compare/v1.13.0...v1.14.0

## Summary

* Security - Don't expose services by default: [#2612](https://github.com/owncloud/ocis/issues/2612)
* Bugfix - Create parent directories for idp configuration: [#2667](https://github.com/owncloud/ocis/issues/2667)
* Change - Configurable default quota: [#2621](https://github.com/owncloud/ocis/issues/2621)
* Change - New default data paths and easier configuration of the data path: [#2590](https://github.com/owncloud/ocis/pull/2590)
* Change - Split spaces webdav url and graph url in base and path: [#2660](https://github.com/owncloud/ocis/pull/2660)
* Change - Update ownCloud Web to v4.4.0: [#2681](https://github.com/owncloud/ocis/pull/2681)
* Enhancement - Add user setting capability: [#2655](https://github.com/owncloud/ocis/pull/2655)
* Enhancement - Broaden bufbuild/Buf usage: [#2630](https://github.com/owncloud/ocis/pull/2630)
* Enhancement - Replace fileb0x with go-embed: [#1199](https://github.com/owncloud/ocis/issues/1199)
* Enhancement - Upgrade to go-micro v4.1.0: [#2616](https://github.com/owncloud/ocis/pull/2616)
* Enhancement - Review and correct http header: [#2666](https://github.com/owncloud/ocis/pull/2666)
* Enhancement - Lower TUS max chunk size: [#2584](https://github.com/owncloud/ocis/pull/2584)
* Enhancement - Add sharees additional info parameter config to ocs: [#2637](https://github.com/owncloud/ocis/pull/2637)
* Enhancement - Add a middleware to authenticate public share requests: [#2536](https://github.com/owncloud/ocis/pull/2536)
* Enhancement - Report quota states: [#2628](https://github.com/owncloud/ocis/pull/2628)
* Enhancement - Start up a new machine auth provider in the storage service: [#2528](https://github.com/owncloud/ocis/pull/2528)
* Enhancement - Enforce permission on update space quota: [#2650](https://github.com/owncloud/ocis/pull/2650)
* Enhancement - Update lico to v0.51.1: [#2654](https://github.com/owncloud/ocis/pull/2654)
* Enhancement - Update reva to v1.15: [#2658](https://github.com/owncloud/ocis/pull/2658)

## Details

* Security - Don't expose services by default: [#2612](https://github.com/owncloud/ocis/issues/2612)

   We've changed the bind behaviour for all non public facing services. Before this PR all
   services would listen on all interfaces. After this PR, all services listen on 127.0.0.1 only,
   except the proxy which is listening on 0.0.0.0:9200.

   https://github.com/owncloud/ocis/issues/2612

* Bugfix - Create parent directories for idp configuration: [#2667](https://github.com/owncloud/ocis/issues/2667)

   The parent directories of the identifier-registration.yaml config file might not exist when
   starting idp. Create them, when that is the case.

   https://github.com/owncloud/ocis/issues/2667

* Change - Configurable default quota: [#2621](https://github.com/owncloud/ocis/issues/2621)

   When creating a new space a (configurable) default quota will be used (instead the hardcoded
   one). One can set the EnvVar `GRAPH_SPACES_DEFAULT_QUOTA` to configure it

   https://github.com/owncloud/ocis/issues/2621
   https://jira.owncloud.com/browse/OCIS-2070

* Change - New default data paths and easier configuration of the data path: [#2590](https://github.com/owncloud/ocis/pull/2590)

   We've changed the default data path for our release artifacts: - oCIS docker images will now
   store all data in `/var/lib/ocis` instead in `/var/tmp/ocis` - binary releases will now store
   all data in `~/.ocis` instead of `/var/tmp/ocis`

   Also if you're a developer and you run oCIS from source, it will store all data in `~/.ocis` from
   now on.

   You can now easily change the data path for all extensions by setting the environment variable
   `OCIS_BASE_DATA_PATH`.

   If you want to package oCIS, you also can set the default data path at compile time, eg. by passing
   `-X "github.com/owncloud/ocis/ocis-pkg/config/defaults.BaseDataPathType=path" -X
   "github.com/owncloud/ocis/ocis-pkg/config/defaults.BaseDataPathValue=/var/lib/ocis"`
   to your go build step.

   https://github.com/owncloud/ocis/pull/2590

* Change - Split spaces webdav url and graph url in base and path: [#2660](https://github.com/owncloud/ocis/pull/2660)

   We've fixed the behavior for the spaces webdav url and graph explorer graph url settings, so
   that they respect the environment variable `OCIS_URL`. Previously oCIS admins needed to set
   these URLs manually to make spaces and the graph explorer work.

   https://github.com/owncloud/ocis/issues/2659
   https://github.com/owncloud/ocis/pull/2660

* Change - Update ownCloud Web to v4.4.0: [#2681](https://github.com/owncloud/ocis/pull/2681)

   Tags: web

   We updated ownCloud Web to v4.4.0. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/ocis/pull/2681
   https://github.com/owncloud/web/releases/tag/v4.4.0

* Enhancement - Add user setting capability: [#2655](https://github.com/owncloud/ocis/pull/2655)

   We've added a capability to communicate the existence of a user settings service to clients.

   https://github.com/owncloud/web/issues/5926
   https://github.com/owncloud/ocis/pull/2655

* Enhancement - Broaden bufbuild/Buf usage: [#2630](https://github.com/owncloud/ocis/pull/2630)

   We've switched the usage of bufbuild/Buf from a protoc replacement only to also using it to
   configure the outputs and pinning dependencies.

   https://github.com/owncloud/ocis/pull/2630
   https://github.com/owncloud/ocis/pull/2616

* Enhancement - Replace fileb0x with go-embed: [#1199](https://github.com/owncloud/ocis/issues/1199)

   Go-embed already brings the functionality we need but with less code. We decided to use it
   instead of 3rd party fileb0x

   https://github.com/owncloud/ocis/issues/1199
   https://github.com/owncloud/ocis/pull/2631
   https://github.com/owncloud/ocis/pull/2649

* Enhancement - Upgrade to go-micro v4.1.0: [#2616](https://github.com/owncloud/ocis/pull/2616)

   We've upgraded to go-micro v4.1.0

   https://github.com/owncloud/ocis/pull/2616

* Enhancement - Review and correct http header: [#2666](https://github.com/owncloud/ocis/pull/2666)

   Reviewed and corrected the necessary http headers. Made CORS configurable.

   https://github.com/owncloud/ocis/pull/2666

* Enhancement - Lower TUS max chunk size: [#2584](https://github.com/owncloud/ocis/pull/2584)

   We've lowered the TUS max chunk size from infinite to 0.1GB so that chunking actually happens.

   https://github.com/owncloud/ocis/pull/2584
   https://github.com/cs3org/reva/pull/2136

* Enhancement - Add sharees additional info parameter config to ocs: [#2637](https://github.com/owncloud/ocis/pull/2637)

   https://github.com/owncloud/ocis/pull/2637

* Enhancement - Add a middleware to authenticate public share requests: [#2536](https://github.com/owncloud/ocis/pull/2536)

   Added a new middleware to authenticate public share requests. This makes it possible to use
   APIs which require an authenticated context with public shares.

   https://github.com/owncloud/ocis/issues/2479
   https://github.com/owncloud/ocis/pull/2536
   https://github.com/owncloud/ocis/pull/2652

* Enhancement - Report quota states: [#2628](https://github.com/owncloud/ocis/pull/2628)

   When listing the available spaces via the GraphAPI we now return quota states to make it easier
   for the clients to add visual indicators.

   https://github.com/owncloud/ocis/pull/2628

* Enhancement - Start up a new machine auth provider in the storage service: [#2528](https://github.com/owncloud/ocis/pull/2528)

   This PR also adds the config to skip encoding user groups in reva tokens

   https://github.com/owncloud/ocis/pull/2528
   https://github.com/owncloud/ocis/pull/2529

* Enhancement - Enforce permission on update space quota: [#2650](https://github.com/owncloud/ocis/pull/2650)

   Added a check that only users with the `set-space-quota` permission can update the space
   quota.

   https://github.com/owncloud/ocis/pull/2650

* Enhancement - Update lico to v0.51.1: [#2654](https://github.com/owncloud/ocis/pull/2654)

   Updated lico to v0.51.1 This update includes: * Apply LibreGraph naming treewide * move to
   go1.17 * Update 3rd party Go dependencies

   https://github.com/owncloud/ocis/pull/2654

* Enhancement - Update reva to v1.15: [#2658](https://github.com/owncloud/ocis/pull/2658)

   Updated reva to v1.15 This update includes:

  * Fix [cs3org/reva#2168](https://github.com/cs3org/reva/pull/2168): Override provider if was previously registered
  * Fix [cs3org/reva#2173](https://github.com/cs3org/reva/pull/2173): Fix archiver max size reached error
  * Fix [cs3org/reva#2167](https://github.com/cs3org/reva/pull/2167): Handle nil quota in decomposedfs
  * Fix [cs3org/reva#2153](https://github.com/cs3org/reva/pull/2153): Restrict EOS project spaces sharing permissions to admins and writers
  * Fix [cs3org/reva#2179](https://github.com/cs3org/reva/pull/2179): Fix the returned permissions for webdav uploads
  * Chg [cs3org/reva#2479](https://github.com/cs3org/reva/pull/2479): Make apps able to work with public shares
  * Enh [cs3org/reva#2174](https://github.com/cs3org/reva/pull/2174): Inherit ACLs for files from parent directories
  * Enh [cs3org/reva#2152](https://github.com/cs3org/reva/pull/2152): Add a reference parameter to the getQuota request
  * Enh [cs3org/reva#2171](https://github.com/cs3org/reva/pull/2171): Add optional claim parameter to machine auth
  * Enh [cs3org/reva#2135](https://github.com/cs3org/reva/pull/2135): Nextcloud test improvements
  * Enh [cs3org/reva#2180](https://github.com/cs3org/reva/pull/2180): Remove OCDAV options namespace parameter
  * Enh [cs3org/reva#2170](https://github.com/cs3org/reva/pull/2170): Handle propfind requests for existing files
  * Enh [cs3org/reva#2165](https://github.com/cs3org/reva/pull/2165): Allow access to recycle bin for arbitrary paths outside homes
  * Enh [cs3org/reva#2189](https://github.com/cs3org/reva/pull/2189): Add user settings capability
  * Enh [cs3org/reva#2162](https://github.com/cs3org/reva/pull/2162): Implement the UpdateStorageSpace method
  * Enh [cs3org/reva#2117](https://github.com/cs3org/reva/pull/2117): Add ocs cache warmup strategy for first request from the user

   https://github.com/owncloud/ocis/pull/2658
   https://github.com/owncloud/ocis/pull/2536
   https://github.com/owncloud/ocis/pull/2650
   https://github.com/owncloud/ocis/pull/2680
# Changelog for [1.13.0] (2021-10-13)

The following sections list the changes for 1.13.0.

[1.13.0]: https://github.com/owncloud/ocis/compare/v1.12.0...v1.13.0

## Summary

* Bugfix - Fix the account resolver middleware: [#2557](https://github.com/owncloud/ocis/pull/2557)
* Bugfix - Fix version information for extensions: [#2575](https://github.com/owncloud/ocis/pull/2575)
* Bugfix - Add the gatewaysvc to all shared configuration in REVA services: [#2597](https://github.com/owncloud/ocis/pull/2597)
* Bugfix - Use proper url path decode on the username: [#2511](https://github.com/owncloud/ocis/pull/2511)
* Bugfix - Remove notifications placeholder: [#2514](https://github.com/owncloud/ocis/pull/2514)
* Bugfix - Remove asset path configuration option from proxy: [#2576](https://github.com/owncloud/ocis/pull/2576)
* Bugfix - Race condition in config parsing: [#2574](https://github.com/owncloud/ocis/pull/2574)
* Change - Configure users and metadata storage separately: [#2598](https://github.com/owncloud/ocis/pull/2598)
* Change - Make the drives create method odata compliant: [#2531](https://github.com/owncloud/ocis/pull/2531)
* Change - Unify Envvar names configuring REVA gateway address: [#2587](https://github.com/owncloud/ocis/pull/2587)
* Change - Update ownCloud Web to v4.3.0: [#2589](https://github.com/owncloud/ocis/pull/2589)
* Enhancement - Updated MimeTypes configuration for AppRegistry: [#2603](https://github.com/owncloud/ocis/pull/2603)
* Enhancement - Add maximum files and size to archiver capabilities: [#2544](https://github.com/owncloud/ocis/pull/2544)
* Enhancement - Reduced repository size: [#2579](https://github.com/owncloud/ocis/pull/2579)
* Enhancement - Return the newly created space: [#2610](https://github.com/owncloud/ocis/pull/2610)
* Enhancement - Expose the reva archiver in OCIS: [#2509](https://github.com/owncloud/ocis/pull/2509)
* Enhancement - Favorites capability: [#2599](https://github.com/owncloud/ocis/pull/2599)
* Enhancement - Upgrade to GO 1.17: [#2605](https://github.com/owncloud/ocis/pull/2605)
* Enhancement - Make mimetype allow list configurable for app provider: [#2553](https://github.com/owncloud/ocis/pull/2553)
* Enhancement - Add allow_creation parameter to mime type config: [#2591](https://github.com/owncloud/ocis/pull/2591)
* Enhancement - Add option to skip generation of demo users and groups: [#2495](https://github.com/owncloud/ocis/pull/2495)
* Enhancement - Allow overriding the cookie based route by claim: [#2508](https://github.com/owncloud/ocis/pull/2508)
* Enhancement - Redirect invalid links to oC Web: [#2493](https://github.com/owncloud/ocis/pull/2493)
* Enhancement - Use reva's Authenticate method instead of spawning token managers: [#2528](https://github.com/owncloud/ocis/pull/2528)
* Enhancement - TLS config options for ldap in reva: [#2492](https://github.com/owncloud/ocis/pull/2492)
* Enhancement - Set reva JWT token expiration time to 24 hours by default: [#2527](https://github.com/owncloud/ocis/pull/2527)
* Enhancement - Update reva to v1.14.0: [#2615](https://github.com/owncloud/ocis/pull/2615)

## Details

* Bugfix - Fix the account resolver middleware: [#2557](https://github.com/owncloud/ocis/pull/2557)

   The accounts resolver middleware put an empty token into the request when the user was already
   present. Added a step to get the token for the user.

   https://github.com/owncloud/ocis/pull/2557

* Bugfix - Fix version information for extensions: [#2575](https://github.com/owncloud/ocis/pull/2575)

   We've fixed the behavior for `ocis version` which previously always showed `0.0.0` as version
   for extensions. Now the real version of the extensions are shown.

   https://github.com/owncloud/ocis/pull/2575

* Bugfix - Add the gatewaysvc to all shared configuration in REVA services: [#2597](https://github.com/owncloud/ocis/pull/2597)

   We've fixed the configuration for REVA services which didn't have a gatewaysvc in their shared
   configuration. This could lead to default gatewaysvc addresses in the auth middleware. Now it
   is set everywhere.

   https://github.com/owncloud/ocis/pull/2597

* Bugfix - Use proper url path decode on the username: [#2511](https://github.com/owncloud/ocis/pull/2511)

   We now properly decode the username when reading it from a url parameter

   https://github.com/owncloud/ocis/pull/2511

* Bugfix - Remove notifications placeholder: [#2514](https://github.com/owncloud/ocis/pull/2514)

   Since Reva was communicating its notification capabilities incorrectly, oCIS relied on a
   hardcoded string to overwrite them. This has been fixed in
   [reva#1819](https://github.com/cs3org/reva/pull/1819) so we now removed the hardcoded
   string and don't modify Reva's notification capabilities anymore in order to fix clients
   having to poll a (non-existent) notifications endpoint.

   https://github.com/owncloud/ocis/pull/2514

* Bugfix - Remove asset path configuration option from proxy: [#2576](https://github.com/owncloud/ocis/pull/2576)

   We've remove the asset path configuration option (`--asset-path` or `PROXY_ASSET_PATH`)
   since it didn't do anything at all.

   https://github.com/owncloud/ocis/pull/2576

* Bugfix - Race condition in config parsing: [#2574](https://github.com/owncloud/ocis/pull/2574)

   There was a race condition in the config parsing when configuring the storage services caused
   by services overwriting a pointer to a config value. We fixed it by setting sane defaults.

   https://github.com/owncloud/ocis/pull/2574

* Change - Configure users and metadata storage separately: [#2598](https://github.com/owncloud/ocis/pull/2598)

   We've fixed the configuration behaviour of the user and metadata service writing in the same
   directory when using oCIS storage.

   Therefore we needed to separate the configuration of the users and metadata storage so that
   they now can be configured totally separate.

   https://github.com/owncloud/ocis/pull/2598

* Change - Make the drives create method odata compliant: [#2531](https://github.com/owncloud/ocis/pull/2531)

   When creating a space on the graph API we now use the POST Body to provide the parameters.

   https://github.com/owncloud/ocis/pull/2531
   https://github.com/owncloud/ocis/pull/2535
   https://www.odata.org/getting-started/basic-tutorial/#modifyData

* Change - Unify Envvar names configuring REVA gateway address: [#2587](https://github.com/owncloud/ocis/pull/2587)

   We've renamed all envvars configuring REVA gateway address to `REVA_GATEWAY`, additionally
   we renamed the cli parameters to `--reva-gateway-addr` and adjusted the description

   https://github.com/owncloud/ocis/issues/2091
   https://github.com/owncloud/ocis/pull/2587

* Change - Update ownCloud Web to v4.3.0: [#2589](https://github.com/owncloud/ocis/pull/2589)

   Tags: web

   We updated ownCloud Web to v4.3.0. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/ocis/pull/2589
   https://github.com/owncloud/web/releases/tag/v4.3.0

* Enhancement - Updated MimeTypes configuration for AppRegistry: [#2603](https://github.com/owncloud/ocis/pull/2603)

   We updated the type of the mime types config to a list, to keep the order of mime types from the
   config.

   https://github.com/owncloud/ocis/pull/2603

* Enhancement - Add maximum files and size to archiver capabilities: [#2544](https://github.com/owncloud/ocis/pull/2544)

   We added the maximum files count and maximum archive size of the archiver to the capabilities
   endpoint. Clients can use this to generate warnings before the actual archive creation fails.

   https://github.com/owncloud/ocis/issues/2537
   https://github.com/owncloud/ocis/pull/2544
   https://github.com/cs3org/reva/pull/2105

* Enhancement - Reduced repository size: [#2579](https://github.com/owncloud/ocis/pull/2579)

   We removed leftover artifacts from the migration to a single repository.

   https://github.com/owncloud/ocis/pull/2579

* Enhancement - Return the newly created space: [#2610](https://github.com/owncloud/ocis/pull/2610)

   Changed the response of the CreateSpace method to include the newly created space.

   https://github.com/owncloud/ocis/pull/2610
   https://github.com/cs3org/reva/pull/2158

* Enhancement - Expose the reva archiver in OCIS: [#2509](https://github.com/owncloud/ocis/pull/2509)

   The reva archiver can now be accessed through the storage frontend service

   https://github.com/owncloud/ocis/pull/2509

* Enhancement - Favorites capability: [#2599](https://github.com/owncloud/ocis/pull/2599)

   We've added a capability for the storage frontend which can be used to announce to clients
   whether or not favorites are supported. By default this is disabled because the listing of
   favorites doesn't survive service restarts at the moment.

   https://github.com/owncloud/ocis/pull/2599

* Enhancement - Upgrade to GO 1.17: [#2605](https://github.com/owncloud/ocis/pull/2605)

   We've upgraded the used GO version from 1.16 to 1.17.

   https://github.com/owncloud/ocis/pull/2605

* Enhancement - Make mimetype allow list configurable for app provider: [#2553](https://github.com/owncloud/ocis/pull/2553)

   We've added a configuration option to configure the mimetype allow list introduced in
   cs3org/reva#2095. This also makes it possible to set one application per mime type as a
   default.

   https://github.com/owncloud/ocis/issues/2563
   https://github.com/owncloud/ocis/pull/2553
   https://github.com/cs3org/reva/pull/2095

* Enhancement - Add allow_creation parameter to mime type config: [#2591](https://github.com/owncloud/ocis/pull/2591)

   https://github.com/owncloud/ocis/pull/2591

* Enhancement - Add option to skip generation of demo users and groups: [#2495](https://github.com/owncloud/ocis/pull/2495)

   We've added a new environment variable to decide whether we should generate the demo users and
   groups or not. This environment variable is set to `true` by default, so the demo users and
   groups will get generated by default as long as oCIS is in its "technical preview" stage.

   In any case, there are still some users and groups automatically generated: for users: Reva
   IOP, Kopano IDP, admin; for groups: sysusers and users.

   https://github.com/owncloud/ocis/pull/2495

* Enhancement - Allow overriding the cookie based route by claim: [#2508](https://github.com/owncloud/ocis/pull/2508)

   When determining the routing policy we now let the claim override the cookie so that users are
   routed to the correct backend after login.

   https://github.com/owncloud/ocis/pull/2508

* Enhancement - Redirect invalid links to oC Web: [#2493](https://github.com/owncloud/ocis/pull/2493)

   Invalid links (eg. https://foo.bar/index.php/apps/pdfviewer) will be redirect to
   ownCloud Web instead of displaying a blank page with a "not found" message.

   https://github.com/owncloud/ocis/pull/2493
   https://github.com/owncloud/ocis/pull/2512

* Enhancement - Use reva's Authenticate method instead of spawning token managers: [#2528](https://github.com/owncloud/ocis/pull/2528)

   When using the CS3 proxy backend, we previously obtained the user from reva's userprovider
   service and minted the token ourselves. This required maintaining a shared JWT secret between
   ocis and reva, as well duplication of logic. This PR delegates this logic by using the
   `Authenticate` method provided by the reva gateway service to obtain this token, making it an
   arbitrary, indestructible entry. Currently, the changes have been made to the proxy service
   but will be extended to others as well.

   https://github.com/owncloud/ocis/pull/2528

* Enhancement - TLS config options for ldap in reva: [#2492](https://github.com/owncloud/ocis/pull/2492)

   We added the new config options "ldap-cacert" and "ldap-insecure" to the auth-, users- and
   groups-provider services to be able to do proper TLS configuration for the LDAP clients.
   "ldap-cacert" is by default configured to add the bundled glauth LDAP servers certificate to
   the trusted set for the LDAP clients. "ldap-insecure" is set to "false" by default and can be
   used to disable certificate checks (only advisable for development and test environments).

   https://github.com/owncloud/ocis/pull/2492

* Enhancement - Set reva JWT token expiration time to 24 hours by default: [#2527](https://github.com/owncloud/ocis/pull/2527)

   https://github.com/owncloud/ocis/pull/2527

* Enhancement - Update reva to v1.14.0: [#2615](https://github.com/owncloud/ocis/pull/2615)

   This update includes:

  * Bugfix [cs3org/reva#2103](https://github.com/cs3org/reva/pull/2103): AppProvider: propagate back errors reported by WOPI
  * Bugfix [cs3org/reva#2149](https://github.com/cs3org/reva/pull/2149): Remove excess info from the http list app providers endpoint
  * Bugfix [cs3org/reva#2114](https://github.com/cs3org/reva/pull/2114): Add as default app while registering and skip unset mimetypes
  * Bugfix [cs3org/reva#2095](https://github.com/cs3org/reva/pull/2095): Fix app open when multiple app providers are present
  * Bugfix [cs3org/reva#2135](https://github.com/cs3org/reva/pull/2135): Make TUS capabilities configurable
  * Bugfix [cs3org/reva#2076](https://github.com/cs3org/reva/pull/2076): Fix chi routing
  * Bugfix [cs3org/reva#2077](https://github.com/cs3org/reva/pull/2077): Fix concurrent registration of mimetypes
  * Bugfix [cs3org/reva#2154](https://github.com/cs3org/reva/pull/2154): Return OK when trying to delete a non existing reference
  * Bugfix [cs3org/reva#2078](https://github.com/cs3org/reva/pull/2078): Fix nil pointer exception in stat
  * Bugfix [cs3org/reva#2073](https://github.com/cs3org/reva/pull/2073): Fix opening a readonly filetype with WOPI
  * Bugfix [cs3org/reva#2140](https://github.com/cs3org/reva/pull/2140): Map GRPC error codes to REVA errors
  * Bugfix [cs3org/reva#2147](https://github.com/cs3org/reva/pull/2147): Follow up of #2138: this is the new expected format
  * Bugfix [cs3org/reva#2116](https://github.com/cs3org/reva/pull/2116): Differentiate share types when retrieving received shares in sql driver
  * Bugfix [cs3org/reva#2074](https://github.com/cs3org/reva/pull/2074): Fix Stat() for EOS storage provider
  * Bugfix [cs3org/reva#2151](https://github.com/cs3org/reva/pull/2151): Fix return code for webdav uploads when the token expired
  * Change [cs3org/reva#2121](https://github.com/cs3org/reva/pull/2121): Sharemanager API change
  * Enhancement [cs3org/reva#2090](https://github.com/cs3org/reva/pull/2090): Return space name during list storage spaces
  * Enhancement [cs3org/reva#2138](https://github.com/cs3org/reva/pull/2138): Default AppProvider on top of the providers list
  * Enhancement [cs3org/reva#2137](https://github.com/cs3org/reva/pull/2137): Revamp app registry and add parameter to control file creation
  * Enhancement [cs3org/reva#145](https://github.com/cs3org/reva/pull/2137): UI improvements for the AppProviders
  * Enhancement [cs3org/reva#2088](https://github.com/cs3org/reva/pull/2088): Add archiver and app provider to ocs capabilities
  * Enhancement [cs3org/reva#2537](https://github.com/cs3org/reva/pull/2537): Add maximum files and size to archiver capabilities
  * Enhancement [cs3org/reva#2100](https://github.com/cs3org/reva/pull/2100): Add support for resource id to the archiver
  * Enhancement [cs3org/reva#2158](https://github.com/cs3org/reva/pull/2158): Augment the Id of new spaces
  * Enhancement [cs3org/reva#2085](https://github.com/cs3org/reva/pull/2085): Make encoding user groups in access tokens configurable
  * Enhancement [cs3org/reva#146](https://github.com/cs3org/reva/pull/146): Filter the denial shares (permission = 0) out of
  * Enhancement [cs3org/reva#2141](https://github.com/cs3org/reva/pull/2141): Use golang v1.17
  * Enhancement [cs3org/reva#2053](https://github.com/cs3org/reva/pull/2053): Safer defaults for TLS verification on LDAP connections
  * Enhancement [cs3org/reva#2115](https://github.com/cs3org/reva/pull/2115): Reduce code duplication in LDAP related drivers
  * Enhancement [cs3org/reva#1989](https://github.com/cs3org/reva/pull/1989): Add redirects from OC10 URL formats
  * Enhancement [cs3org/reva#2479](https://github.com/cs3org/reva/pull/2479): Limit publicshare and resourceinfo scope content
  * Enhancement [cs3org/reva#2071](https://github.com/cs3org/reva/pull/2071): Implement listing favorites via the dav report API
  * Enhancement [cs3org/reva#2091](https://github.com/cs3org/reva/pull/2091): Nextcloud share managers
  * Enhancement [cs3org/reva#2070](https://github.com/cs3org/reva/pull/2070): More unit tests for the Nextcloud storage provider
  * Enhancement [cs3org/reva#2087](https://github.com/cs3org/reva/pull/2087): More unit tests for the Nextcloud auth and user managers
  * Enhancement [cs3org/reva#2075](https://github.com/cs3org/reva/pull/2075): Make owncloudsql leverage existing filecache index
  * Enhancement [cs3org/reva#2050](https://github.com/cs3org/reva/pull/2050): Add a share types filter to the OCS API
  * Enhancement [cs3org/reva#2134](https://github.com/cs3org/reva/pull/2134): Use space Type from request
  * Enhancement [cs3org/reva#2132](https://github.com/cs3org/reva/pull/2132): Align local tests with drone setup
  * Enhancement [cs3org/reva#2095](https://github.com/cs3org/reva/pull/2095): Whitelisting for apps
  * Enhancement [cs3org/reva#2155](https://github.com/cs3org/reva/pull/2155): Pass an extra query parameter to WOPI /openinapp with a

   https://github.com/owncloud/ocis/pull/2615
   https://github.com/owncloud/ocis/pull/2566
   https://github.com/owncloud/ocis/pull/2520
# Changelog for [1.12.0] (2021-09-14)

The following sections list the changes for 1.12.0.

[1.12.0]: https://github.com/owncloud/ocis/compare/v1.11.0...v1.12.0

## Summary

* Bugfix - Remove non working proxy route and fix cs3 users example: [#2474](https://github.com/owncloud/ocis/pull/2474)
* Bugfix - Set English as default language in the dropdown in the settings page: [#2465](https://github.com/owncloud/ocis/pull/2465)
* Change - Remove OnlyOffice extension: [#2433](https://github.com/owncloud/ocis/pull/2433)
* Change - Remove OnlyOffice extension: [#2433](https://github.com/owncloud/ocis/pull/2433)
* Change - Update ownCloud Web to v4.2.0: [#2501](https://github.com/owncloud/ocis/pull/2501)
* Enhancement - Add app provider and app provider registry: [#2204](https://github.com/owncloud/ocis/pull/2204)
* Enhancement - Add the create space permission: [#2461](https://github.com/owncloud/ocis/pull/2461)
* Enhancement - Add set space quota permission: [#2459](https://github.com/owncloud/ocis/pull/2459)
* Enhancement - Create a Space using the Graph API: [#2471](https://github.com/owncloud/ocis/pull/2471)
* Enhancement - Update go-chi/chi to version 5.0.3: [#2429](https://github.com/owncloud/ocis/pull/2429)
* Enhancement - Upgrade go micro to v3.6.0: [#2451](https://github.com/owncloud/ocis/pull/2451)
* Enhancement - Update reva to v1.13.0: [#2477](https://github.com/owncloud/ocis/pull/2477)

## Details

* Bugfix - Remove non working proxy route and fix cs3 users example: [#2474](https://github.com/owncloud/ocis/pull/2474)

   We removed a non working route from the proxy default configuration and fixed the cs3 users
   deployment example since it still used the accounts service. It now only uses the configured
   LDAP.

   https://github.com/owncloud/ocis/pull/2474

* Bugfix - Set English as default language in the dropdown in the settings page: [#2465](https://github.com/owncloud/ocis/pull/2465)

   The language dropdown didn't have a default language selected, and it was showing an empty
   value. Now it shows English instead.

   https://github.com/owncloud/ocis/pull/2465

* Change - Remove OnlyOffice extension: [#2433](https://github.com/owncloud/ocis/pull/2433)

   Tags: OnlyOffice

   We've removed the OnlyOffice extension in oCIS. OnlyOffice has their own web extension for
   OC10 backend now with [a dedicated
   guide](https://owncloud.dev/clients/web/deployments/oc10-app/#onlyoffice). In
   oCIS, we will follow up with a guide on how to start a WOPI server providing OnlyOffice soon.

   https://github.com/owncloud/ocis/pull/2433

* Change - Remove OnlyOffice extension: [#2433](https://github.com/owncloud/ocis/pull/2433)

   Tags: OnlyOffice

   We've removed the OnlyOffice extension in oCIS. OnlyOffice has their own web extension for
   OC10 backend now with [a dedicated
   guide](https://owncloud.dev/clients/web/deployments/oc10-app/#onlyoffice). In
   oCIS, we will follow up with a guide on how to start a WOPI server providing OnlyOffice soon.

   https://github.com/owncloud/ocis/pull/2433

* Change - Update ownCloud Web to v4.2.0: [#2501](https://github.com/owncloud/ocis/pull/2501)

   Tags: web

   We updated ownCloud Web to v4.2.0. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/ocis/pull/2501
   https://github.com/owncloud/web/releases/tag/v4.2.0

* Enhancement - Add app provider and app provider registry: [#2204](https://github.com/owncloud/ocis/pull/2204)

   We added the app provider and app provider registry. Now the CS3org WOPI server can be
   registered and OpenInApp requests can be done.

   https://github.com/owncloud/ocis/pull/2204
   https://github.com/cs3org/reva/pull/1785

* Enhancement - Add the create space permission: [#2461](https://github.com/owncloud/ocis/pull/2461)

   In preparation for the upcoming spaces features a `Create Space` permission was added.

   https://github.com/owncloud/ocis/pull/2461

* Enhancement - Add set space quota permission: [#2459](https://github.com/owncloud/ocis/pull/2459)

   In preparation for the upcoming spaces features a `SetSpaceQuota` permission was added.

   https://github.com/owncloud/ocis/pull/2459

* Enhancement - Create a Space using the Graph API: [#2471](https://github.com/owncloud/ocis/pull/2471)

   Spaces can now be created on `POST /drives/{drive-name}`. Only users with the `create-space`
   permissions can perform this operation.

   Allowed body form values are:

   - `quota` (bytes) maximum amount of bytes stored in the space. - `maxQuotaFiles` (integer)
   maximum amount of files supported by the space.

   https://github.com/owncloud/ocis/pull/2471

* Enhancement - Update go-chi/chi to version 5.0.3: [#2429](https://github.com/owncloud/ocis/pull/2429)

   Updated go-chi/chi to the latest release

   https://github.com/owncloud/ocis/pull/2429

* Enhancement - Upgrade go micro to v3.6.0: [#2451](https://github.com/owncloud/ocis/pull/2451)

   Go micro and all go micro plugins are now on v3.6.0

   https://github.com/owncloud/ocis/pull/2451

* Enhancement - Update reva to v1.13.0: [#2477](https://github.com/owncloud/ocis/pull/2477)

   This update includes:

  * Bugfix [cs3org/reva#2054](https://github.com/cs3org/reva/pull/2054): Fix the response after deleting a share
  * Bugfix [cs3org/reva#2026](https://github.com/cs3org/reva/pull/2026): Fix moving of a shared file
  * Bugfix [cs3org/reva#1605](https://github.com/cs3org/reva/pull/1605): Allow to expose full paths in OCS API
  * Bugfix [cs3org/reva#2033](https://github.com/cs3org/reva/pull/2033): Fix the storage id of shares
  * Bugfix [cs3org/reva#1991](https://github.com/cs3org/reva/pull/1991): Remove share references when declining shares
  * Enhancement [cs3org/reva#1994](https://github.com/cs3org/reva/pull/1994): Add owncloudsql driver for the userprovider
  * Enhancement [cs3org/reva#2065](https://github.com/cs3org/reva/pull/2065): New sharing role Manager
  * Enhancement [cs3org/reva#2015](https://github.com/cs3org/reva/pull/2015): Add spaces to the list of capabilities
  * Enhancement [cs3org/reva#2041](https://github.com/cs3org/reva/pull/2041): Create operations for Spaces
  * Enhancement [cs3org/reva#2029](https://github.com/cs3org/reva/pull/2029): Tracing agent configuration

   https://github.com/owncloud/ocis/pull/2477
# Changelog for [1.11.0] (2021-08-24)

The following sections list the changes for 1.11.0.

[1.11.0]: https://github.com/owncloud/ocis/compare/v1.10.0...v1.11.0

## Summary

* Bugfix - Specify primary user type for all accounts: [#2364](https://github.com/owncloud/ocis/pull/2364)
* Bugfix - Fix naming of the user- and groupprovider services: [#2388](https://github.com/owncloud/ocis/pull/2388)
* Change - Update ownCloud Web to v4.1.0: [#2426](https://github.com/owncloud/ocis/pull/2426)
* Enhancement - Use non root user for the owncloud/ocis docker image: [#2380](https://github.com/owncloud/ocis/pull/2380)
* Enhancement - Replace unmaintained jwt library: [#2386](https://github.com/owncloud/ocis/pull/2386)
* Enhancement - Update bleve to version 2.1.0: [#2391](https://github.com/owncloud/ocis/pull/2391)
* Enhancement - Update github.com/coreos/go-oidc to v3.0.0: [#2393](https://github.com/owncloud/ocis/pull/2393)
* Enhancement - Update reva to v1.12: [#2423](https://github.com/owncloud/ocis/pull/2423)

## Details

* Bugfix - Specify primary user type for all accounts: [#2364](https://github.com/owncloud/ocis/pull/2364)

   https://github.com/owncloud/ocis/pull/2364

* Bugfix - Fix naming of the user- and groupprovider services: [#2388](https://github.com/owncloud/ocis/pull/2388)

   The services are called "storage-userprovider" and "storage-groupprovider". The 'ocis
   help' output was misleading.

   https://github.com/owncloud/ocis/pull/2388

* Change - Update ownCloud Web to v4.1.0: [#2426](https://github.com/owncloud/ocis/pull/2426)

   Tags: web

   We updated ownCloud Web to v4.1.0. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/ocis/pull/2426
   https://github.com/owncloud/web/releases/tag/v4.1.0

* Enhancement - Use non root user for the owncloud/ocis docker image: [#2380](https://github.com/owncloud/ocis/pull/2380)

   The owncloud/ocis docker image now uses a non root user and enables you to set a different user
   with the docker `--user` parameter. The default user has the UID 1000 is part of a group with the
   GID 1000.

   This is a breaking change for existing docker deployments. The permission on the files and
   folders in persistent volumes need to be changed to the UID and GID used for oCIS (default
   1000:1000 if not changed by the user).

   https://github.com/owncloud/ocis/pull/2380

* Enhancement - Replace unmaintained jwt library: [#2386](https://github.com/owncloud/ocis/pull/2386)

   The old library [github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)
   is unmaintained and was replaced by the community maintained fork
   [github.com/golang-jwt/jwt](https://github.com/golang-jwt/jwt).

   https://github.com/owncloud/ocis/pull/2386

* Enhancement - Update bleve to version 2.1.0: [#2391](https://github.com/owncloud/ocis/pull/2391)

   Updated bleve to the current version.

   https://github.com/owncloud/ocis/pull/2391

* Enhancement - Update github.com/coreos/go-oidc to v3.0.0: [#2393](https://github.com/owncloud/ocis/pull/2393)

   Updated the github.com/coreos/go-oidc library to the version 3.0.0.

   https://github.com/owncloud/ocis/pull/2393

* Enhancement - Update reva to v1.12: [#2423](https://github.com/owncloud/ocis/pull/2423)

  * Enhancement cs3org/reva#1803: Introduce new webdav spaces endpoint
  * Bugfix cs3org/reva#1819: Disable notifications
  * Enhancement cs3org/reva#1861: Add support for runtime plugins
  * Bugfix cs3org/reva#1913: Logic to restore files to readonly nodes
  * Enhancement cs3org/reva#1946: Add share manager that connects to oc10 databases
  * Bugfix cs3org/reva#1954: Fix response format of the sharees API
  * Bugfix cs3org/reva#1956: Fix trashbin listing with depth 0
  * Bugfix cs3org/reva#1957: Fix etag propagation on deletes
  * Bugfix cs3org/reva#1960: Return the updated share after updating
  * Bugfix cs3org/reva#1965 cs3org/reva#1967: Fix the file target of user and group shares
  * Bugfix cs3org/reva#1980: Propagate the etag after restoring a file version
  * Enhancement cs3org/reva#1984: Replace OpenCensus with OpenTelemetry
  * Bugfix cs3org/reva#1985: Add quota stubs
  * Bugfix cs3org/reva#1987: Fix windows build
  * Bugfix cs3org/reva#1990: Increase oc10 compatibility of owncloudsql
  * Bugfix cs3org/reva#1992: Check if symlink exists instead of spamming the console
  * Bugfix cs3org/reva#1993: fix owncloudsql GetMD

   https://github.com/owncloud/ocis/pull/2423
# Changelog for [1.10.0] (2021-08-06)

The following sections list the changes for 1.10.0.

[1.10.0]: https://github.com/owncloud/ocis/compare/v1.9.0...v1.10.0

## Summary

* Bugfix - Improve IDP Login Accessibility: [#5376](https://github.com/owncloud/web/issues/5376)
* Bugfix - Forward basic auth to OpenID connect token authentication endpoint: [#2095](https://github.com/owncloud/ocis/issues/2095)
* Bugfix - Log all requests in the proxy access log: [#2301](https://github.com/owncloud/ocis/pull/2301)
* Bugfix - Update glauth to 20210729125545-b9aecdfcac31: [#2336](https://github.com/owncloud/ocis/pull/2336)
* Change - Update ownCloud Web to v4.0.0: [#2353](https://github.com/owncloud/ocis/pull/2353)
* Enhancement - Proxy: Add claims policy selector: [#2248](https://github.com/owncloud/ocis/pull/2248)
* Enhancement - Add ocs cache warmup config and warn on protobuf ns conflicts: [#2328](https://github.com/owncloud/ocis/pull/2328)
* Enhancement - Refactor graph API: [#2277](https://github.com/owncloud/ocis/pull/2277)
* Enhancement - Update REVA: [#2355](https://github.com/owncloud/ocis/pull/2355)
* Enhancement - Use only one go.mod file for project dependencies: [#2344](https://github.com/owncloud/ocis/pull/2344)

## Details

* Bugfix - Improve IDP Login Accessibility: [#5376](https://github.com/owncloud/web/issues/5376)

   We have addressed the feedback from the `a11y` audit and improved the IDP login screen
   accordingly.

   https://github.com/owncloud/web/issues/5376
   https://github.com/owncloud/web/issues/5377

* Bugfix - Forward basic auth to OpenID connect token authentication endpoint: [#2095](https://github.com/owncloud/ocis/issues/2095)

   When using `PROXY_ENABLE_BASIC_AUTH=true` we now forward request to the idp instead of
   trying to authenticate the request ourself.

   https://github.com/owncloud/ocis/issues/2095
   https://github.com/owncloud/ocis/issues/2094

* Bugfix - Log all requests in the proxy access log: [#2301](https://github.com/owncloud/ocis/pull/2301)

   We now use a dedicated middleware to log all requests, regardless of routing selector outcome.
   While the log now includes the remote address, the selected routing policy is only logged when
   log level is set to debug because the request context cannot be changed in the
   `directorSelectionDirector`, as per the `ReverseProxy.Director` documentation.

   https://github.com/owncloud/ocis/pull/2301

* Bugfix - Update glauth to 20210729125545-b9aecdfcac31: [#2336](https://github.com/owncloud/ocis/pull/2336)

  * Fixes the backend config not being passed correctly in ocis
  * Fixes a mutex being copied, leading to concurrent writes
  * Fixes UTF8 chars in filters
  * Fixes case insensitive strings

   https://github.com/owncloud/ocis/pull/2336
   https://github.com/glauth/glauth/pull/198
   https://github.com/glauth/glauth/pull/194

* Change - Update ownCloud Web to v4.0.0: [#2353](https://github.com/owncloud/ocis/pull/2353)

   Tags: web

   We updated ownCloud Web to v4.0.0. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/ocis/pull/2353
   https://github.com/owncloud/web/releases/tag/v4.0.0

* Enhancement - Proxy: Add claims policy selector: [#2248](https://github.com/owncloud/ocis/pull/2248)

   Using the proxy config file, it is now possible to let let the IdP determine the routing policy by
   sending an `ocis.routing.policy` claim. Its value will be used to determine the set of routes
   for the logged in user.

   https://github.com/owncloud/ocis/pull/2248

* Enhancement - Add ocs cache warmup config and warn on protobuf ns conflicts: [#2328](https://github.com/owncloud/ocis/pull/2328)

   https://github.com/owncloud/ocis/pull/2328

* Enhancement - Refactor graph API: [#2277](https://github.com/owncloud/ocis/pull/2277)

   We refactored the `/graph/v1.0/` endpoint which now relies on the internal access token fer
   authentication, getting rid of any LDAP or OIDC code to authenticate requests. This allows
   using the graph api when using basic auth or any other auth mechanism provided by the CS3 auth
   providers / reva gateway / ocis proxy.

   https://github.com/owncloud/ocis/pull/2277

* Enhancement - Update REVA: [#2355](https://github.com/owncloud/ocis/pull/2355)

   Update REVA from v1.10.1-0.20210730095301-fcb7a30a44a6 to
   v1.11.1-0.20210809134415-3fe79c870fb5 * Fix cs3org/reva#1978: Fix owner type is optional
   * Fix cs3org/reva#1965: fix value of file_target in shares * Fix cs3org/reva#1960: fix
   updating shares in the memory share manager * Fix cs3org/reva#1956: fix trashbin listing with
   depth 0 * Fix cs3org/reva#1957: fix etag propagation on deletes * Enh cs3org/reva#1861: [WIP]
   Runtime plugins * Fix cs3org/reva#1954: fix response format of the sharees API * Fix
   cs3org/reva#1819: Remove notifications key from ocs response * Enh cs3org/reva#1946: Add a
   share manager that connects to oc10 databases * Fix cs3org/reva#1899: Fix chunked uploads for
   new versions * Fix cs3org/reva#1906: Fix copy over existing resource * Fix cs3org/reva#1891:
   Delete Shared Resources as Receiver * Fix cs3org/reva#1907: Error when creating folder with
   existing name * Fix cs3org/reva#1937: Do not overwrite more specific matches when finding
   storage providers * Fix cs3org/reva#1939: Fix the share jail permissions in the decomposedfs
   * Fix cs3org/reva#1932: Numerous fixes to the owncloudsql storage driver * Fix
   cs3org/reva#1912: Fix response when listing versions of another user * Fix
   cs3org/reva#1910: Get user groups recursively in the cbox rest user driver * Fix
   cs3org/reva#1904: Set Content-Length to 0 when swallowing body in the datagateway * Fix
   cs3org/reva#1911: Fix version order in propfind responses * Fix cs3org/reva#1926: Trash Bin
   in oCIS Storage Operations * Fix cs3org/reva#1901: Fix response code when folder doesnt exist
   on upload * Enh cs3org/reva#1785: Extend app registry with AddProvider method and mimetype
   filters * Enh cs3org/reva#1938: Add methods to get and put context values * Enh
   cs3org/reva#1798: Add support for a deny-all permission on references * Enh
   cs3org/reva#1916: Generate updated protobuf bindings for EOS GRPC * Enh cs3org/reva#1887:
   Add "a" and "l" filter for grappa queries * Enh cs3org/reva#1919: Run gofmt before building *
   Enh cs3org/reva#1927: Implement RollbackToVersion for eosgrpc (needs a newer EOS MGM) * Enh
   cs3org/reva#1944: Implement listing supported mime types in app registry * Enh
   cs3org/reva#1870: Be defensive about wrongly quoted etags * Enh cs3org/reva#1940: Reduce
   memory usage when uploading with S3ng storage * Enh cs3org/reva#1888: Refactoring of the
   webdav code * Enh cs3org/reva#1900: Check for illegal names while uploading or moving files *
   Enh cs3org/reva#1925: Refactor listing and statting across providers for virtual views * Fix
   cs3org/reva#1883: Pass directories with trailing slashes to eosclient.GenerateToken * Fix
   cs3org/reva#1878: Improve the webdav error handling in the trashbin * Fix cs3org/reva#1884:
   Do not send body on failed range request * Enh cs3org/reva#1744: Add support for lightweight
   user types * Fix cs3org/reva#1904: Set Content-Length to 0 when swallowing body in the
   datagateway * Fix cs3org/reva#1899: Bugfix: Fix chunked uploads for new versions * Enh
   cs3org/reva#1888: Refactoring of the webdav code * Enh cs3org/reva#1887: Add "a" and "l"
   filter for grappa queries

   https://github.com/owncloud/ocis/pull/2355
   https://github.com/owncloud/ocis/pull/2295
   https://github.com/owncloud/ocis/pull/2314

* Enhancement - Use only one go.mod file for project dependencies: [#2344](https://github.com/owncloud/ocis/pull/2344)

   We now use one single go.mod file at the root of the repository rather than one per core
   extension.

   https://github.com/owncloud/ocis/pull/2344
# Changelog for [1.9.0] (2021-07-13)

The following sections list the changes for 1.9.0.

[1.9.0]: https://github.com/owncloud/ocis/compare/v1.8.0...v1.9.0

## Summary

* Bugfix - Panic when service fails to start: [#2252](https://github.com/owncloud/ocis/pull/2252)
* Bugfix - Dont use port 80 as debug for GroupsProvider: [#2271](https://github.com/owncloud/ocis/pull/2271)
* Change - Update ownCloud Web to v3.4.0: [#2276](https://github.com/owncloud/ocis/pull/2276)
* Change - Update WEB to v3.4.1: [#2283](https://github.com/owncloud/ocis/pull/2283)
* Enhancement - Runtime support for cherry picking extensions: [#2229](https://github.com/owncloud/ocis/pull/2229)
* Enhancement - Add readonly mode for storagehome and storageusers: [#2230](https://github.com/owncloud/ocis/pull/2230)
* Enhancement - Remove unnecessary Service.Init(): [#1705](https://github.com/owncloud/ocis/pull/1705)
* Enhancement - Update REVA to v1.9.1-0.20210628143859-9d29c36c0c3f: [#2227](https://github.com/owncloud/ocis/pull/2227)
* Enhancement - Update REVA to v1.9.1: [#2280](https://github.com/owncloud/ocis/pull/2280)

## Details

* Bugfix - Panic when service fails to start: [#2252](https://github.com/owncloud/ocis/pull/2252)

   Tags: runtime

   When attempting to run a service through the runtime that is currently running and fails to
   start, a race condition still redirect os Interrupt signals to a closed channel.

   https://github.com/owncloud/ocis/pull/2252

* Bugfix - Dont use port 80 as debug for GroupsProvider: [#2271](https://github.com/owncloud/ocis/pull/2271)

   A copy/paste error where the configuration for the groupsprovider's debug address was not
   present leaves go-micro to start the debug service in port 80 by default.

   https://github.com/owncloud/ocis/pull/2271

* Change - Update ownCloud Web to v3.4.0: [#2276](https://github.com/owncloud/ocis/pull/2276)

   Tags: web

   We updated ownCloud Web to v3.4.0. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/ocis/pull/2276
   https://github.com/owncloud/web/releases/tag/v3.4.0

* Change - Update WEB to v3.4.1: [#2283](https://github.com/owncloud/ocis/pull/2283)

  * Fix [5501](https://github.com/owncloud/web/pull/5501): loading previews in the right sidebar
  * Fix [5493](https://github.com/owncloud/web/pull/5493): view options position

   https://github.com/owncloud/ocis/pull/2283
   https://github.com/owncloud/web/releases/tag/v3.4.1

* Enhancement - Runtime support for cherry picking extensions: [#2229](https://github.com/owncloud/ocis/pull/2229)

   Support for running certain extensions supervised via cli flags. Example usage:

   ``` > ocis server --extensions="proxy, idp, storage-metadata, accounts" ```

   https://github.com/owncloud/ocis/pull/2229

* Enhancement - Add readonly mode for storagehome and storageusers: [#2230](https://github.com/owncloud/ocis/pull/2230)

   To enable the readonly mode use `STORAGE_HOME_READ_ONLY=true` and
   `STORAGE_USERS_READ_ONLY=true`. Alternative: use `OCIS_STORAGE_READ_ONLY=true`

   https://github.com/owncloud/ocis/pull/2230

* Enhancement - Remove unnecessary Service.Init(): [#1705](https://github.com/owncloud/ocis/pull/1705)

   As it turns out oCIS already calls this method. Invoking it twice would end in accidentally
   resetting values.

   https://github.com/owncloud/ocis/pull/1705

* Enhancement - Update REVA to v1.9.1-0.20210628143859-9d29c36c0c3f: [#2227](https://github.com/owncloud/ocis/pull/2227)

   https://github.com/owncloud/ocis/pull/2227

* Enhancement - Update REVA to v1.9.1: [#2280](https://github.com/owncloud/ocis/pull/2280)

  * Fix cs3org/reva#1843: Correct Dockerfile path for the reva CLI and alpine3.13 as builder
  * Fix cs3org/reva#1835: Cleanup owncloudsql driver
  * Fix cs3org/reva#1868: Minor fixes to the grpc/http plugin: checksum, url escaping
  * Fix cs3org/reva#1885: Fix template in eoshomewrapper to use context user rather than resource
  * Fix cs3org/reva#1833: Properly handle name collisions for deletes in the owncloud driver
  * Fix cs3org/reva#1874: Use the original file mtime during upload
  * Fix cs3org/reva#1854: Add the uid/gid to the url for eos
  * Fix cs3org/reva#1848: Fill in missing gid/uid number with nobody
  * Fix cs3org/reva#1831: Make the ocm-provider endpoint in the ocmd service unprotected
  * Fix cs3org/reva#1808: Use empty array in OCS Notifications endpoints
  * Fix cs3org/reva#1825: Raise max grpc message size
  * Fix cs3org/reva#1828: Send a proper XML header with error messages
  * Chg cs3org/reva#1828: Remove the oidc provider in order to upgrad mattn/go-sqlite3 to v1.14.7
  * Enh cs3org/reva#1834: Add API key to Mentix GOCDB connector
  * Enh cs3org/reva#1855: Minor optimization in parsing EOS ACLs
  * Enh cs3org/reva#1873: Update the EOS image tag to be for revad-eos image
  * Enh cs3org/reva#1802: Introduce list spaces
  * Enh cs3org/reva#1849: Add readonly interceptor
  * Enh cs3org/reva#1875: Simplify resource comparison
  * Enh cs3org/reva#1827: Support trashbin sub paths in the recycle API

   https://github.com/owncloud/ocis/pull/2280
# Changelog for [1.8.0] (2021-06-28)

The following sections list the changes for 1.8.0.

[1.8.0]: https://github.com/owncloud/ocis/compare/v1.7.0...v1.8.0

## Summary

* Bugfix - External storage registration used wrong config: [#2120](https://github.com/owncloud/ocis/pull/2120)
* Bugfix - Remove authentication from /status.php completely: [#2188](https://github.com/owncloud/ocis/pull/2188)
* Bugfix - Make webdav namespace configurable across services: [#2198](https://github.com/owncloud/ocis/pull/2198)
* Change - Update ownCloud Web to v3.3.0: [#2187](https://github.com/owncloud/ocis/pull/2187)
* Enhancement - Properly configure graph-explorer client registration: [#2118](https://github.com/owncloud/ocis/pull/2118)
* Enhancement - Use system default location to store TLS artefacts: [#2129](https://github.com/owncloud/ocis/pull/2129)
* Enhancement - Update REVA to v1.9: [#2205](https://github.com/owncloud/ocis/pull/2205)

## Details

* Bugfix - External storage registration used wrong config: [#2120](https://github.com/owncloud/ocis/pull/2120)

   The go-micro registry-singleton ignores the ocis configuration and defaults to mdns

   https://github.com/owncloud/ocis/pull/2120

* Bugfix - Remove authentication from /status.php completely: [#2188](https://github.com/owncloud/ocis/pull/2188)

   Despite requests without Authentication header being successful, requests with an invalid
   bearer token in the Authentication header were rejected in the proxy with an 401
   unauthenticated. Now the Authentication header is completely ignored for the /status.php
   route.

   https://github.com/owncloud/client/issues/8538
   https://github.com/owncloud/ocis/pull/2188

* Bugfix - Make webdav namespace configurable across services: [#2198](https://github.com/owncloud/ocis/pull/2198)

   The WebDAV namespace is used across various services, but it was previously hardcoded in some
   of the services. This PR uses the same environment variable to set the config correctly across
   the services.

   https://github.com/owncloud/ocis/pull/2198

* Change - Update ownCloud Web to v3.3.0: [#2187](https://github.com/owncloud/ocis/pull/2187)

   Tags: web

   We updated ownCloud Web to v3.3.0. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/ocis/pull/2187
   https://github.com/owncloud/web/releases/tag/v3.3.0

* Enhancement - Properly configure graph-explorer client registration: [#2118](https://github.com/owncloud/ocis/pull/2118)

   The client registration in the `identifier-registration.yaml` for the graph-explorer
   didn't contain `redirect_uris` nor `origins`. Both were added to prevent exploitation.

   https://github.com/owncloud/ocis/pull/2118

* Enhancement - Use system default location to store TLS artefacts: [#2129](https://github.com/owncloud/ocis/pull/2129)

   This used to default to the current location of the binary, which is not ideal after a first run as
   it leaves traces behind. It now uses the system's location for artefacts with the help of
   https://golang.org/pkg/os/#UserConfigDir.

   https://github.com/owncloud/ocis/pull/2129

* Enhancement - Update REVA to v1.9: [#2205](https://github.com/owncloud/ocis/pull/2205)

   This update includes * [set Content-Type
   correctly](https://github.com/cs3org/reva/pull/1750) * [Return file checksum
   available from the metadata for the EOS
   driver](https://github.com/cs3org/reva/pull/1755) * [Sort share entries
   alphabetically](https://github.com/cs3org/reva/pull/1772) * [Initial work on the
   owncloudsql driver](https://github.com/cs3org/reva/pull/1710) * [Add user ID cache
   warmup to EOS storage driver](https://github.com/cs3org/reva/pull/1774) * [Use
   UidNumber and GidNumber fields in User
   objects](https://github.com/cs3org/reva/pull/1573) * [EOS GRPC
   interface](https://github.com/cs3org/reva/pull/1471) * [switch
   references](https://github.com/cs3org/reva/pull/1721) * [remove user's uuid from
   trashbin file key](https://github.com/cs3org/reva/pull/1793) * [fix restore behavior of
   the trashbin API](https://github.com/cs3org/reva/pull/1795) * [eosfs: add arbitrary
   metadata support](https://github.com/cs3org/reva/pull/1811)

   https://github.com/owncloud/ocis/pull/2205
   https://github.com/owncloud/ocis/pull/2210
# Changelog for [1.7.0] (2021-06-04)

The following sections list the changes for 1.7.0.

[1.7.0]: https://github.com/owncloud/ocis/compare/v1.6.0...v1.7.0

## Summary

* Bugfix - Change the groups index to be case sensitive: [#2109](https://github.com/owncloud/ocis/pull/2109)
* Change - Update ownCloud Web to v3.2.0: [#2096](https://github.com/owncloud/ocis/pull/2096)
* Enhancement - Enable the s3ng storage driver: [#1886](https://github.com/owncloud/ocis/pull/1886)
* Enhancement - Color contrasts on IDP/OIDC login pages: [#2088](https://github.com/owncloud/ocis/pull/2088)
* Enhancement - Announce user profile picture capability: [#2036](https://github.com/owncloud/ocis/pull/2036)
* Enhancement - Update reva to v1.7.1-0.20210531093513-b74a2b156af6: [#2104](https://github.com/owncloud/ocis/pull/2104)

## Details

* Bugfix - Change the groups index to be case sensitive: [#2109](https://github.com/owncloud/ocis/pull/2109)

   Groups are considered to be case-sensitive. The index must handle them case-sensitive too
   otherwise we will have non-deterministic behavior while editing or deleting groups.

   https://github.com/owncloud/ocis/pull/2109

* Change - Update ownCloud Web to v3.2.0: [#2096](https://github.com/owncloud/ocis/pull/2096)

   Tags: web

   We updated ownCloud Web to v3.2.0. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/ocis/pull/2096
   https://github.com/owncloud/web/releases/tag/v3.2.0

* Enhancement - Enable the s3ng storage driver: [#1886](https://github.com/owncloud/ocis/pull/1886)

   We made it possible to use the new s3ng storage driver by adding according commandline flags and
   environment variables.

   https://github.com/owncloud/ocis/pull/1886

* Enhancement - Color contrasts on IDP/OIDC login pages: [#2088](https://github.com/owncloud/ocis/pull/2088)

   We have updated the color contrasts on the IDP pages in order to improve accessibility.

   https://github.com/owncloud/ocis/pull/2088

* Enhancement - Announce user profile picture capability: [#2036](https://github.com/owncloud/ocis/pull/2036)

   Added a new capability (through https://github.com/cs3org/reva/pull/1694) to prevent the
   web frontend from fetching (nonexistent) user avatar profile pictures which added latency &
   console errors.

   https://github.com/owncloud/ocis/pull/2036

* Enhancement - Update reva to v1.7.1-0.20210531093513-b74a2b156af6: [#2104](https://github.com/owncloud/ocis/pull/2104)

   This reva update includes: * [fix move in the owncloud storage
   driver](https://github.com/cs3org/reva/pull/1696) * [add checksum header to the tus
   preflight response](https://github.com/cs3org/reva/pull/1702) * [Add reliability
   calculations support to Mentix](https://github.com/cs3org/reva/pull/1649) * [fix
   response format when accepting shares](https://github.com/cs3org/reva/pull/1724) *
   [Datatx createtransfershare](https://github.com/cs3org/reva/pull/1725)

   https://github.com/owncloud/ocis/issues/2102
   https://github.com/owncloud/ocis/pull/2104
# Changelog for [1.6.0] (2021-05-12)

The following sections list the changes for 1.6.0.

[1.6.0]: https://github.com/owncloud/ocis/compare/v1.5.0...v1.6.0

## Summary

* Bugfix - Fix STORAGE_METADATA_ROOT default value override: [#1956](https://github.com/owncloud/ocis/pull/1956)
* Bugfix - Stop the supervisor if a service fails to start: [#1963](https://github.com/owncloud/ocis/pull/1963)
* Change - Update ownCloud Web to v3.1.0: [#2045](https://github.com/owncloud/ocis/pull/2045)
* Enhancement - Added dictionary files: [#2003](https://github.com/owncloud/ocis/pull/2003)
* Enhancement - Introduce login form with h1 tag for screen readers only: [#1991](https://github.com/owncloud/ocis/pull/1991)
* Enhancement - User Deprovisioning for the OCS API: [#1962](https://github.com/owncloud/ocis/pull/1962)
* Enhancement - Support thumbnails for txt files: [#1988](https://github.com/owncloud/ocis/pull/1988)
* Enhancement - Update reva to v1.7.1-0.20210430154404-69bd21f2cc97: [#2010](https://github.com/owncloud/ocis/pull/2010)
* Enhancement - Update reva to v1.7.1-0.20210507160327-e2c3841d0dbc: [#2044](https://github.com/owncloud/ocis/pull/2044)
* Enhancement - Use oc-select: [#1979](https://github.com/owncloud/ocis/pull/1979)
* Enhancement - Set SameSite settings to Strict for Web: [#2019](https://github.com/owncloud/ocis/pull/2019)

## Details

* Bugfix - Fix STORAGE_METADATA_ROOT default value override: [#1956](https://github.com/owncloud/ocis/pull/1956)

   The way the value was being set ensured that it was NOT being overridden where it should have
   been. This patch ensures the correct loading order of values.

   https://github.com/owncloud/ocis/pull/1956

* Bugfix - Stop the supervisor if a service fails to start: [#1963](https://github.com/owncloud/ocis/pull/1963)

   Steps to make the supervisor fail:

   `PROXY_HTTP_ADDR=0.0.0.0:9144 bin/ocis server`

   https://github.com/owncloud/ocis/pull/1963

* Change - Update ownCloud Web to v3.1.0: [#2045](https://github.com/owncloud/ocis/pull/2045)

   Tags: web

   We updated ownCloud Web to v3.1.0. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/ocis/pull/2045
   https://github.com/owncloud/web/releases/tag/v3.1.0

* Enhancement - Added dictionary files: [#2003](https://github.com/owncloud/ocis/pull/2003)

   Added the dictionary.js file for package settings and accounts which contains strings that
   should be synced to transifex but not exist in the UI directly.

   https://github.com/owncloud/ocis/pull/2003

* Enhancement - Introduce login form with h1 tag for screen readers only: [#1991](https://github.com/owncloud/ocis/pull/1991)

   https://github.com/owncloud/ocis/pull/1991

* Enhancement - User Deprovisioning for the OCS API: [#1962](https://github.com/owncloud/ocis/pull/1962)

   Use the CS3 API and Reva to deprovision users completely.

   Two new environment variables introduced: ``` OCS_IDM_ADDRESS OCS_STORAGE_USERS_DRIVER
   ```

   `OCS_IDM_ADDRESS` is also an alias for `OCIS_URL`; allows the OCS service to mint jwt tokens
   for the authenticated user that will be read by the reva authentication middleware.

   `OCS_STORAGE_USERS_DRIVER` determines how a user is deprovisioned. This kind of behavior is
   needed since every storage driver deals with deleting differently.

   https://github.com/owncloud/ocis/pull/1962

* Enhancement - Support thumbnails for txt files: [#1988](https://github.com/owncloud/ocis/pull/1988)

   Implemented support for thumbnails for txt files in the thumbnails service.

   https://github.com/owncloud/ocis/pull/1988

* Enhancement - Update reva to v1.7.1-0.20210430154404-69bd21f2cc97: [#2010](https://github.com/owncloud/ocis/pull/2010)

  * Fix recycle to different locations (https://github.com/cs3org/reva/pull/1541)
  * Fix user share as grantee in json backend (https://github.com/cs3org/reva/pull/1650)
  * Introduce named services (https://github.com/cs3org/reva/pull/1509)
  * Improve json marshalling of share protobuf messages (https://github.com/cs3org/reva/pull/1655)
  * Cache resources from share getter methods in OCS (https://github.com/cs3org/reva/pull/1643)
  * Fix public file shares (https://github.com/cs3org/reva/pull/1666)

   https://github.com/owncloud/ocis/pull/2010

* Enhancement - Update reva to v1.7.1-0.20210507160327-e2c3841d0dbc: [#2044](https://github.com/owncloud/ocis/pull/2044)

  * Add user profile picture to capabilities (https://github.com/cs3org/reva/pull/1694)
  * Mint scope-based access tokens for RBAC (https://github.com/cs3org/reva/pull/1669)
  * Add cache warmup strategy for OCS resource infos (https://github.com/cs3org/reva/pull/1664)
  * Filter shares based on type in OCS (https://github.com/cs3org/reva/pull/1683)

   https://github.com/owncloud/ocis/pull/2044

* Enhancement - Use oc-select: [#1979](https://github.com/owncloud/ocis/pull/1979)

   Replace oc-drop with oc select in settings

   https://github.com/owncloud/ocis/pull/1979

* Enhancement - Set SameSite settings to Strict for Web: [#2019](https://github.com/owncloud/ocis/pull/2019)

   Changed SameSite settings to Strict for Web to prevent warnings in Firefox

   https://github.com/owncloud/ocis/pull/2019
# Changelog for [1.5.0] (2021-04-21)

The following sections list the changes for 1.5.0.

[1.5.0]: https://github.com/owncloud/ocis/compare/v1.4.0...v1.5.0

## Summary

* Bugfix - Fixes "unaligned 64-bit atomic operation" panic on 32-bit ARM: [#1888](https://github.com/owncloud/ocis/pull/1888)
* Change - Make Protobuf package names unique: [#1875](https://github.com/owncloud/ocis/pull/1875)
* Change - Update ownCloud Web to v3.0.0: [#1938](https://github.com/owncloud/ocis/pull/1938)
* Enhancement - Change default path for thumbnails: [#1892](https://github.com/owncloud/ocis/pull/1892)
* Enhancement - Parse config on supervised mode with run subcommand: [#1931](https://github.com/owncloud/ocis/pull/1931)
* Enhancement - Update ODS in accounts & settings extension: [#1934](https://github.com/owncloud/ocis/pull/1934)
* Enhancement - Add config for public share SQL driver: [#1916](https://github.com/owncloud/ocis/pull/1916)
* Enhancement - Remove dead runtime code: [#1923](https://github.com/owncloud/ocis/pull/1923)
* Enhancement - Add option to reading registry rules from json file: [#1917](https://github.com/owncloud/ocis/pull/1917)
* Enhancement - Update reva to v1.6.1-0.20210414111318-a4b5148cbfb2: [#1872](https://github.com/owncloud/ocis/pull/1872)

## Details

* Bugfix - Fixes "unaligned 64-bit atomic operation" panic on 32-bit ARM: [#1888](https://github.com/owncloud/ocis/pull/1888)

   Sync/cache had uint64s that were not 64-bit aligned causing panics on 32-bit systems during
   atomic access

   https://github.com/owncloud/ocis/issues/1887
   https://github.com/owncloud/ocis/pull/1888

* Change - Make Protobuf package names unique: [#1875](https://github.com/owncloud/ocis/pull/1875)

   Introduce unique `package` and `go_package` names for our Protobuf definitions

   https://github.com/owncloud/ocis/pull/1875

* Change - Update ownCloud Web to v3.0.0: [#1938](https://github.com/owncloud/ocis/pull/1938)

   Tags: web

   We updated ownCloud Web to v3.0.0. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/ocis/pull/1938
   https://github.com/owncloud/web/releases/tag/v3.0.0

* Enhancement - Change default path for thumbnails: [#1892](https://github.com/owncloud/ocis/pull/1892)

   Changes the default path for thumbnails from `<os tmp dir>/ocis-thumbnails` to
   `/var/tmp/ocis/thumbnails`

   https://github.com/owncloud/ocis/issues/1891
   https://github.com/owncloud/ocis/pull/1892

* Enhancement - Parse config on supervised mode with run subcommand: [#1931](https://github.com/owncloud/ocis/pull/1931)

   Currently it is not possible to parse a single config file from an extension when running on
   supervised mode.

   https://github.com/owncloud/ocis/pull/1931

* Enhancement - Update ODS in accounts & settings extension: [#1934](https://github.com/owncloud/ocis/pull/1934)

   The accounts and settings extensions were updated to reflect the latest changes in the
   ownCloud design system. In addition, a couple of quick wins in terms of accessibility are
   included.

   https://github.com/owncloud/ocis/pull/1934

* Enhancement - Add config for public share SQL driver: [#1916](https://github.com/owncloud/ocis/pull/1916)

   https://github.com/owncloud/ocis/pull/1916

* Enhancement - Remove dead runtime code: [#1923](https://github.com/owncloud/ocis/pull/1923)

   When moving from the old runtime to the new one there were lots of files left behind that are
   essentially dead code and should be removed. The original code lives here
   github.com/refs/pman/ if someone finds it interesting to read.

   https://github.com/owncloud/ocis/pull/1923

* Enhancement - Add option to reading registry rules from json file: [#1917](https://github.com/owncloud/ocis/pull/1917)

   https://github.com/owncloud/ocis/pull/1917

* Enhancement - Update reva to v1.6.1-0.20210414111318-a4b5148cbfb2: [#1872](https://github.com/owncloud/ocis/pull/1872)

  * enforce quota (https://github.com/cs3org/reva/pull/1557)
  * Make additional info attribute configurable (https://github.com/cs3org/reva/pull/1588)
  * check ENOTDIR for readlink (https://github.com/cs3org/reva/pull/1597)
  * Add wrappers for EOS and EOS Home storage drivers (https://github.com/cs3org/reva/pull/1624)
  * eos: fixes for enabling file sharing (https://github.com/cs3org/reva/pull/1619)
  * implement checksums in the owncloud storage driver (https://github.com/cs3org/reva/pull/1629)

   https://github.com/owncloud/ocis/pull/1872
# Changelog for [1.4.0] (2021-03-30)

The following sections list the changes for 1.4.0.

[1.4.0]: https://github.com/owncloud/ocis/compare/v1.3.0...v1.4.0

## Summary

* Bugfix - Fix thumbnail generation for jpegs: [#1785](https://github.com/owncloud/ocis/pull/1785)
* Change - Update ownCloud Web to v2.1.0: [#1870](https://github.com/owncloud/ocis/pull/1870)
* Enhancement - Add focus to input elements on login page: [#1792](https://github.com/owncloud/ocis/pull/1792)
* Enhancement - Improve accessibility to input elements on login page: [#1794](https://github.com/owncloud/ocis/pull/1794)
* Enhancement - Add new build targets: [#1824](https://github.com/owncloud/ocis/pull/1824)
* Enhancement - Clarify expected failures: [#1790](https://github.com/owncloud/ocis/pull/1790)
* Enhancement - Replace special character in login page title with a regular minus: [#1813](https://github.com/owncloud/ocis/pull/1813)
* Enhancement - File Logging: [#1816](https://github.com/owncloud/ocis/pull/1816)
* Enhancement - Runtime Hostname and Port are now configurable: [#1822](https://github.com/owncloud/ocis/pull/1822)
* Enhancement - Generate thumbnails for .gif files: [#1791](https://github.com/owncloud/ocis/pull/1791)
* Enhancement - Tracing Refactor: [#1819](https://github.com/owncloud/ocis/pull/1819)
* Enhancement - Update reva to v1.6.1-0.20210326165326-e8a00d9b2368: [#1683](https://github.com/owncloud/ocis/pull/1683)

## Details

* Bugfix - Fix thumbnail generation for jpegs: [#1785](https://github.com/owncloud/ocis/pull/1785)

   Images with the extension `.jpeg` were not properly supported.

   https://github.com/owncloud/ocis/issues/1490
   https://github.com/owncloud/ocis/pull/1785

* Change - Update ownCloud Web to v2.1.0: [#1870](https://github.com/owncloud/ocis/pull/1870)

   Tags: web

   We updated ownCloud Web to v2.1.0. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/ocis/pull/1870
   https://github.com/owncloud/web/releases/tag/v2.1.0

* Enhancement - Add focus to input elements on login page: [#1792](https://github.com/owncloud/ocis/pull/1792)

   https://github.com/owncloud/web/issues/4322
   https://github.com/owncloud/ocis/pull/1792

* Enhancement - Improve accessibility to input elements on login page: [#1794](https://github.com/owncloud/ocis/pull/1794)

   https://github.com/owncloud/web/issues/4319
   https://github.com/owncloud/ocis/pull/1794
   https://github.com/owncloud/ocis/pull/1811

* Enhancement - Add new build targets: [#1824](https://github.com/owncloud/ocis/pull/1824)

   Make build target `build` used to build a binary twice, the second occurrence having symbols
   for debugging. We split this step in two and added `build-all` and `build-debug` targets.

   - `build-all` now behaves as the previous `build` target, it will generate 2 binaries, one for
   debug. - `build-debug` will build a single binary for debugging.

   https://github.com/owncloud/ocis/pull/1824

* Enhancement - Clarify expected failures: [#1790](https://github.com/owncloud/ocis/pull/1790)

   Some features, while covered by the ownCloud 10 acceptance tests, will not be implemented for
   now: - blacklisted / ignored files, because ocis does not need to blacklist `.htaccess` files -
   `OC-LazyOps` support was [removed from the
   clients](https://github.com/owncloud/client/pull/8398). We are thinking about [a state
   machine for uploads to properly solve that scenario and also list the state of files in progress
   in the web ui](https://github.com/owncloud/ocis/issues/214). The expected failures
   files now have a dedicated _Won't fix_ section for these items.

   https://github.com/owncloud/ocis/issues/214
   https://github.com/owncloud/ocis/pull/1790
   https://github.com/owncloud/client/pull/8398

* Enhancement - Replace special character in login page title with a regular minus: [#1813](https://github.com/owncloud/ocis/pull/1813)

   https://github.com/owncloud/ocis/pull/1813

* Enhancement - File Logging: [#1816](https://github.com/owncloud/ocis/pull/1816)

   When running supervised, support for configuring all logs to a single log file:
   `OCIS_LOG_FILE=/Users/foo/bar/ocis.log MICRO_REGISTRY=etcd bin/ocis server`

   Supports directing log from single extensions to a log file:
   `PROXY_LOG_FILE=/Users/foo/bar/proxy.log MICRO_REGISTRY=etcd bin/ocis proxy`

   https://github.com/owncloud/ocis/pull/1816

* Enhancement - Runtime Hostname and Port are now configurable: [#1822](https://github.com/owncloud/ocis/pull/1822)

   Without any configuration the ocis runtime will start on `localhost:9250` unless specified
   otherwise. Usage:

   - `OCIS_RUNTIME_PORT=6061 bin/ocis server` - overrides the oCIS runtime and starts on port
   6061 - `OCIS_RUNTIME_PORT=6061 bin/ocis list` - lists running extensions for the runtime on
   `localhost:6061`

   All subcommands are updated and expected to work with the following environment variables:

   ``` OCIS_RUNTIME_HOST OCIS_RUNTIME_PORT ```

   https://github.com/owncloud/ocis/pull/1822

* Enhancement - Generate thumbnails for .gif files: [#1791](https://github.com/owncloud/ocis/pull/1791)

   Added support for gifs to the thumbnails service.

   https://github.com/owncloud/ocis/pull/1791

* Enhancement - Tracing Refactor: [#1819](https://github.com/owncloud/ocis/pull/1819)

   Centralize tracing handling per extension.

   https://github.com/owncloud/ocis/pull/1819

* Enhancement - Update reva to v1.6.1-0.20210326165326-e8a00d9b2368: [#1683](https://github.com/owncloud/ocis/pull/1683)

  * quota querying and tree accounting [cs3org/reva#1405](https://github.com/cs3org/reva/pull/1405)
  * Fix webdav file versions endpoint bugs [cs3org/reva#1526](https://github.com/cs3org/reva/pull/1526)
  * Fix etag changing only once a second [cs3org/reva#1576](https://github.com/cs3org/reva/pull/1576)
  * Trashbin API parity [cs3org/reva#1552](https://github.com/cs3org/reva/pull/1552)
  * Signature authentication for public links [cs3org/reva#1590](https://github.com/cs3org/reva/pull/1590)

   https://github.com/owncloud/ocis/pull/1683
   https://github.com/cs3org/reva/pull/1405
   https://github.com/owncloud/ocis/pull/1861
# Changelog for [1.3.0] (2021-03-09)

The following sections list the changes for 1.3.0.

[1.3.0]: https://github.com/owncloud/ocis/compare/v1.2.0...v1.3.0

## Summary

* Bugfix - Purposely delay accounts service startup: [#1734](https://github.com/owncloud/ocis/pull/1734)
* Bugfix - Add missing gateway config: [#1716](https://github.com/owncloud/ocis/pull/1716)
* Bugfix - Fix accounts initialization: [#1696](https://github.com/owncloud/ocis/pull/1696)
* Bugfix - Fix the ttl of the authentication middleware cache: [#1699](https://github.com/owncloud/ocis/pull/1699)
* Change - Update ownCloud Web to v2.0.1: [#1683](https://github.com/owncloud/ocis/pull/1683)
* Change - Update ownCloud Web to v2.0.2: [#1776](https://github.com/owncloud/ocis/pull/1776)
* Enhancement - Remove the JWT from the log: [#1758](https://github.com/owncloud/ocis/pull/1758)
* Enhancement - Update go-micro to v3.5.1-0.20210217182006-0f0ace1a44a9: [#1670](https://github.com/owncloud/ocis/pull/1670)
* Enhancement - Update reva to v1.6.1-0.20210223065028-53f39499762e: [#1683](https://github.com/owncloud/ocis/pull/1683)
* Enhancement - Add initial nats and kubernetes registry support: [#1697](https://github.com/owncloud/ocis/pull/1697)

## Details

* Bugfix - Purposely delay accounts service startup: [#1734](https://github.com/owncloud/ocis/pull/1734)

   As it turns out the race condition between `accounts <-> storage-metadata` still remains.
   This PR is a hotfix, and it should be followed up with a proper fix. Either:

   - block the accounts' initialization until the storage metadata is ready (using the registry)
   or - allow the accounts service to initialize and use a message broker to signal the accounts the
   metadata storage is ready to receive requests.

   https://github.com/owncloud/ocis/pull/1734

* Bugfix - Add missing gateway config: [#1716](https://github.com/owncloud/ocis/pull/1716)

   The auth provider `ldap` and `oidc` drivers now need to be able talk to the reva gateway. We added
   the `gatewayscv` to the config that is passed to reva.

   https://github.com/owncloud/ocis/pull/1716

* Bugfix - Fix accounts initialization: [#1696](https://github.com/owncloud/ocis/pull/1696)

   Originally the accounts service relies on both the `settings` and `storage-metadata` to be up
   and running at the moment it starts. This is an antipattern as it will cause the entire service to
   panic if the dependants are not present.

   We inverted this dependency and moved the default initialization data (i.e: creating roles,
   permissions, settings bundles) and instead of notifying the settings service that the
   account has to provide with such options, the settings is instead initialized with the options
   the accounts rely on. Essentially saving bandwidth as there is no longer a gRPC call to the
   settings service.

   For the `storage-metadata` a retry mechanism was added that retries by default 20 times to
   fetch the `com.owncloud.storage.metadata` from the service registry every `500`
   milliseconds. If this retry expires the accounts panics, as its dependency on the
   `storage-metadata` service cannot be resolved.

   We also introduced a client wrapper that acts as middleware between a client and a server. For
   more information on how it works further read [here](https://github.com/sony/gobreaker)

   https://github.com/owncloud/ocis/pull/1696

* Bugfix - Fix the ttl of the authentication middleware cache: [#1699](https://github.com/owncloud/ocis/pull/1699)

   The authentication cache ttl was multiplied with `time.Second` multiple times. This
   resulted in a ttl that was not intended.

   https://github.com/owncloud/ocis/pull/1699

* Change - Update ownCloud Web to v2.0.1: [#1683](https://github.com/owncloud/ocis/pull/1683)

   Tags: web

   We updated ownCloud Web to v2.0.1. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/ocis/pull/1683
   https://github.com/owncloud/web/releases/tag/v2.0.1

* Change - Update ownCloud Web to v2.0.2: [#1776](https://github.com/owncloud/ocis/pull/1776)

   Tags: web

   We updated ownCloud Web to v2.0.2. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/ocis/pull/1776
   https://github.com/owncloud/web/releases/tag/v2.0.2

* Enhancement - Remove the JWT from the log: [#1758](https://github.com/owncloud/ocis/pull/1758)

   We were logging the JWT in some places. Secrets should not be exposed in logs so it got removed.

   https://github.com/owncloud/ocis/pull/1758

* Enhancement - Update go-micro to v3.5.1-0.20210217182006-0f0ace1a44a9: [#1670](https://github.com/owncloud/ocis/pull/1670)

   - We updated from go micro v2 (v2.9.1) go-micro v3 (v3.5.1 edge). - oCIS runtime is now aware of
   `MICRO_LOG_LEVEL` and is set to `error` by default. This decision was made because ownCloud,
   as framework builders, want to log everything oCIS related and hide everything unrelated by
   default. It can be re-enabled by setting it to a log level other than `error`. i.e:
   `MICRO_LOG_LEVEL=info`. - Updated `protoc-gen-micro` to the [latest
   version](https://github.com/asim/go-micro/tree/master/cmd/protoc-gen-micro). -
   We're using Prometheus wrappers from go-micro.

   https://github.com/owncloud/ocis/pull/1670
   https://github.com/asim/go-micro/pull/2126

* Enhancement - Update reva to v1.6.1-0.20210223065028-53f39499762e: [#1683](https://github.com/owncloud/ocis/pull/1683)

  * quota querying and tree accounting [cs3org/reva#1405](https://github.com/cs3org/reva/pull/1405)

   https://github.com/owncloud/ocis/pull/1683
   https://github.com/cs3org/reva/pull/1405

* Enhancement - Add initial nats and kubernetes registry support: [#1697](https://github.com/owncloud/ocis/pull/1697)

   We added initial support to use nats and kubernetes as a service registry using
   `MICRO_REGISTRY=nats` and `MICRO_REGISTRY=kubernetes` respectively. Multiple nodes can
   be given with `MICRO_REGISTRY_ADDRESS=1.2.3.4,5.6.7.8,9.10.11.12`.

   https://github.com/owncloud/ocis/pull/1697
# Changelog for [1.2.0] (2021-02-17)

The following sections list the changes for 1.2.0.

[1.2.0]: https://github.com/owncloud/ocis/compare/v1.1.0...v1.2.0

## Summary

* Bugfix - Check if roles are present in user object before looking those up: [#1388](https://github.com/owncloud/ocis/pull/1388)
* Bugfix - Fix etcd address configuration: [#1546](https://github.com/owncloud/ocis/pull/1546)
* Bugfix - Remove unimplemented config file option for oCIS root command: [#1636](https://github.com/owncloud/ocis/pull/1636)
* Bugfix - Fix thumbnail generation when using different idp: [#1624](https://github.com/owncloud/ocis/issues/1624)
* Change - Initial release of graph and graph explorer: [#1594](https://github.com/owncloud/ocis/pull/1594)
* Change - Move runtime code on refs/pman over to owncloud/ocis/ocis: [#1483](https://github.com/owncloud/ocis/pull/1483)
* Change - Update ownCloud Web to v2.0.0: [#1661](https://github.com/owncloud/ocis/pull/1661)
* Enhancement - Make use of new design-system oc-table: [#1597](https://github.com/owncloud/ocis/pull/1597)
* Enhancement - Use a default protocol parameter instead of explicitly disabling tus: [#1331](https://github.com/cs3org/reva/pull/1331)
* Enhancement - Functionality to map home directory to different storage providers: [#1186](https://github.com/owncloud/ocis/pull/1186)
* Enhancement - Introduce ADR: [#1042](https://github.com/owncloud/ocis/pull/1042)
* Enhancement - Switch to opencontainers annotation scheme: [#1381](https://github.com/owncloud/ocis/pull/1381)
* Enhancement - Migrate ocis-graph-explorer to ocis monorepo: [#1596](https://github.com/owncloud/ocis/pull/1596)
* Enhancement - Migrate ocis-graph to ocis monorepo: [#1594](https://github.com/owncloud/ocis/pull/1594)
* Enhancement - Enable group sharing and add config for sharing SQL driver: [#1626](https://github.com/owncloud/ocis/pull/1626)
* Enhancement - Update reva to v1.5.2-0.20210125114636-0c10b333ee69: [#1482](https://github.com/owncloud/ocis/pull/1482)

## Details

* Bugfix - Check if roles are present in user object before looking those up: [#1388](https://github.com/owncloud/ocis/pull/1388)

   https://github.com/owncloud/ocis/pull/1388

* Bugfix - Fix etcd address configuration: [#1546](https://github.com/owncloud/ocis/pull/1546)

   The etcd server address in `MICRO_REGISTRY_ADDRESS` was not picked up when etcd was set as
   service discovery registry `MICRO_REGISTRY=etcd`. Therefore etcd was only working if
   available on localhost / 127.0.0.1.

   https://github.com/owncloud/ocis/pull/1546

* Bugfix - Remove unimplemented config file option for oCIS root command: [#1636](https://github.com/owncloud/ocis/pull/1636)

   https://github.com/owncloud/ocis/pull/1636

* Bugfix - Fix thumbnail generation when using different idp: [#1624](https://github.com/owncloud/ocis/issues/1624)

   The thumbnail service was relying on a konnectd specific field in the access token. This logic
   was now replaced by a service parameter for the username.

   https://github.com/owncloud/ocis/issues/1624
   https://github.com/owncloud/ocis/pull/1628

* Change - Initial release of graph and graph explorer: [#1594](https://github.com/owncloud/ocis/pull/1594)

   Tags: graph, graph-explorer

   We brought initial basic Graph and Graph-Explorer support for the ownCloud Infinite Scale
   project.

   https://github.com/owncloud/ocis/pull/1594
   https://github.com/owncloud/ocis-graph-explorer/pull/3

* Change - Move runtime code on refs/pman over to owncloud/ocis/ocis: [#1483](https://github.com/owncloud/ocis/pull/1483)

   Tags: ocis, runtime

   Currently, the runtime is under the private account of an oCIS developer. For future-proofing
   we don't want oCIS mission critical components to depend on external repositories, so we're
   including refs/pman module as an oCIS package instead.

   https://github.com/owncloud/ocis/pull/1483

* Change - Update ownCloud Web to v2.0.0: [#1661](https://github.com/owncloud/ocis/pull/1661)

   Tags: web

   We updated ownCloud Web to v2.0.0. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/ocis/pull/1661
   https://github.com/owncloud/web/releases/tag/v2.0.0

* Enhancement - Make use of new design-system oc-table: [#1597](https://github.com/owncloud/ocis/pull/1597)

   Tags: ui, accounts

   The design-system table component has changed the way it's used. We updated accounts-ui to use
   the new 'oc-table-simple' component.

   https://github.com/owncloud/ocis/pull/1597

* Enhancement - Use a default protocol parameter instead of explicitly disabling tus: [#1331](https://github.com/cs3org/reva/pull/1331)

   https://github.com/cs3org/reva/pull/1331
   https://github.com/owncloud/ocis/pull/1374

* Enhancement - Functionality to map home directory to different storage providers: [#1186](https://github.com/owncloud/ocis/pull/1186)

   We added a parameter in reva that allows us to redirect /home requests to different storage
   providers based on a mapping derived from the user attributes, which was previously not
   possible since we hardcode the /home path for all users. For example, having its value as
   `/home/{{substr 0 1 .Username}}` can be used to redirect home requests for different users to
   different storage providers.

   https://github.com/owncloud/ocis/pull/1186
   https://github.com/cs3org/reva/pull/1142

* Enhancement - Introduce ADR: [#1042](https://github.com/owncloud/ocis/pull/1042)

   We will keep track of [Architectural Decision Records using
   Markdown](https://adr.github.io/madr/) in `/docs/adr`.

   https://github.com/owncloud/ocis/pull/1042

* Enhancement - Switch to opencontainers annotation scheme: [#1381](https://github.com/owncloud/ocis/pull/1381)

   Switch docker image annotation scheme to org.opencontainers standard because
   org.label-schema is depreciated.

   https://github.com/owncloud/ocis/pull/1381

* Enhancement - Migrate ocis-graph-explorer to ocis monorepo: [#1596](https://github.com/owncloud/ocis/pull/1596)

   Tags: ocis, ocis-graph-explorer

   Ocis-graph-explorer was not migrated during the monorepo conversion.

   https://github.com/owncloud/ocis/pull/1596

* Enhancement - Migrate ocis-graph to ocis monorepo: [#1594](https://github.com/owncloud/ocis/pull/1594)

   Tags: ocis, ocis-graph

   Ocis-graph was not migrated during the monorepo conversion.

   https://github.com/owncloud/ocis/pull/1594

* Enhancement - Enable group sharing and add config for sharing SQL driver: [#1626](https://github.com/owncloud/ocis/pull/1626)

   This PR adds config to support sharing with groups. It also introduces a breaking change for the
   CS3APIs definitions since grantees can now refer to both users as well as groups. Since we store
   the grantee information in a json file, `/var/tmp/ocis/storage/shares.json`, its previous
   version needs to be removed as we won't be able to unmarshal data corresponding to the previous
   definitions.

   https://github.com/owncloud/ocis/pull/1626
   https://github.com/cs3org/reva/pull/1453

* Enhancement - Update reva to v1.5.2-0.20210125114636-0c10b333ee69: [#1482](https://github.com/owncloud/ocis/pull/1482)

  * initial checksum support for ocis [cs3org/reva#1400](https://github.com/cs3org/reva/pull/1400)
  * Use updated etag of home directory even if it is cached [cs3org/reva#1416](https://github.com/cs3org/reva/pull/#1416)
  * Indicate in EOS containers that TUS is not supported [cs3org/reva#1415](https://github.com/cs3org/reva/pull/#1415)
  * Get status code from recycle response [cs3org/reva#1408](https://github.com/cs3org/reva/pull/#1408)

   https://github.com/owncloud/ocis/pull/1482
   https://github.com/cs3org/reva/pull/1400
   https://github.com/cs3org/reva/pull/1416
   https://github.com/cs3org/reva/pull/1415
   https://github.com/cs3org/reva/pull/1408
# Changelog for [1.1.0] (2021-01-22)

The following sections list the changes for 1.1.0.

[1.1.0]: https://github.com/owncloud/ocis/compare/v1.0.0...v1.1.0

## Summary

* Change - Disable pretty logging by default: [#1133](https://github.com/owncloud/ocis/pull/1133)
* Change - Add "volume" declaration to docker images: [#1375](https://github.com/owncloud/ocis/pull/1375)
* Change - Add "expose" information to docker images: [#1366](https://github.com/owncloud/ocis/pull/1366)
* Change - Generate cryptographically secure state token: [#1203](https://github.com/owncloud/ocis/pull/1203)
* Change - Move k6 to cdperf: [#1358](https://github.com/owncloud/ocis/pull/1358)
* Change - Update go version: [#1364](https://github.com/owncloud/ocis/pull/1364)
* Change - Update ownCloud Web to v1.0.1: [#1191](https://github.com/owncloud/ocis/pull/1191)
* Enhancement - Add OCIS_URL env var: [#1148](https://github.com/owncloud/ocis/pull/1148)
* Enhancement - Use sync.cache for roles cache: [#1367](https://github.com/owncloud/ocis/pull/1367)
* Enhancement - Add named locks and refactor cache: [#1212](https://github.com/owncloud/ocis/pull/1212)
* Enhancement - Update reva to v1.5.1: [#1372](https://github.com/owncloud/ocis/pull/1372)
* Enhancement - Update reva to v1.4.1-0.20210111080247-f2b63bfd6825: [#1194](https://github.com/owncloud/ocis/pull/1194)

## Details

* Change - Disable pretty logging by default: [#1133](https://github.com/owncloud/ocis/pull/1133)

   Tags: ocis

   Disable pretty logging default for performance reasons.

   https://github.com/owncloud/ocis/pull/1133

* Change - Add "volume" declaration to docker images: [#1375](https://github.com/owncloud/ocis/pull/1375)

   Tags: docker

   Add "volume" declaration to docker images. This makes it easier for Docker users to see where
   oCIS stores data.

   https://github.com/owncloud/ocis/pull/1375

* Change - Add "expose" information to docker images: [#1366](https://github.com/owncloud/ocis/pull/1366)

   Tags: docker

   Add "expose" information to docker images. Docker users will now see that we offer services on
   port 9200.

   https://github.com/owncloud/ocis/pull/1366

* Change - Generate cryptographically secure state token: [#1203](https://github.com/owncloud/ocis/pull/1203)

   Replaced Math.random with a cryptographically secure way to generate the oidc state token
   using the javascript crypto api.

   https://github.com/owncloud/ocis/pull/1203
   https://developer.mozilla.org/en-US/docs/Web/API/Crypto/getRandomValues
   https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math/random

* Change - Move k6 to cdperf: [#1358](https://github.com/owncloud/ocis/pull/1358)

   Tags: performance, testing, k6

   The ownCloud performance tests can not only be used to test oCIS. This is why we have decided to
   move the k6 tests to https://github.com/owncloud/cdperf

   https://github.com/owncloud/ocis/pull/1358

* Change - Update go version: [#1364](https://github.com/owncloud/ocis/pull/1364)

   Tags: go

   Update go from 1.13 to 1.15

   https://github.com/owncloud/ocis/pull/1364

* Change - Update ownCloud Web to v1.0.1: [#1191](https://github.com/owncloud/ocis/pull/1191)

   Tags: web

   We updated ownCloud Web to v1.0.1. Please refer to the changelog (linked) for details on the web
   release.

   https://github.com/owncloud/ocis/pull/1191
   https://github.com/owncloud/web/releases/tag/v1.0.1

* Enhancement - Add OCIS_URL env var: [#1148](https://github.com/owncloud/ocis/pull/1148)

   Tags: ocis

   We introduced a new environment variable `OCIS_URL` that expects a URL including protocol,
   host and optionally port to simplify configuring all the different services. These existing
   environment variables still take precedence, but will also fall back to `OCIS_URL`:
   `STORAGE_LDAP_IDP`, `STORAGE_OIDC_ISSUER`, `PROXY_OIDC_ISSUER`,
   `STORAGE_FRONTEND_PUBLIC_URL`, `KONNECTD_ISS`, `WEB_OIDC_AUTHORITY`, and
   `WEB_UI_CONFIG_SERVER`.

   Some environment variables are now built dynamically if they are not set: -
   `STORAGE_DATAGATEWAY_PUBLIC_URL` defaults to `<STORAGE_FRONTEND_PUBLIC_URL>/data`,
   also falling back to `OCIS_URL` - `WEB_OIDC_METADATA_URL` defaults to
   `<WEB_OIDC_AUTHORITY>/.well-known/openid-configuration`, also falling back to
   `OCIS_URL`

   Furthermore, the built in konnectd will generate an `identifier-registration.yaml` that
   uses the `KONNECTD_ISS` in the allowed `redirect_uris` and `origins`. It simplifies the
   default `https://localhost:9200` and remote deployment with `OCIS_URL` which is evaluated
   as a fallback if `KONNECTD_ISS` is not set.

   An oCIS server can now be started on a remote machine as easy as
   `OCIS_URL=https://cloud.ocis.test PROXY_HTTP_ADDR=0.0.0.0:443 ocis server`.

   Note that the `OCIS_DOMAIN` environment variable is not used by oCIS, but by the docker
   containers.

   https://github.com/owncloud/ocis/pull/1148

* Enhancement - Use sync.cache for roles cache: [#1367](https://github.com/owncloud/ocis/pull/1367)

   Tags: ocis-pkg

   Update ocis-pkg/roles cache to use ocis-pkg/sync cache

   https://github.com/owncloud/ocis/pull/1367

* Enhancement - Add named locks and refactor cache: [#1212](https://github.com/owncloud/ocis/pull/1212)

   Tags: ocis-pkg, accounts

   We had the case that we needed kind of a named locking mechanism which enables us to lock only
   under certain conditions. It's used in the indexer package where we do not need to lock
   everything, instead just lock the requested parts and differentiate between reads and
   writes.

   This made it possible to entirely remove locks from the accounts service and move them to the
   ocis-pkg indexer. Another part of this refactor was to make the cache atomic and write tests for
   it.

   - remove locking from accounts service - add sync package with named mutex - add named locking to
   indexer - move cache to sync package

   https://github.com/owncloud/ocis/issues/966
   https://github.com/owncloud/ocis/pull/1212

* Enhancement - Update reva to v1.5.1: [#1372](https://github.com/owncloud/ocis/pull/1372)

   Summary -------

  * Fix #1401: Use the user in request for deciding the layout for non-home DAV requests
  * Fix #1413: Re-include the '.git' dir in the Docker images to pass the version tag
  * Fix #1399: Fix ocis trash-bin purge
  * Enh #1397: Bump the Copyright date to 2021
  * Enh #1398: Support site authorization status in Mentix
  * Enh #1393: Allow setting favorites, mtime and a temporary etag
  * Enh #1403: Support remote cloud gathering metrics

   Details -------

  * Bugfix #1401: Use the user in request for deciding the layout for non-home DAV requests

   For the incoming /dav/files/userID requests, we have different namespaces depending on
   whether the request is for the logged-in user's namespace or not. Since in the storage drivers,
   we specify the layout depending only on the user whose resources are to be accessed, this fails
   when a user wants to access another user's namespace when the storage provider depends on the
   logged in user's namespace. This PR fixes that.

   For example, consider the following case. The owncloud fs uses a layout {{substr 0 1
   .Id.OpaqueId}}/{{.Id.OpaqueId}}. The user einstein sends a request to access a resource
   shared with him, say /dav/files/marie/abcd, which should be allowed. However, based on the
   way we applied the layout, there's no way in which this can be translated to /m/marie/.

   Https://github.com/cs3org/reva/pull/1401

  * Bugfix #1413: Re-include the '.git' dir in the Docker images to pass the version tag

   And git SHA to the release tool.

   Https://github.com/cs3org/reva/pull/1413

  * Bugfix #1399: Fix ocis trash-bin purge

   Fixes the empty trash-bin functionality for ocis-storage

   Https://github.com/owncloud/product/issues/254
   https://github.com/cs3org/reva/pull/1399

  * Enhancement #1397: Bump the Copyright date to 2021

   Https://github.com/cs3org/reva/pull/1397

  * Enhancement #1398: Support site authorization status in Mentix

   This enhancement adds support for a site authorization status to Mentix. This way, sites
   registered via a web app can now be excluded until authorized manually by an administrator.

   Furthermore, Mentix now sets the scheme for Prometheus targets. This allows us to also support
   monitoring of sites that do not support the default HTTPS scheme.

   Https://github.com/cs3org/reva/pull/1398

  * Enhancement #1393: Allow setting favorites, mtime and a temporary etag

   We now let the oCIS driver persist favorites, set temporary etags and the mtime as arbitrary
   metadata.

   Https://github.com/owncloud/ocis/issues/567
   https://github.com/cs3org/reva/issues/1394
   https://github.com/cs3org/reva/pull/1393

  * Enhancement #1403: Support remote cloud gathering metrics

   The current metrics package can only gather metrics either from json files. With this feature,
   the metrics can be gathered polling the http endpoints exposed by the owncloud/nextcloud
   sciencemesh apps.

   Https://github.com/cs3org/reva/pull/1403

   https://github.com/owncloud/ocis/pull/1372

* Enhancement - Update reva to v1.4.1-0.20210111080247-f2b63bfd6825: [#1194](https://github.com/owncloud/ocis/pull/1194)

  * Enhancement: calculate and expose actual file permission set [cs3org/reva#1368](https://github.com/cs3org/reva/pull/1368)
  * initial range request support [cs3org/reva#1326](https://github.com/cs3org/reva/pull/1388)

   https://github.com/owncloud/ocis/pull/1194
   https://github.com/cs3org/reva/pull/1368
   https://github.com/cs3org/reva/pull/1388
# Changelog for [1.0.0] (2020-12-17)

The following sections list the changes for 1.0.0.

## Summary

* Bugfix - Enable scrolling in accounts list: [#909](https://github.com/owncloud/ocis/pull/909)
* Bugfix - Add missing env vars to docker compose: [#392](https://github.com/owncloud/ocis/pull/392)
* Bugfix - Don't enforce empty external apps slice: [#473](https://github.com/owncloud/ocis/pull/473)
* Bugfix - Lower Bound was not working for the cs3 api index implementation: [#741](https://github.com/owncloud/ocis/pull/741)
* Bugfix - Accounts config sometimes being overwritten: [#808](https://github.com/owncloud/ocis/pull/808)
* Bugfix - Make settings service start without go coroutines: [#835](https://github.com/owncloud/ocis/pull/835)
* Bugfix - Fix button layout after phoenix update: [#625](https://github.com/owncloud/ocis/pull/625)
* Bugfix - Fix choose account dialogue: [#846](https://github.com/owncloud/ocis/pull/846)
* Bugfix - Fix id or username query handling: [#745](https://github.com/owncloud/ocis/pull/745)
* Bugfix - Fix konnectd build: [#809](https://github.com/owncloud/ocis/pull/809)
* Bugfix - Fix path of files shared with me in ocs api: [#204](https://github.com/owncloud/product/issues/204)
* Bugfix - Use micro default client: [#718](https://github.com/owncloud/ocis/pull/718)
* Bugfix - Allow consent-prompt with switch-account: [#788](https://github.com/owncloud/ocis/pull/788)
* Bugfix - Mint token with uid and gid: [#737](https://github.com/owncloud/ocis/pull/737)
* Bugfix - Serve index.html for directories: [#912](https://github.com/owncloud/ocis/pull/912)
* Bugfix - Don't create account if id/mail/username already taken: [#709](https://github.com/owncloud/ocis/pull/709)
* Bugfix - Fix director selection in proxy: [#521](https://github.com/owncloud/ocis/pull/521)
* Bugfix - Permission checks for settings write access: [#1092](https://github.com/owncloud/ocis/pull/1092)
* Bugfix - Fix minor ui bugs: [#1043](https://github.com/owncloud/ocis/issues/1043)
* Bugfix - Disable public link expiration by default: [#987](https://github.com/owncloud/ocis/issues/987)
* Bugfix - Build docker images with alpine:latest instead of alpine:edge: [#416](https://github.com/owncloud/ocis/pull/416)
* Change - Accounts UI shows message when no permissions: [#656](https://github.com/owncloud/ocis/pull/656)
* Change - Cache password validation: [#958](https://github.com/owncloud/ocis/pull/958)
* Change - Filesystem based index: [#709](https://github.com/owncloud/ocis/pull/709)
* Change - Rebuild index command for accounts: [#748](https://github.com/owncloud/ocis/pull/748)
* Change - Add the thumbnails command: [#156](https://github.com/owncloud/ocis/issues/156)
* Change - CS3 can be used as accounts-backend: [#1020](https://github.com/owncloud/ocis/pull/1020)
* Change - Use bcrypt to hash the user passwords: [#510](https://github.com/owncloud/ocis/issues/510)
* Change - Replace the library which scales the images: [#910](https://github.com/owncloud/ocis/pull/910)
* Change - Choose disk or cs3 storage for accounts and groups: [#623](https://github.com/owncloud/ocis/pull/623)
* Change - Enable OpenID dynamic client registration: [#811](https://github.com/owncloud/ocis/issues/811)
* Change - Integrate import command from ocis-migration: [#249](https://github.com/owncloud/ocis/pull/249)
* Change - Improve reva service descriptions: [#536](https://github.com/owncloud/ocis/pull/536)
* Change - Initial release of basic version: [#2](https://github.com/owncloud/ocis/issues/2)
* Change - Add cli-commands to manage accounts: [#115](https://github.com/owncloud/product/issues/115)
* Change - Start ocis-accounts with the ocis server command: [#25](https://github.com/owncloud/product/issues/25)
* Change - Properly style konnectd consent page: [#754](https://github.com/owncloud/ocis/pull/754)
* Change - Make all paths configurable and default to a common temp dir: [#1080](https://github.com/owncloud/ocis/pull/1080)
* Change - Move the indexer package from ocis/accounts to ocis/ocis-pkg: [#794](https://github.com/owncloud/ocis/pull/794)
* Change - Switch over to a new custom-built runtime: [#287](https://github.com/owncloud/ocis/pull/287)
* Change - Move ocis default config to root level: [#842](https://github.com/owncloud/ocis/pull/842)
* Change - Remove username field in OCS: [#709](https://github.com/owncloud/ocis/pull/709)
* Change - Account management permissions for Admin role: [#124](https://github.com/owncloud/product/issues/124)
* Change - Update phoenix to v0.18.0: [#651](https://github.com/owncloud/ocis/pull/651)
* Change - Default apps in ownCloud Web: [#688](https://github.com/owncloud/ocis/pull/688)
* Change - Proxy allow insecure upstreams: [#1007](https://github.com/owncloud/ocis/pull/1007)
* Change - Make ocis-settings available: [#287](https://github.com/owncloud/ocis/pull/287)
* Change - Start ocis-proxy with the ocis server command: [#119](https://github.com/owncloud/ocis/issues/119)
* Change - Theme welcome and choose account pages: [#887](https://github.com/owncloud/ocis/pull/887)
* Change - Bring oC theme: [#698](https://github.com/owncloud/ocis/pull/698)
* Change - Unify Configuration Parsing: [#675](https://github.com/owncloud/ocis/pull/675)
* Change - Update phoenix to v0.20.0: [#674](https://github.com/owncloud/ocis/pull/674)
* Change - Update phoenix to v0.21.0: [#728](https://github.com/owncloud/ocis/pull/728)
* Change - Update phoenix to v0.22.0: [#757](https://github.com/owncloud/ocis/pull/757)
* Change - Update phoenix to v0.23.0: [#785](https://github.com/owncloud/ocis/pull/785)
* Change - Update phoenix to v0.24.0: [#817](https://github.com/owncloud/ocis/pull/817)
* Change - Update phoenix to v0.25.0: [#868](https://github.com/owncloud/ocis/pull/868)
* Change - Update phoenix to v0.26.0: [#935](https://github.com/owncloud/ocis/pull/935)
* Change - Update phoenix to v0.27.0: [#943](https://github.com/owncloud/ocis/pull/943)
* Change - Update phoenix to v0.28.0: [#1027](https://github.com/owncloud/ocis/pull/1027)
* Change - Update phoenix to v0.29.0: [#1034](https://github.com/owncloud/ocis/pull/1034)
* Change - Update reva config: [#336](https://github.com/owncloud/ocis/pull/336)
* Change - Update reva to v1.4.1-0.20201209113234-e791b5599a89: [#1089](https://github.com/owncloud/ocis/pull/1089)
* Change - Clarify storage driver env vars: [#729](https://github.com/owncloud/ocis/pull/729)
* Change - Update ownCloud Web to v1.0.0-beta3: [#1105](https://github.com/owncloud/ocis/pull/1105)
* Change - Update ownCloud Web to v1.0.0-beta4: [#1110](https://github.com/owncloud/ocis/pull/1110)
* Change - Settings and accounts appear in the user menu: [#656](https://github.com/owncloud/ocis/pull/656)
* Enhancement - Add tracing to the accounts service: [#1016](https://github.com/owncloud/ocis/issues/1016)
* Enhancement - Add the accounts service: [#244](https://github.com/owncloud/product/issues/244)
* Enhancement - Add basic auth option: [#627](https://github.com/owncloud/ocis/pull/627)
* Enhancement - Document how to run OCIS on top of EOS: [#172](https://github.com/owncloud/ocis/pull/172)
* Enhancement - Add the glauth service: [#244](https://github.com/owncloud/product/issues/244)
* Enhancement - Add k6: [#941](https://github.com/owncloud/ocis/pull/941)
* Enhancement - Add the konnectd service: [#244](https://github.com/owncloud/product/issues/244)
* Enhancement - Add the ocis-phoenix service: [#244](https://github.com/owncloud/product/issues/244)
* Enhancement - Add the ocis-pkg package: [#244](https://github.com/owncloud/product/issues/244)
* Enhancement - Add the ocs service: [#244](https://github.com/owncloud/product/issues/244)
* Enhancement - Add the proxy service: [#244](https://github.com/owncloud/product/issues/244)
* Enhancement - Add the settings service: [#244](https://github.com/owncloud/product/issues/244)
* Enhancement - Add the storage service: [#244](https://github.com/owncloud/product/issues/244)
* Enhancement - Add the store service: [#244](https://github.com/owncloud/product/issues/244)
* Enhancement - Add the thumbnails service: [#244](https://github.com/owncloud/product/issues/244)
* Enhancement - Add a command to list the versions of running instances: [#226](https://github.com/owncloud/product/issues/226)
* Enhancement - Add the webdav service: [#244](https://github.com/owncloud/product/issues/244)
* Enhancement - Better adopt Go-Micro: [#840](https://github.com/owncloud/ocis/pull/840)
* Enhancement - Add permission check when assigning and removing roles: [#879](https://github.com/owncloud/ocis/issues/879)
* Enhancement - Create OnlyOffice extension: [#857](https://github.com/owncloud/ocis/pull/857)
* Enhancement - Show basic-auth warning only once: [#886](https://github.com/owncloud/ocis/pull/886)
* Enhancement - Add glauth fallback backend: [#649](https://github.com/owncloud/ocis/pull/649)
* Enhancement - Tidy dependencies: [#845](https://github.com/owncloud/ocis/pull/845)
* Enhancement - Launch a storage to store ocis-metadata: [#602](https://github.com/owncloud/ocis/pull/602)
* Enhancement - Add a version command to ocis: [#915](https://github.com/owncloud/ocis/pull/915)
* Enhancement - Create a proxy access-log: [#889](https://github.com/owncloud/ocis/pull/889)
* Enhancement - Cache userinfo in proxy: [#877](https://github.com/owncloud/ocis/pull/877)
* Enhancement - Update reva to v1.4.1-0.20201125144025-57da0c27434c: [#1320](https://github.com/cs3org/reva/pull/1320)
* Enhancement - Runtime Cleanup: [#1066](https://github.com/owncloud/ocis/pull/1066)
* Enhancement - Update OCIS Runtime: [#1108](https://github.com/owncloud/ocis/pull/1108)
* Enhancement - Simplify tracing config: [#92](https://github.com/owncloud/product/issues/92)
* Enhancement - Update glauth to dev fd3ac7e4bbdc93578655d9a08d8e23f105aaa5b2: [#834](https://github.com/owncloud/ocis/pull/834)
* Enhancement - Update glauth to dev 4f029234b2308: [#786](https://github.com/owncloud/ocis/pull/786)
* Enhancement - Update konnectd to v0.33.8: [#744](https://github.com/owncloud/ocis/pull/744)
* Enhancement - Update reva to v1.4.1-0.20201123062044-b2c4af4e897d: [#823](https://github.com/owncloud/ocis/pull/823)
* Enhancement - Update reva to v1.4.1-0.20201130061320-ac85e68e0600: [#980](https://github.com/owncloud/ocis/pull/980)
* Enhancement - Update reva to cdb3d6688da5: [#748](https://github.com/owncloud/ocis/pull/748)
* Enhancement - Update reva to dd3a8c0f38: [#725](https://github.com/owncloud/ocis/pull/725)
* Enhancement - Update reva to v1.4.1-0.20201127111856-e6a6212c1b7b: [#971](https://github.com/owncloud/ocis/pull/971)
* Enhancement - Update reva to 063b3db9162b: [#1091](https://github.com/owncloud/ocis/pull/1091)
* Enhancement - Add www-authenticate based on user agent: [#1009](https://github.com/owncloud/ocis/pull/1009)

## Details

* Bugfix - Enable scrolling in accounts list: [#909](https://github.com/owncloud/ocis/pull/909)

   Tags: accounts

   We've fixed the accounts list to enable scrolling.

   https://github.com/owncloud/ocis/pull/909

* Bugfix - Add missing env vars to docker compose: [#392](https://github.com/owncloud/ocis/pull/392)

   Tags: docker

   Without setting `REVA_FRONTEND_URL` and `REVA_DATAGATEWAY_URL` uploads would default to
   localhost and fail if `OCIS_DOMAIN` was used to run ocis on a remote host.

   https://github.com/owncloud/ocis/pull/392

* Bugfix - Don't enforce empty external apps slice: [#473](https://github.com/owncloud/ocis/pull/473)

   Tags: web

   The command for ocis-phoenix enforced an empty external apps configuration. This was
   removed, as it was blocking a new set of default external apps in ocis-phoenix.

   https://github.com/owncloud/ocis/pull/473

* Bugfix - Lower Bound was not working for the cs3 api index implementation: [#741](https://github.com/owncloud/ocis/pull/741)

   Tags: accounts

   Lower bound working on the cs3 index implementation

   https://github.com/owncloud/ocis/pull/741

* Bugfix - Accounts config sometimes being overwritten: [#808](https://github.com/owncloud/ocis/pull/808)

   Tags: accounts

   Sometimes when running the accounts extensions flags were not being taken into
   consideration.

   https://github.com/owncloud/ocis/pull/808

* Bugfix - Make settings service start without go coroutines: [#835](https://github.com/owncloud/ocis/pull/835)

   The go routines cause a race condition that sometimes causes the tests to fail. The ListRoles
   request would not return all permissions.

   https://github.com/owncloud/ocis/pull/835

* Bugfix - Fix button layout after phoenix update: [#625](https://github.com/owncloud/ocis/pull/625)

   Tags: accounts

   With the phoenix update to v0.17.0 a new ODS version was released which has a breaking change for
   buttons regarding their layout. We adjusted the button layout in the accounts UI accordingly.

   https://github.com/owncloud/ocis/pull/625

* Bugfix - Fix choose account dialogue: [#846](https://github.com/owncloud/ocis/pull/846)

   Tags: konnectd

   We've fixed the choose account dialogue in konnectd bug that the user hasn't been logged in
   after selecting account.

   https://github.com/owncloud/ocis/pull/846

* Bugfix - Fix id or username query handling: [#745](https://github.com/owncloud/ocis/pull/745)

   Tags: accounts

   The code was stopping execution when encountering an error while loading an account by id. But
   for or queries we can continue execution.

   https://github.com/owncloud/ocis/pull/745

* Bugfix - Fix konnectd build: [#809](https://github.com/owncloud/ocis/pull/809)

   Tags: konnectd

   We fixed the default config for konnectd and updated the Makefile to include the `yarn
   install`and `yarn build` steps if the static assets are missing.

   https://github.com/owncloud/ocis/pull/809

* Bugfix - Fix path of files shared with me in ocs api: [#204](https://github.com/owncloud/product/issues/204)

   The path of files shared with me using the ocs api was pointing to an incorrect location.

   https://github.com/owncloud/product/issues/204
   https://github.com/owncloud/ocis/pull/994

* Bugfix - Use micro default client: [#718](https://github.com/owncloud/ocis/pull/718)

   Tags: glauth

   We found a file descriptor leak in the glauth connections to the accounts service. Fixed it by
   using the micro default client.

   https://github.com/owncloud/ocis/pull/718

* Bugfix - Allow consent-prompt with switch-account: [#788](https://github.com/owncloud/ocis/pull/788)

   Multiple prompt values are allowed and this change fixes the check for select_account if it was
   used together with other prompt values. Where select_account previously was ignored, it is
   now processed as required, fixing the use case when a RP wants to trigger select_account first
   while at the same time wants also to request interactive consent.

   https://github.com/owncloud/ocis/pull/788

* Bugfix - Mint token with uid and gid: [#737](https://github.com/owncloud/ocis/pull/737)

   Tags: accounts

   The eos driver expects the uid and gid from the opaque map of a user. While the proxy does mint
   tokens correctly, the accounts service wasn't.

   https://github.com/owncloud/ocis/pull/737

* Bugfix - Serve index.html for directories: [#912](https://github.com/owncloud/ocis/pull/912)

   The static middleware in ocis-pkg now serves index.html instead of returning 404 on paths with
   a trailing `/`.

   https://github.com/owncloud/ocis-pkg/issues/63
   https://github.com/owncloud/ocis/pull/912

* Bugfix - Don't create account if id/mail/username already taken: [#709](https://github.com/owncloud/ocis/pull/709)

   Tags: accounts

   We don't allow anymore to create a new account if the provided id/mail/username is already
   taken.

   https://github.com/owncloud/ocis/pull/709

* Bugfix - Fix director selection in proxy: [#521](https://github.com/owncloud/ocis/pull/521)

   Tags: proxy

   We fixed a bug in ocis-proxy where simultaneous requests could be executed on the wrong
   backend.

   https://github.com/owncloud/ocis/pull/521
   https://github.com/owncloud/ocis-proxy/pull/99

* Bugfix - Permission checks for settings write access: [#1092](https://github.com/owncloud/ocis/pull/1092)

   Tags: settings

   There were several endpoints with write access to the settings service that were not protected
   by permission checks. We introduced a generic settings management permission to fix this for
   now. Will be more fine grained later on.

   https://github.com/owncloud/ocis/pull/1092

* Bugfix - Fix minor ui bugs: [#1043](https://github.com/owncloud/ocis/issues/1043)

   - the ui haven't updated the language of the items in the settings view menu. Now we listen to the
   selected language and update the ui - deduplicate resetMenuItems call

   https://github.com/owncloud/ocis/issues/1043
   https://github.com/owncloud/ocis/pull/1044

* Bugfix - Disable public link expiration by default: [#987](https://github.com/owncloud/ocis/issues/987)

   Tags: storage

   The public link expiration was enabled by default and didn't have a default expiration span by
   default, which resulted in already expired public links coming from the public link quick
   action. We fixed this by disabling the public link expiration by default.

   https://github.com/owncloud/ocis/issues/987
   https://github.com/owncloud/ocis/pull/1035

* Bugfix - Build docker images with alpine:latest instead of alpine:edge: [#416](https://github.com/owncloud/ocis/pull/416)

   Tags: docker

   ARM builds were failing when built on alpine:edge, so we switched to alpine:latest instead.

   https://github.com/owncloud/ocis/pull/416

* Change - Accounts UI shows message when no permissions: [#656](https://github.com/owncloud/ocis/pull/656)

   We improved the UX of the accounts UI by showing a message information the user about missing
   permissions when the accounts or roles fail to load. This was showing an indeterminate
   progress bar before.

   https://github.com/owncloud/ocis/pull/656

* Change - Cache password validation: [#958](https://github.com/owncloud/ocis/pull/958)

   Tags: accounts

   The password validity check for requests like `login eq '%s' and password eq '%s'` is now cached
   for 10 minutes. This improves the performance for basic auth requests.

   https://github.com/owncloud/ocis/pull/958

* Change - Filesystem based index: [#709](https://github.com/owncloud/ocis/pull/709)

   Tags: accounts, storage

   We replaced `bleve` with a new filesystem based index implementation. There is an `indexer`
   which is capable of orchestrating different index types to build indices on documents by
   field. You can choose from the index types `unique`, `non-unique` or `autoincrement`.
   Indices can be utilized to run search queries (full matches or globbing) on document fields.
   The accounts service is using this index internally to run the search queries coming in via
   `ListAccounts` and `ListGroups` and to generate UIDs for new accounts as well as GIDs for new
   groups.

   The accounts service can be configured to store the index on the local FS / a NFS (`disk`
   implementation of the index) or to use an arbitrary storage ( `cs3` implementation of the
   index). `cs3` is the new default, which is configured to use the `metadata` storage.

   https://github.com/owncloud/ocis/pull/709

* Change - Rebuild index command for accounts: [#748](https://github.com/owncloud/ocis/pull/748)

   Tags: accounts

   The index for the accounts service can now be rebuilt by running the cli command `./bin/ocis
   accounts rebuild`. It deletes all configured indices and rebuilds them from the documents
   found on storage. For this we also introduced a `LoadAccounts` and `LoadGroups` function on
   storage for loading all existing documents.

   https://github.com/owncloud/ocis/pull/748

* Change - Add the thumbnails command: [#156](https://github.com/owncloud/ocis/issues/156)

   Tags: thumbnails

   Added the thumbnails command so that the thumbnails service can get started via ocis.

   https://github.com/owncloud/ocis/issues/156

* Change - CS3 can be used as accounts-backend: [#1020](https://github.com/owncloud/ocis/pull/1020)

   Tags: proxy

   PROXY_ACCOUNT_BACKEND_TYPE=cs3 PROXY_ACCOUNT_BACKEND_TYPE=accounts (default)

   By using a backend which implements the CS3 user-api (currently provided by reva/storage) it
   is possible to bypass the ocis-accounts service and for example use ldap directly.

   https://github.com/owncloud/ocis/pull/1020

* Change - Use bcrypt to hash the user passwords: [#510](https://github.com/owncloud/ocis/issues/510)

   Change the hashing algorithm from SHA-512 to bcrypt since the latter is better suitable for
   password hashing. This is a breaking change. Existing deployments need to regenerate the
   accounts folder.

   https://github.com/owncloud/ocis/issues/510

* Change - Replace the library which scales the images: [#910](https://github.com/owncloud/ocis/pull/910)

   The library went out of support. Also did some refactoring of the thumbnails service code.

   https://github.com/owncloud/ocis/pull/910

* Change - Choose disk or cs3 storage for accounts and groups: [#623](https://github.com/owncloud/ocis/pull/623)

   Tags: accounts

   The accounts service now has an abstraction layer for the storage. In addition to the local disk
   implementation we implemented a cs3 storage, which is the new default for the accounts
   service.

   https://github.com/owncloud/ocis/pull/623

* Change - Enable OpenID dynamic client registration: [#811](https://github.com/owncloud/ocis/issues/811)

   Enable OpenID dynamic client registration

   https://github.com/owncloud/ocis/issues/811
   https://github.com/owncloud/ocis/pull/813

* Change - Integrate import command from ocis-migration: [#249](https://github.com/owncloud/ocis/pull/249)

   Tags: migration

   https://github.com/owncloud/ocis/pull/249
   https://github.com/owncloud/ocis-migration

* Change - Improve reva service descriptions: [#536](https://github.com/owncloud/ocis/pull/536)

   Tags: docs

   The descriptions make it clearer that the services actually represent a mount point in the
   combined storage. Each mount point can have a different driver.

   https://github.com/owncloud/ocis/pull/536

* Change - Initial release of basic version: [#2](https://github.com/owncloud/ocis/issues/2)

   Just prepared an initial basic version which simply embeds the minimum of required services in
   the context of the ownCloud Infinite Scale project.

   https://github.com/owncloud/ocis/issues/2

* Change - Add cli-commands to manage accounts: [#115](https://github.com/owncloud/product/issues/115)

   Tags: accounts

   COMMANDS:

  * list, ls        List existing accounts
  * add, create     Create a new account
  * update          Make changes to an existing account
  * remove, rm      Removes an existing account
  * inspect         Show detailed data on an existing account
  * help, h         Shows a list of commands or help for one command

   https://github.com/owncloud/product/issues/115

* Change - Start ocis-accounts with the ocis server command: [#25](https://github.com/owncloud/product/issues/25)

   Tags: accounts

   Starts ocis-accounts in single binary mode (./ocis server). This service stores the
   user-account information.

   https://github.com/owncloud/product/issues/25
   https://github.com/owncloud/ocis/pull/239/files

* Change - Properly style konnectd consent page: [#754](https://github.com/owncloud/ocis/pull/754)

   Tags: konnectd

   After bringing our theme into konnectd, we've had to adjust the styles of the consent page so the
   text is visible and button reflects our theme.

   https://github.com/owncloud/ocis/pull/754

* Change - Make all paths configurable and default to a common temp dir: [#1080](https://github.com/owncloud/ocis/pull/1080)

   Aligned all services to use a dir following`/var/tmp/ocis/<service>/...` by default. Also
   made some missing temp paths configurable via env vars and config flags.

   https://github.com/owncloud/ocis/pull/1080

* Change - Move the indexer package from ocis/accounts to ocis/ocis-pkg: [#794](https://github.com/owncloud/ocis/pull/794)

   We are making that change for semantic reasons. So consumers of any index don't necessarily
   need to know of the accounts service.

   https://github.com/owncloud/ocis/pull/794

* Change - Switch over to a new custom-built runtime: [#287](https://github.com/owncloud/ocis/pull/287)

   We moved away from using the go-micro runtime and are now using [our own
   runtime](https://github.com/refs/pman). This allows us to spawn service processes even
   when they are using different versions of go-micro. On top of that we now have the commands `ocis
   list`, `ocis kill` and `ocis run` available for service runtime management.

   https://github.com/owncloud/ocis/pull/287

* Change - Move ocis default config to root level: [#842](https://github.com/owncloud/ocis/pull/842)

   Tags: ocis

   We moved the tracing config to the `root` flagset so that they are parsed on all commands. We also
   introduced a `JWTSecret` flag in the root flagset, in order to apply a common default JWTSecret
   to all services that have one.

   https://github.com/owncloud/ocis/pull/842
   https://github.com/owncloud/ocis/pull/843

* Change - Remove username field in OCS: [#709](https://github.com/owncloud/ocis/pull/709)

   Tags: ocs

   We use the incoming userid as both the `id` and the `on_premises_sam_account_name` for new
   accounts in the accounts service. The userid in OCS requests is in fact the username, not our
   internal account id. We need to enforce the userid as our internal account id though, because
   the account id is part of various `path` formats.

   https://github.com/owncloud/ocis/pull/709
   https://github.com/owncloud/ocis/pull/816

* Change - Account management permissions for Admin role: [#124](https://github.com/owncloud/product/issues/124)

   Tags: accounts, settings

   We created an `AccountManagement` permission and added it to the default admin role. There are
   permission checks in place to protected http endpoints in ocis-accounts against requests
   without the permission. All existing default users (einstein, marie, richard) have the
   default user role now (doesn't have the `AccountManagement` permission). Additionally,
   there is a new default Admin user with credentials `moss:vista`.

   Known issue: for users without the `AccountManagement` permission, the accounts UI
   extension is still available in the ocis-web app switcher, but the requests for loading the
   users will fail (as expected). We are working on a way to hide the accounts UI extension if the
   user doesn't have the `AccountManagement` permission.

   https://github.com/owncloud/product/issues/124
   https://github.com/owncloud/ocis-settings/pull/59
   https://github.com/owncloud/ocis-settings/pull/66
   https://github.com/owncloud/ocis-settings/pull/67
   https://github.com/owncloud/ocis-settings/pull/69
   https://github.com/owncloud/ocis-proxy/pull/95
   https://github.com/owncloud/ocis-pkg/pull/59
   https://github.com/owncloud/ocis-accounts/pull/95
   https://github.com/owncloud/ocis-accounts/pull/100
   https://github.com/owncloud/ocis-accounts/pull/102

* Change - Update phoenix to v0.18.0: [#651](https://github.com/owncloud/ocis/pull/651)

   Tags: web

   We updated phoenix to v0.18.0. Please refer to the changelog (linked) for details on the
   phoenix release. With the ODS release brought in by phoenix we now have proper oc-checkbox and
   oc-radio components for the settings and accounts UI.

   https://github.com/owncloud/ocis/pull/651
   https://github.com/owncloud/phoenix/releases/tag/v0.18.0
   https://github.com/owncloud/owncloud-design-system/releases/tag/v1.12.1

* Change - Default apps in ownCloud Web: [#688](https://github.com/owncloud/ocis/pull/688)

   Tags: web

   We changed the default apps for ownCloud Web to be only files and media-viewer.
   Markdown-editor and draw-io have been removed as defaults.

   https://github.com/owncloud/ocis/pull/688

* Change - Proxy allow insecure upstreams: [#1007](https://github.com/owncloud/ocis/pull/1007)

   Tags: proxy

   We can now configure the proxy if insecure upstream servers are allowed. This was added since
   you need to disable certificate checks fore some situations like testing.

   https://github.com/owncloud/ocis/pull/1007

* Change - Make ocis-settings available: [#287](https://github.com/owncloud/ocis/pull/287)

   Tags: settings

   This version delivers `settings` as a new service. It is part of the array of services in the
   `server` command.

   https://github.com/owncloud/ocis/pull/287

* Change - Start ocis-proxy with the ocis server command: [#119](https://github.com/owncloud/ocis/issues/119)

   Tags: proxy

   Starts the proxy in single binary mode (./ocis server) on port 9200. The proxy serves as a
   single-entry point for all http-clients.

   https://github.com/owncloud/ocis/issues/119
   https://github.com/owncloud/ocis/issues/136

* Change - Theme welcome and choose account pages: [#887](https://github.com/owncloud/ocis/pull/887)

   Tags: konnectd

   We've themed the konnectd pages Welcome and Choose account. All text has a white color now to be
   easily readable on the dark background.

   https://github.com/owncloud/ocis/pull/887

* Change - Bring oC theme: [#698](https://github.com/owncloud/ocis/pull/698)

   Tags: konnectd

   We've styled our konnectd login page to reflect ownCloud theme.

   https://github.com/owncloud/ocis/pull/698

* Change - Unify Configuration Parsing: [#675](https://github.com/owncloud/ocis/pull/675)

   Tags: ocis

   - responsibility for config parsing should be on the subcommand - if there is a config file in the
   environment location, env var should take precedence - general rule of thumb: the more
   explicit the config file is that would be picked up. Order from less to more explicit: - config
   location (/etc/ocis) - environment variable - cli flag

   https://github.com/owncloud/ocis/pull/675

* Change - Update phoenix to v0.20.0: [#674](https://github.com/owncloud/ocis/pull/674)

   Tags: web

   We updated phoenix to v0.20.0. Please refer to the changelog (linked) for details on the
   phoenix release.

   https://github.com/owncloud/ocis/pull/674
   https://github.com/owncloud/phoenix/releases/tag/v0.20.0

* Change - Update phoenix to v0.21.0: [#728](https://github.com/owncloud/ocis/pull/728)

   Tags: web

   We updated phoenix to v0.21.0. Please refer to the changelog (linked) for details on the
   phoenix release.

   https://github.com/owncloud/ocis/pull/728
   https://github.com/owncloud/phoenix/releases/tag/v0.21.0

* Change - Update phoenix to v0.22.0: [#757](https://github.com/owncloud/ocis/pull/757)

   Tags: web

   We updated phoenix to v0.22.0. Please refer to the changelog (linked) for details on the
   phoenix release.

   https://github.com/owncloud/ocis/pull/757
   https://github.com/owncloud/phoenix/releases/tag/v0.22.0

* Change - Update phoenix to v0.23.0: [#785](https://github.com/owncloud/ocis/pull/785)

   Tags: web

   We updated phoenix to v0.23.0. Please refer to the changelog (linked) for details on the
   phoenix release.

   https://github.com/owncloud/ocis/pull/785
   https://github.com/owncloud/phoenix/releases/tag/v0.23.0

* Change - Update phoenix to v0.24.0: [#817](https://github.com/owncloud/ocis/pull/817)

   Tags: web

   We updated phoenix to v0.24.0. Please refer to the changelog (linked) for details on the
   phoenix release.

   https://github.com/owncloud/ocis/pull/817
   https://github.com/owncloud/phoenix/releases/tag/v0.24.0

* Change - Update phoenix to v0.25.0: [#868](https://github.com/owncloud/ocis/pull/868)

   Tags: web

   We updated phoenix to v0.25.0. Please refer to the changelog (linked) for details on the
   phoenix release.

   https://github.com/owncloud/ocis/pull/868
   https://github.com/owncloud/phoenix/releases/tag/v0.25.0

* Change - Update phoenix to v0.26.0: [#935](https://github.com/owncloud/ocis/pull/935)

   Tags: web

   We updated phoenix to v0.26.0. Please refer to the changelog (linked) for details on the
   phoenix release.

   https://github.com/owncloud/ocis/pull/935
   https://github.com/owncloud/phoenix/releases/tag/v0.26.0

* Change - Update phoenix to v0.27.0: [#943](https://github.com/owncloud/ocis/pull/943)

   Tags: web

   We updated phoenix to v0.27.0. Please refer to the changelog (linked) for details on the
   phoenix release.

   https://github.com/owncloud/ocis/pull/943
   https://github.com/owncloud/phoenix/releases/tag/v0.27.0

* Change - Update phoenix to v0.28.0: [#1027](https://github.com/owncloud/ocis/pull/1027)

   Tags: web

   We updated phoenix to v0.28.0. Please refer to the changelog (linked) for details on the
   phoenix release.

   https://github.com/owncloud/ocis/pull/1027
   https://github.com/owncloud/phoenix/releases/tag/v0.28.0

* Change - Update phoenix to v0.29.0: [#1034](https://github.com/owncloud/ocis/pull/1034)

   Tags: web

   We updated phoenix to v0.29.0. Please refer to the changelog (linked) for details on the
   phoenix release.

   https://github.com/owncloud/ocis/pull/1034
   https://github.com/owncloud/phoenix/releases/tag/v0.29.0

* Change - Update reva config: [#336](https://github.com/owncloud/ocis/pull/336)

  * EOS homes are not configured with an enable-flag anymore, but with a dedicated storage driver.
  * We're using it now and adapted default configs of storages

   https://github.com/owncloud/ocis/pull/336
   https://github.com/owncloud/ocis/pull/337
   https://github.com/owncloud/ocis/pull/338
   https://github.com/owncloud/ocis-reva/pull/891

* Change - Update reva to v1.4.1-0.20201209113234-e791b5599a89: [#1089](https://github.com/owncloud/ocis/pull/1089)

   Updated reva to v1.4.1-0.20201209113234-e791b5599a89

   https://github.com/owncloud/ocis/pull/1089

* Change - Clarify storage driver env vars: [#729](https://github.com/owncloud/ocis/pull/729)

   After renaming ocsi-reva to storage and combining the storage and data providers some env vars
   were confusingly named `STORAGE_STORAGE_...`. We are changing the prefix for driver related
   env vars to `STORAGE_DRIVER_...`. This makes changing the storage driver using eg.:
   `STORAGE_HOME_DRIVER=eos` and setting driver options using
   `STORAGE_DRIVER_EOS_LAYOUT=...` less confusing.

   https://github.com/owncloud/ocis/pull/729

* Change - Update ownCloud Web to v1.0.0-beta3: [#1105](https://github.com/owncloud/ocis/pull/1105)

   Tags: web

   We updated ownCloud Web to v1.0.0-beta3. Please refer to the changelog (linked) for details on
   the web release.

   https://github.com/owncloud/ocis/pull/1105
   https://github.com/owncloud/phoenix/releases/tag/v1.0.0-beta3

* Change - Update ownCloud Web to v1.0.0-beta4: [#1110](https://github.com/owncloud/ocis/pull/1110)

   Tags: web

   We updated ownCloud Web to v1.0.0-beta4. Please refer to the changelog (linked) for details on
   the web release.

   https://github.com/owncloud/ocis/pull/1110
   https://github.com/owncloud/phoenix/releases/tag/v1.0.0-beta4

* Change - Settings and accounts appear in the user menu: [#656](https://github.com/owncloud/ocis/pull/656)

   We moved settings and accounts to the user menu.

   https://github.com/owncloud/ocis/pull/656

* Enhancement - Add tracing to the accounts service: [#1016](https://github.com/owncloud/ocis/issues/1016)

   Added tracing to the accounts service.

   https://github.com/owncloud/ocis/issues/1016

* Enhancement - Add the accounts service: [#244](https://github.com/owncloud/product/issues/244)

   Tags: accounts

  * Bugfix - Initialize roleService client in GRPC server: [#114](https://github.com/owncloud/ocis-accounts/pull/114)
  * Bugfix - Cleanup separated indices in memory: [#224](https://github.com/owncloud/product/issues/224)
  * Change - Set user role on builtin users: [#102](https://github.com/owncloud/ocis-accounts/pull/102)
  * Change - Add new builtin admin user: [#102](https://github.com/owncloud/ocis-accounts/pull/102)
  * Change - We make use of the roles cache to enforce permission checks: [#100](https://github.com/owncloud/ocis-accounts/pull/100)
  * Change - We make use of the roles manager to enforce permission checks: [#108](https://github.com/owncloud/ocis-accounts/pull/108)
  * Enhancement - Add create account form: [#148](https://github.com/owncloud/product/issues/148)
  * Enhancement - Add delete accounts action: [#148](https://github.com/owncloud/product/issues/148)
  * Enhancement - Add enable/disable capabilities to the WebUI: [#118](https://github.com/owncloud/product/issues/118)
  * Enhancement - Improve visual appearance of accounts UI: [#222](https://github.com/owncloud/product/issues/222)
  * Bugfix - Adapting to new settings API for fetching roles: [#96](https://github.com/owncloud/ocis-accounts/pull/96)
  * Change - Create account api-call implicitly adds "default-user" role: [#173](https://github.com/owncloud/product/issues/173)
  * Change - Add role selection to accounts UI: [#103](https://github.com/owncloud/product/issues/103)
  * Bugfix - Atomic Requests: [#82](https://github.com/owncloud/ocis-accounts/pull/82)
  * Bugfix - Unescape value for prefix query: [#76](https://github.com/owncloud/ocis-accounts/pull/76)
  * Change - Adapt to new ocis-settings data model: [#87](https://github.com/owncloud/ocis-accounts/pull/87)
  * Change - Add permissions for language to default roles: [#88](https://github.com/owncloud/ocis-accounts/pull/88)
  * Bugfix - Add write mutexes: [#71](https://github.com/owncloud/ocis-accounts/pull/71)
  * Bugfix - Fix the accountId and groupId mismatch in DeleteGroup Method: [#60](https://github.com/owncloud/ocis-accounts/pull/60)
  * Bugfix - Fix index mapping: [#73](https://github.com/owncloud/ocis-accounts/issues/73)
  * Bugfix - Use NewNumericRangeInclusiveQuery for numeric literals: [#28](https://github.com/owncloud/ocis-glauth/issues/28)
  * Bugfix - Prevent segfault when no password is set: [#65](https://github.com/owncloud/ocis-accounts/pull/65)
  * Bugfix - Update account return value not used: [#70](https://github.com/owncloud/ocis-accounts/pull/70)
  * Bugfix - Build docker images with alpine:latest instead of alpine:edge: [#64](https://github.com/owncloud/ocis-accounts/pull/64)
  * Change - Align structure of this extension with other extensions: [#51](https://github.com/owncloud/ocis-accounts/pull/51)
  * Change - Change api errors: [#11](https://github.com/owncloud/ocis-accounts/issues/11)
  * Change - Enable accounts on creation: [#43](https://github.com/owncloud/ocis-accounts/issues/43)
  * Change - Fix index update on create/update: [#57](https://github.com/owncloud/ocis-accounts/issues/57)
  * Change - Pass around the correct logger throughout the code: [#41](https://github.com/owncloud/ocis-accounts/issues/41)
  * Change - Remove timezone setting: [#33](https://github.com/owncloud/ocis-accounts/pull/33)
  * Change - Tighten screws on usernames and email addresses: [#65](https://github.com/owncloud/ocis-accounts/pull/65)
  * Enhancement - Add early version of cli tools for user-management: [#69](https://github.com/owncloud/ocis-accounts/pull/69)
  * Enhancement - Update accounts API: [#30](https://github.com/owncloud/ocis-accounts/pull/30)
  * Enhancement - Add simple user listing UI: [#51](https://github.com/owncloud/ocis-accounts/pull/51)
  * Enhancement - Logging is configurable: [#24](https://github.com/owncloud/ocis-accounts/pull/24)
  * Change - Initial release of basic version: [#1](https://github.com/owncloud/ocis-accounts/issues/1)
  * Enhancement - Configuration: [#15](https://github.com/owncloud/ocis-accounts/pull/15)

   https://github.com/owncloud/product/issues/244

* Enhancement - Add basic auth option: [#627](https://github.com/owncloud/ocis/pull/627)

   We added a new `enable-basic-auth` option and `PROXY_ENABLE_BASIC_AUTH` environment
   variable that can be set to `true` to make the proxy verify the basic auth header with the
   accounts service. This should only be used for testing and development and is disabled by
   default.

   https://github.com/owncloud/product/issues/198
   https://github.com/owncloud/ocis/pull/627

* Enhancement - Document how to run OCIS on top of EOS: [#172](https://github.com/owncloud/ocis/pull/172)

   Tags: eos

   We have added rules to the Makefile that use the official [eos docker
   images](https://gitlab.cern.ch/eos/eos-docker) to boot an eos cluster and configure OCIS
   to use it.

   https://github.com/owncloud/ocis/pull/172

* Enhancement - Add the glauth service: [#244](https://github.com/owncloud/product/issues/244)

   Tags: glauth

  * Bugfix - Return invalid credentials when user was not found: [#30](https://github.com/owncloud/ocis-glauth/pull/30)
  * Bugfix - Query numeric attribute values without quotes: [#28](https://github.com/owncloud/ocis-glauth/issues/28)
  * Bugfix - Use searchBaseDN if already a user/group name: [#214](https://github.com/owncloud/product/issues/214)
  * Bugfix - Fix LDAP substring startswith filters: [#31](https://github.com/owncloud/ocis-glauth/pull/31)
  * Enhancement - Add build information to the metrics: [#226](https://github.com/owncloud/product/issues/226)
  * Enhancement - Reenable configuring backends: [#600](https://github.com/owncloud/ocis/pull/600)
  * Bugfix - Ignore case when comparing objectclass values: [#26](https://github.com/owncloud/ocis-glauth/pull/26)
  * Bugfix - Build docker images with alpine:latest instead of alpine:edge: [#24](https://github.com/owncloud/ocis-glauth/pull/24)
  * Enhancement - Handle ownCloudUUID attribute: [#27](https://github.com/owncloud/ocis-glauth/pull/27)
  * Enhancement - Implement group queries: [#22](https://github.com/owncloud/ocis-glauth/issues/22)
  * Enhancement - Configuration: [#11](https://github.com/owncloud/ocis-glauth/pull/11)
  * Enhancement - Improve default settings: [#12](https://github.com/owncloud/ocis-glauth/pull/12)
  * Enhancement - Generate temporary ldap certificates if LDAPS is enabled: [#12](https://github.com/owncloud/ocis-glauth/pull/12)
  * Enhancement - Provide additional tls-endpoint: [#12](https://github.com/owncloud/ocis-glauth/pull/12)
  * Change - Use physicist demo users: [#5](https://github.com/owncloud/ocis-glauth/issues/5)
  * Change - Default to config based user backend: [#6](https://github.com/owncloud/ocis-glauth/pull/6)

   https://github.com/owncloud/product/issues/244

* Enhancement - Add k6: [#941](https://github.com/owncloud/ocis/pull/941)

   Tags: tests

   Add k6 as a performance testing framework

   https://github.com/owncloud/ocis/pull/941
   https://github.com/owncloud/ocis/pull/983

* Enhancement - Add the konnectd service: [#244](https://github.com/owncloud/product/issues/244)

   Tags: konnectd

  * Enhancement - Add version command: [#226](https://github.com/owncloud/product/issues/226)
  * Bugfix - Add silent redirect url: [#69](https://github.com/owncloud/ocis-konnectd/issues/69)
  * Bugfix - Build docker images with alpine:latest instead of alpine:edge: [#71](https://github.com/owncloud/ocis-konnectd/pull/71)
  * Bugfix - Include the assets for #62: [#64](https://github.com/owncloud/ocis-konnectd/pull/64)
  * Bugfix - Redirect to the provided uri: [#26](https://github.com/owncloud/ocis-konnectd/issues/26)
  * Change - Add a trailing slash to trusted redirect uris: [#26](https://github.com/owncloud/ocis-konnectd/issues/26)
  * Change - Improve client identifiers for end users: [#62](https://github.com/owncloud/ocis-konnectd/pull/62)
  * Enhancement - Use upstream version of konnect library: [#14](https://github.com/owncloud/product/issues/14)
  * Enhancement - Change default config for single-binary: [#55](https://github.com/owncloud/ocis-konnectd/pull/55)
  * Bugfix - Generate a random CSP-Nonce in the webapp: [#17](https://github.com/owncloud/ocis-konnectd/issues/17)
  * Change - Dummy index.html is not required anymore by upstream: [#25](https://github.com/owncloud/ocis-konnectd/issues/25)
  * Change - Initial release of basic version: [#1](https://github.com/owncloud/ocis-konnectd/issues/1)
  * Change - Use glauth as ldap backend, default to running behind ocis-proxy: [#52](https://github.com/owncloud/ocis-konnectd/pull/52)

   https://github.com/owncloud/product/issues/244

* Enhancement - Add the ocis-phoenix service: [#244](https://github.com/owncloud/product/issues/244)

   Tags: web

  * Bugfix - Fix external app URLs: [#218](https://github.com/owncloud/product/issues/218)
  * Change - Remove pdf-viewer from default apps: [#85](https://github.com/owncloud/ocis-phoenix/pull/85)
  * Change - Enable Settings and Accounts apps by default: [#80](https://github.com/owncloud/ocis-phoenix/pull/80)
  * Bugfix - Exit when assets or config are not found: [#76](https://github.com/owncloud/ocis-phoenix/pull/76)
  * Bugfix - Build docker images with alpine:latest instead of alpine:edge: [#73](https://github.com/owncloud/ocis-phoenix/pull/73)
  * Change - Hide searchbar by default: [#116](https://github.com/owncloud/product/issues/116)
  * Bugfix - Allow silent refresh of access token: [#69](https://github.com/owncloud/ocis-konnectd/issues/69)
  * Change - Update Phoenix: [#60](https://github.com/owncloud/ocis-phoenix/pull/60)
  * Enhancement - Configuration: [#57](https://github.com/owncloud/ocis-phoenix/pull/57)
  * Bugfix - Config file value not being read: [#45](https://github.com/owncloud/ocis-phoenix/pull/45)
  * Change - Default to running behind ocis-proxy: [#55](https://github.com/owncloud/ocis-phoenix/pull/55)

   https://github.com/owncloud/product/issues/244

* Enhancement - Add the ocis-pkg package: [#244](https://github.com/owncloud/product/issues/244)

   Tags: ocis-pkg

  * Change - Unwrap roleIDs from access-token into metadata context: [#59](https://github.com/owncloud/ocis-pkg/pull/59)
  * Change - Provide cache for roles: [#59](https://github.com/owncloud/ocis-pkg/pull/59)
  * Change - Roles manager: [#60](https://github.com/owncloud/ocis-pkg/pull/60)
  * Change - Use go-micro's metadata context for account id: [#56](https://github.com/owncloud/ocis-pkg/pull/56)
  * Bugfix - Remove redigo 2.0.0+incompatible dependency: [#33](https://github.com/owncloud/ocis-graph/pull/33)
  * Change - Add middleware for x-access-token dismantling: [#46](https://github.com/owncloud/ocis-pkg/pull/46)
  * Enhancement - Add `ocis.id` and numeric id claims: [#50](https://github.com/owncloud/ocis-pkg/pull/50)
  * Bugfix - Pass flags to micro service: [#44](https://github.com/owncloud/ocis-pkg/pull/44)
  * Change - Add header to cors handler: [#41](https://github.com/owncloud/ocis-pkg/issues/41)
  * Enhancement - Tracing middleware: [#35](https://github.com/owncloud/ocis-pkg/pull/35/)
  * Enhancement - Allow http services to register handlers: [#33](https://github.com/owncloud/ocis-pkg/pull/33)
  * Change - Upgrade the micro libraries: [#22](https://github.com/owncloud/ocis-pkg/pull/22)
  * Bugfix - Fix Module Path: [#25](https://github.com/owncloud/ocis-pkg/pull/25)
  * Bugfix - Change import paths to ocis-pkg/v2: [#27](https://github.com/owncloud/ocis-pkg/pull/27)
  * Bugfix - Fix serving static assets: [#14](https://github.com/owncloud/ocis-pkg/pull/14)
  * Change - Add TLS support for http services: [#19](https://github.com/owncloud/ocis-pkg/issues/19)
  * Enhancement - Introduce OpenID Connect middleware: [#8](https://github.com/owncloud/ocis-pkg/issues/8)
  * Change - Add root path to static middleware: [#9](https://github.com/owncloud/ocis-pkg/issues/9)
  * Change - Better log level handling within micro: [#2](https://github.com/owncloud/ocis-pkg/issues/2)

   https://github.com/owncloud/product/issues/244

* Enhancement - Add the ocs service: [#244](https://github.com/owncloud/product/issues/244)

   Tags: ocs

  * Bugfix - Match the user response to the OC10 format: [#181](https://github.com/owncloud/product/issues/181)
  * Enhancement - Add version command: [#226](https://github.com/owncloud/product/issues/226)
  * Bugfix - Add the top level response structure to json responses: [#181](https://github.com/owncloud/product/issues/181)
  * Enhancement - Update ocis-accounts: [#42](https://github.com/owncloud/ocis-ocs/pull/42)
  * Bugfix - Mimic oc10 user enabled as string in provisioning api: [#39](https://github.com/owncloud/ocis-ocs/pull/39)
  * Bugfix - Use opaque ID of a user for signing keys: [#436](https://github.com/owncloud/ocis/issues/436)
  * Enhancement - Add option to create user with uidnumber and gidnumber: [#34](https://github.com/owncloud/ocis-ocs/pull/34)
  * Bugfix - Fix file descriptor leak: [#79](https://github.com/owncloud/ocis-accounts/issues/79)
  * Enhancement - Add Group management for OCS Provisioning API: [#25](https://github.com/owncloud/ocis-ocs/pull/25)
  * Enhancement - Basic Support for the User Provisioning API: [#23](https://github.com/owncloud/ocis-ocs/pull/23)
  * Bugfix - Build docker images with alpine:latest instead of alpine:edge: [#20](https://github.com/owncloud/ocis-ocs/pull/20)
  * Change - Initial release of basic version: [#1](https://github.com/owncloud/ocis-ocs/issues/1)
  * Change - Upgrade micro libraries: [#11](https://github.com/owncloud/ocis-ocs/issues/11)
  * Enhancement - Configuration: [#14](https://github.com/owncloud/ocis-ocs/pull/14)
  * Enhancement - Support signing key: [#18](https://github.com/owncloud/ocis-ocs/pull/18)

   https://github.com/owncloud/product/issues/244

* Enhancement - Add the proxy service: [#244](https://github.com/owncloud/product/issues/244)

   Tags: proxy

  * Bugfix - Fix director selection: [#99](https://github.com/owncloud/ocis-proxy/pull/99)
  * Bugfix - Add settings API and app endpoints to example config: [#93](https://github.com/owncloud/ocis-proxy/pull/93)
  * Change - Remove accounts caching: [#100](https://github.com/owncloud/ocis-proxy/pull/100)
  * Enhancement - Add autoprovision accounts flag: [#219](https://github.com/owncloud/product/issues/219)
  * Enhancement - Add hello API and app endpoints to example config and builtin config: [#96](https://github.com/owncloud/ocis-proxy/pull/96)
  * Enhancement - Add roleIDs to the access token: [#95](https://github.com/owncloud/ocis-proxy/pull/95)
  * Enhancement - Add version command: [#226](https://github.com/owncloud/product/issues/226)
  * Enhancement - Add numeric uid and gid to the access token: [#89](https://github.com/owncloud/ocis-proxy/pull/89)
  * Enhancement - Add configuration options for the pre-signed url middleware: [#91](https://github.com/owncloud/ocis-proxy/issues/91)
  * Bugfix - Enable new accounts by default: [#79](https://github.com/owncloud/ocis-proxy/pull/79)
  * Bugfix - Lookup user by id for presigned URLs: [#85](https://github.com/owncloud/ocis-proxy/pull/85)
  * Bugfix - Build docker images with alpine:latest instead of alpine:edge: [#78](https://github.com/owncloud/ocis-proxy/pull/78)
  * Change - Add settings and ocs group routes: [#81](https://github.com/owncloud/ocis-proxy/pull/81)
  * Change - Add route for user provisioning API in ocis-ocs: [#80](https://github.com/owncloud/ocis-proxy/pull/80)
  * Bugfix - Provide token configuration from config: [#69](https://github.com/owncloud/ocis-proxy/pull/69)
  * Bugfix - Provide token configuration from config: [#76](https://github.com/owncloud/ocis-proxy/pull/76)
  * Change - Add OIDC config flags: [#66](https://github.com/owncloud/ocis-proxy/pull/66)
  * Change - Mint new username property in the reva token: [#62](https://github.com/owncloud/ocis-proxy/pull/62)
  * Enhancement - Add Accounts UI routes: [#65](https://github.com/owncloud/ocis-proxy/pull/65)
  * Enhancement - Add option to disable TLS: [#71](https://github.com/owncloud/ocis-proxy/issues/71)
  * Enhancement - Only send create home request if an account has been migrated: [#52](https://github.com/owncloud/ocis-proxy/issues/52)
  * Enhancement - Create a root span on proxy that propagates down to consumers: [#64](https://github.com/owncloud/ocis-proxy/pull/64)
  * Enhancement - Support signed URLs: [#73](https://github.com/owncloud/ocis-proxy/issues/73)
  * Bugfix - Accounts service response was ignored: [#43](https://github.com/owncloud/ocis-proxy/pull/43)
  * Bugfix - Fix x-access-token in header: [#41](https://github.com/owncloud/ocis-proxy/pull/41)
  * Change - Point /data endpoint to reva frontend: [#45](https://github.com/owncloud/ocis-proxy/pull/45)
  * Change - Send autocreate home request to reva gateway: [#51](https://github.com/owncloud/ocis-proxy/pull/51)
  * Change - Update to new accounts API: [#39](https://github.com/owncloud/ocis-proxy/issues/39)
  * Enhancement - Retrieve Account UUID From User Claims: [#36](https://github.com/owncloud/ocis-proxy/pull/36)
  * Enhancement - Create account if it doesn't exist in ocis-accounts: [#55](https://github.com/owncloud/ocis-proxy/issues/55)
  * Enhancement - Disable keep-alive on server-side OIDC requests: [#268](https://github.com/owncloud/ocis/issues/268)
  * Enhancement - Make jwt secret configurable: [#41](https://github.com/owncloud/ocis-proxy/pull/41)
  * Enhancement - Respect account_enabled flag: [#53](https://github.com/owncloud/ocis-proxy/issues/53)
  * Change - Update ocis-pkg: [#30](https://github.com/owncloud/ocis-proxy/pull/30)
  * Change - Insecure http-requests are now redirected to https: [#29](https://github.com/owncloud/ocis-proxy/pull/29)
  * Enhancement - Configurable OpenID Connect client: [#27](https://github.com/owncloud/ocis-proxy/pull/27)
  * Enhancement - Add policy selectors: [#4](https://github.com/owncloud/ocis-proxy/issues/4)
  * Bugfix - Set TLS-Certificate correctly: [#25](https://github.com/owncloud/ocis-proxy/pull/25)
  * Change - Route requests based on regex or query parameters: [#21](https://github.com/owncloud/ocis-proxy/issues/21)
  * Enhancement - Proxy client urls in default configuration: [#19](https://github.com/owncloud/ocis-proxy/issues/19)
  * Enhancement - Make TLS-Cert configurable: [#14](https://github.com/owncloud/ocis-proxy/pull/14)
  * Enhancement - Load Proxy Policies at Runtime: [#17](https://github.com/owncloud/ocis-proxy/issues/17)

   https://github.com/owncloud/product/issues/244

* Enhancement - Add the settings service: [#244](https://github.com/owncloud/product/issues/244)

   Tags: settings

  * Bugfix - Fix loading and saving system scoped values: [#66](https://github.com/owncloud/ocis-settings/pull/66)
  * Bugfix - Complete input validation: [#66](https://github.com/owncloud/ocis-settings/pull/66)
  * Change - Add filter option for bundle ids in ListBundles and ListRoles: [#59](https://github.com/owncloud/ocis-settings/pull/59)
  * Change - Reuse roleIDs from the metadata context: [#69](https://github.com/owncloud/ocis-settings/pull/69)
  * Change - Update ocis-pkg/v2: [#72](https://github.com/owncloud/ocis-settings/pull/72)
  * Enhancement - Add version command: [#226](https://github.com/owncloud/product/issues/226)
  * Bugfix - Fix fetching bundles in settings UI: [#61](https://github.com/owncloud/ocis-settings/pull/61)
  * Change - Filter settings by permissions: [#99](https://github.com/owncloud/product/issues/99)
  * Change - Add role service: [#110](https://github.com/owncloud/product/issues/110)
  * Change - Rename endpoints and message types: [#36](https://github.com/owncloud/ocis-settings/issues/36)
  * Change - Use UUIDs instead of alphanumeric identifiers: [#46](https://github.com/owncloud/ocis-settings/pull/46)
  * Bugfix - Adjust UUID validation to be more tolerant: [#41](https://github.com/owncloud/ocis-settings/issues/41)
  * Bugfix - Fix runtime error when type asserting on nil value: [#38](https://github.com/owncloud/ocis-settings/pull/38)
  * Bugfix - Fix multiple submits on string and number form elements: [#745](https://github.com/owncloud/owncloud-design-system/issues/745)
  * Bugfix - Build docker images with alpine:latest instead of alpine:edge: [#39](https://github.com/owncloud/ocis-settings/pull/39)
  * Change - Dynamically add navItems for extensions with settings bundles: [#25](https://github.com/owncloud/ocis-settings/pull/25)
  * Change - Introduce input validation: [#22](https://github.com/owncloud/ocis-settings/pull/22)
  * Change - Use account uuid from x-access-token: [#14](https://github.com/owncloud/ocis-settings/pull/14)
  * Change - Use server config variable from ocis-web: [#34](https://github.com/owncloud/ocis-settings/pull/34)
  * Enhancement - Remove paths from Makefile: [#33](https://github.com/owncloud/ocis-settings/pull/33)
  * Enhancement - Extend the docs: [#11](https://github.com/owncloud/ocis-settings/issues/11)
  * Enhancement - Update ocis-pkg/v2: [#42](https://github.com/owncloud/ocis-settings/pull/42)

   https://github.com/owncloud/product/issues/244

* Enhancement - Add the storage service: [#244](https://github.com/owncloud/product/issues/244)

   Tags: storage, reva

  * Enhancement - Enable ocis driver treetime accounting: [#620](https://github.com/owncloud/ocis/pull/620)
  * Enhancement - Launch a storage to store ocis-metadata: [#602](https://github.com/owncloud/ocis/pull/602)

   In the future accounts, settings etc. should be stored in a dedicated metadata storage. The
   services should talk to this storage directly, bypassing reva-gateway.

   Https://github.com/owncloud/ocis/pull/602

  * Enhancement - Update reva to v1.2.2-0.20200924071957-e6676516e61e: [#601](https://github.com/owncloud/ocis/pull/601)

   - Update reva to v1.2.2-0.20200924071957-e6676516e61e - eos client: Handle eos EPERM as
   permission denied [(reva/#1183)](https://github.com/cs3org/reva/pull/1183) - ocis
   driver: synctime based etag propagation
   [(reva/#1180)](https://github.com/cs3org/reva/pull/1180) - ocis driver: fix litmus
   [(reva/#1179)](https://github.com/cs3org/reva/pull/1179) - ocis driver: fix move
   [(reva/#1177)](https://github.com/cs3org/reva/pull/1177) - ocs service: cache
   displaynames [(reva/#1161)](https://github.com/cs3org/reva/pull/1161)

   Https://github.com/owncloud/ocis-reva/issues/262
   https://github.com/owncloud/ocis-reva/issues/357
   https://github.com/owncloud/ocis-reva/issues/301
   https://github.com/owncloud/ocis-reva/issues/302
   https://github.com/owncloud/ocis/pull/601

  * Bugfix - Fix default configuration for accessing shares: [#205](https://github.com/owncloud/product/issues/205)

   The storage provider mounted at `/home` should always have EnableHome set to `true`. The other
   storage providers should have it set to `false`.

   Https://github.com/owncloud/product/issues/205
   https://github.com/owncloud/ocis-reva/pull/461

  * Enhancement - Allow configuring arbitrary storage registry rules: [#193](https://github.com/owncloud/product/issues/193)

   We added a new config flag `storage-registry-rule` that can be given multiple times for the
   gateway to specify arbitrary storage registry rules. You can also use a comma separated list of
   rules in the `REVA_STORAGE_REGISTRY_RULES` environment variable.

   Https://github.com/owncloud/product/issues/193
   https://github.com/owncloud/ocis-reva/pull/461

  * Enhancement - Update reva to v1.2.1-0.20200826162318-c0f54e1f37ea: [#454](https://github.com/owncloud/ocis-reva/pull/454)

   - Update reva to v1.2.1-0.20200826162318-c0f54e1f37ea - Do not swallow 'not found' errors in
   Stat [(reva/#1124)](https://github.com/cs3org/reva/pull/1124) - Rewire dav files to the
   home storage [(reva/#1125)](https://github.com/cs3org/reva/pull/1125) - Do not restore
   recycle entry on purge [(reva/#1099)](https://github.com/cs3org/reva/pull/1099) -
   Allow listing the trashbin [(reva/#1091)](https://github.com/cs3org/reva/pull/1091) -
   Restore and delete trash items via ocs
   [(reva/#1103)](https://github.com/cs3org/reva/pull/1103) - Ensure ignoring public
   stray shares [(reva/#1090)](https://github.com/cs3org/reva/pull/1090) - Ensure
   ignoring stray shares [(reva/#1064)](https://github.com/cs3org/reva/pull/1064) -
   Minor fixes in reva cmd, gateway uploads and smtpclient
   [(reva/#1082)](https://github.com/cs3org/reva/pull/1082) - Owncloud driver -
   propagate mtime on RemoveGrant
   [(reva/#1115)](https://github.com/cs3org/reva/pull/1115) - Handle redirection
   prefixes when extracting destination from URL
   [(reva/#1111)](https://github.com/cs3org/reva/pull/1111) - Add UID and GID in ldap auth
   driver [(reva/#1101)](https://github.com/cs3org/reva/pull/1101) - Add calens check to
   verify changelog entries in CI
   [(reva/#1077)](https://github.com/cs3org/reva/pull/1077) - Refactor Reva CLI with
   prompts [(reva/#1072)](https://github.com/cs3org/reva/pull/1072j) - Get file info
   using fxids from EOS [(reva/#1079)](https://github.com/cs3org/reva/pull/1079) - Update
   LDAP user driver [(reva/#1088)](https://github.com/cs3org/reva/pull/1088) - System
   information metrics cleanup
   [(reva/#1114)](https://github.com/cs3org/reva/pull/1114) - System information
   included in Prometheus metrics
   [(reva/#1071)](https://github.com/cs3org/reva/pull/1071) - Add logic for resolving
   storage references over webdav
   [(reva/#1094)](https://github.com/cs3org/reva/pull/1094)

   Https://github.com/owncloud/ocis-reva/pull/454

  * Enhancement - Update reva to v1.2.1-0.20200911111727-51649e37df2d: [#466](https://github.com/owncloud/ocis-reva/pull/466)

   - Update reva to v1.2.1-0.20200911111727-51649e37df2d - Added new OCIS storage driver ocis
   [(reva/#1155)](https://github.com/cs3org/reva/pull/1155) - App provider: fallback to
   env. variable if 'iopsecret' unset
   [(reva/#1146)](https://github.com/cs3org/reva/pull/1146) - Add switch to database
   [(reva/#1135)](https://github.com/cs3org/reva/pull/1135) - Add the ocdav HTTP svc to the
   standalone config [(reva/#1128)](https://github.com/cs3org/reva/pull/1128)

   Https://github.com/owncloud/ocis-reva/pull/466

  * Enhancement - Separate user and auth providers, add config for rest user: [#412](https://github.com/owncloud/ocis-reva/pull/412)

   Previously, the auth and user provider services used to have the same driver, which restricted
   using separate drivers and configs for both. This PR separates the two and adds the config for
   the rest user driver and the gatewaysvc parameter to EOS fs.

   Https://github.com/owncloud/ocis-reva/pull/412
   https://github.com/cs3org/reva/pull/995

  * Enhancement - Update reva to v1.1.1-0.20200819100654-dcbf0c8ea187: [#447](https://github.com/owncloud/ocis-reva/pull/447)

   - Update reva to v1.1.1-0.20200819100654-dcbf0c8ea187 - fix restoring and deleting trash
   items via ocs [(reva/#1103)](https://github.com/cs3org/reva/pull/1103) - Add UID and GID
   in ldap auth driver [(reva/#1101)](https://github.com/cs3org/reva/pull/1101) - Allow
   listing the trashbin [(reva/#1091)](https://github.com/cs3org/reva/pull/1091) -
   Ignore Stray Public Shares [(reva/#1090)](https://github.com/cs3org/reva/pull/1090) -
   Implement GetUserByClaim for LDAP user driver
   [(reva/#1088)](https://github.com/cs3org/reva/pull/1088) - eosclient: get file info by
   fxid [(reva/#1079)](https://github.com/cs3org/reva/pull/1079) - Ensure stray shares
   get ignored [(reva/#1064)](https://github.com/cs3org/reva/pull/1064) - Improve
   timestamp precision while logging
   [(reva/#1059)](https://github.com/cs3org/reva/pull/1059) - Ocfs lookup userid
   (update) [(reva/#1052)](https://github.com/cs3org/reva/pull/1052) - Disallow sharing
   the shares directory [(reva/#1051)](https://github.com/cs3org/reva/pull/1051) - Local
   storage provider: Fixed resolution of fileid
   [(reva/#1046)](https://github.com/cs3org/reva/pull/1046) - List public shares only
   created by the current user [(reva/#1042)](https://github.com/cs3org/reva/pull/1042)

   Https://github.com/owncloud/ocis-reva/pull/447

  * Bugfix - Update LDAP filters: [#399](https://github.com/owncloud/ocis-reva/pull/399)

   With the separation of use and find filters we can now use a filter that taken into account a users
   uuid as well as his username. This is necessary to make sharing work with the new account service
   which assigns accounts an immutable account id that is different from the username.
   Furthermore, the separate find filters now allows searching users by their displayname or
   email as well.

   ``` userfilter =
   "(&(objectclass=posixAccount)(|(ownclouduuid={{.OpaqueId}})(cn={{.OpaqueId}})))"
   findfilter =
   "(&(objectclass=posixAccount)(|(cn={{query}}*)(displayname={{query}}*)(mail={{query}}*)))"
   ```

   Https://github.com/owncloud/ocis-reva/pull/399
   https://github.com/cs3org/reva/pull/996

  * Change - Environment updates for the username userid split: [#420](https://github.com/owncloud/ocis-reva/pull/420)

   We updated the owncloud storage driver in reva to properly look up users by userid or username
   using the userprovider instead of taking the path segment as is. This requires the user service
   address as well as changing the default layout to the userid instead of the username. The latter
   is not considered a stable and persistent identifier.

   Https://github.com/owncloud/ocis-reva/pull/420
   https://github.com/cs3org/reva/pull/1033

  * Enhancement - Update storage documentation: [#384](https://github.com/owncloud/ocis-reva/pull/384)

   We added details to the documentation about storage requirements known from ownCloud 10, the
   local storage driver and the ownCloud storage driver.

   Https://github.com/owncloud/ocis-reva/pull/384
   https://github.com/owncloud/ocis-reva/pull/390

  * Enhancement - Update reva to v0.1.1-0.20200724135750-b46288b375d6: [#399](https://github.com/owncloud/ocis-reva/pull/399)

   - Update reva to v0.1.1-0.20200724135750-b46288b375d6 - Split LDAP user filters
   (reva/#996) - meshdirectory: Add invite forward API to provider links (reva/#1000) - OCM:
   Pass the link to the meshdirectory service in token mail (reva/#1002) - Update
   github.com/go-ldap/ldap to v3 (reva/#1004)

   Https://github.com/owncloud/ocis-reva/pull/399
   https://github.com/cs3org/reva/pull/996 https://github.com/cs3org/reva/pull/1000
   https://github.com/cs3org/reva/pull/1002 https://github.com/cs3org/reva/pull/1004

  * Enhancement - Update reva to v0.1.1-0.20200728071211-c948977dd3a0: [#407](https://github.com/owncloud/ocis-reva/pull/407)

   - Update reva to v0.1.1-0.20200728071211-c948977dd3a0 - Use proper logging for ldap auth
   requests (reva/#1008) - Update github.com/eventials/go-tus to
   v0.0.0-20200718001131-45c7ec8f5d59 (reva/#1007) - Check if SMTP credentials are nil
   (reva/#1006)

   Https://github.com/owncloud/ocis-reva/pull/407
   https://github.com/cs3org/reva/pull/1008 https://github.com/cs3org/reva/pull/1007
   https://github.com/cs3org/reva/pull/1006

  * Bugfix - Build docker images with alpine:latest instead of alpine:edge: [#393](https://github.com/owncloud/ocis-reva/pull/393)

   ARM builds were failing when built on alpine:edge, so we switched to alpine:latest instead.

   Https://github.com/owncloud/ocis-reva/pull/393

  * Enhancement - Update reva to v0.1.1-0.20200710143425-cf38a45220c5: [#371](https://github.com/owncloud/ocis-reva/pull/371)

   - Update reva to v0.1.1-0.20200710143425-cf38a45220c5 (#371) - Add wopi open (reva/#920) -
   Added a CS3API compliant data exporter to Mentix (reva/#955) - Read SMTP password from env if
   not set in config (reva/#953) - OCS share fix including file info after update (reva/#958) - Add
   flag to smtpclient for for unauthenticated SMTP (reva/#963)

   Https://github.com/owncloud/ocis-reva/pull/371
   https://github.com/cs3org/reva/pull/920 https://github.com/cs3org/reva/pull/953
   https://github.com/cs3org/reva/pull/955 https://github.com/cs3org/reva/pull/958
   https://github.com/cs3org/reva/pull/963

  * Enhancement - Update reva to v0.1.1-0.20200722125752-6dea7936f9d1: [#392](https://github.com/owncloud/ocis-reva/pull/392)

   - Update reva to v0.1.1-0.20200722125752-6dea7936f9d1 - Added signing key capability
   (reva/#986) - Add functionality to create webdav references for OCM shares (reva/#974) -
   Added a site locations exporter to Mentix (reva/#972) - Add option to config to allow requests
   to hosts with unverified certificates (reva/#969)

   Https://github.com/owncloud/ocis-reva/pull/392
   https://github.com/cs3org/reva/pull/986 https://github.com/cs3org/reva/pull/974
   https://github.com/cs3org/reva/pull/972 https://github.com/cs3org/reva/pull/969

  * Enhancement - Make frontend prefixes configurable: [#363](https://github.com/owncloud/ocis-reva/pull/363)

   We introduce three new environment variables and preconfigure them the following way:

  * `REVA_FRONTEND_DATAGATEWAY_PREFIX="data"`
  * `REVA_FRONTEND_OCDAV_PREFIX=""`
  * `REVA_FRONTEND_OCS_PREFIX="ocs"`

   This restores the reva defaults that were changed upstream.

   Https://github.com/owncloud/ocis-reva/pull/363
   https://github.com/cs3org/reva/pull/936/files#diff-51bf4fb310f7362f5c4306581132fc3bR63

  * Enhancement - Update reva to v0.1.1-0.20200701152626-2f6cc60e2f66: [#341](https://github.com/owncloud/ocis-reva/pull/341)

   - Update reva to v0.1.1-0.20200701152626-2f6cc60e2f66 (#341) - Added country information
   to Mentix (reva/#924) - Refactor metrics package to implement reader interface (reva/#934) -
   Fix OCS public link share update values logic (#252, #288, reva/#930)

   Https://github.com/owncloud/ocis-reva/issues/252
   https://github.com/owncloud/ocis-reva/issues/288
   https://github.com/owncloud/ocis-reva/pull/341
   https://github.com/cs3org/reva/pull/924 https://github.com/cs3org/reva/pull/934
   https://github.com/cs3org/reva/pull/930

  * Enhancement - Update reva to v0.1.1-0.20200709064551-91eed007038f: [#362](https://github.com/owncloud/ocis-reva/pull/362)

   - Update reva to v0.1.1-0.20200709064551-91eed007038f (#362) - Fix config for uploads when
   data server is not exposed (reva/#936) - Update OCM partners endpoints (reva/#937) - Update
   Ailleron endpoint (reva/#938) - OCS: Fix initialization of shares json file (reva/#940) -
   OCS: Fix returned public link URL (#336, reva/#945) - OCS: Share wrap resource id correctly
   (#344, reva/#951) - OCS: Implement share handling for accepting and listing shares (#11,
   reva/#929) - ocm: dynamically lookup IPs for provider check (reva/#946) - ocm: add
   functionality to mail OCM invite tokens (reva/#944) - Change percentagused to
   percentageused (reva/#903) - Fix file-descriptor leak (reva/#954)

   Https://github.com/owncloud/ocis-reva/issues/344
   https://github.com/owncloud/ocis-reva/issues/336
   https://github.com/owncloud/ocis-reva/issues/11
   https://github.com/owncloud/ocis-reva/pull/362
   https://github.com/cs3org/reva/pull/936 https://github.com/cs3org/reva/pull/937
   https://github.com/cs3org/reva/pull/938 https://github.com/cs3org/reva/pull/940
   https://github.com/cs3org/reva/pull/951 https://github.com/cs3org/reva/pull/945
   https://github.com/cs3org/reva/pull/929 https://github.com/cs3org/reva/pull/946
   https://github.com/cs3org/reva/pull/944 https://github.com/cs3org/reva/pull/903
   https://github.com/cs3org/reva/pull/954

  * Enhancement - Add new config options for the http client: [#330](https://github.com/owncloud/ocis-reva/pull/330)

   The internal certificates are checked for validity after
   https://github.com/cs3org/reva/pull/914, which causes the acceptance tests to fail. This
   change sets new hardcoded defaults.

   Https://github.com/owncloud/ocis-reva/pull/330

  * Enhancement - Allow datagateway transfers to take 24h: [#323](https://github.com/owncloud/ocis-reva/pull/323)

   - Increase transfer token life time to 24h (PR #323)

   Https://github.com/owncloud/ocis-reva/pull/323

  * Enhancement - Update reva to v0.1.1-0.20200630075923-39a90d431566: [#320](https://github.com/owncloud/ocis-reva/pull/320)

   - Update reva to v0.1.1-0.20200630075923-39a90d431566 (#320) - Return special value for
   public link password (#294, reva/#904) - Fix public stat and listcontainer response to
   contain the correct prefix (#310, reva/#902)

   Https://github.com/owncloud/ocis-reva/issues/310
   https://github.com/owncloud/ocis-reva/issues/294
   https://github.com/owncloud/ocis-reva/pull/320
   https://github.com/cs3org/reva/pull/902 https://github.com/cs3org/reva/pull/904

  * Enhancement - Update reva to v0.1.1-0.20200701152626-2f6cc60e2f66: [#328](https://github.com/owncloud/ocis-reva/pull/328)

   - Update reva to v0.1.1-0.20200701152626-2f6cc60e2f66 (#328) - Use sync.Map on pool package
   (reva/#909) - Use mutex instead of sync.Map (reva/#915) - Use gatewayProviders instead of
   storageProviders on conn pool (reva/#916) - Add logic to ls and stat to process arbitrary
   metadata keys (reva/#905) - Preliminary implementation of Set/UnsetArbitraryMetadata
   (reva/#912) - Make datagateway forward headers (reva/#913, reva/#926) - Add option to cmd
   upload to disable tus (reva/#911) - OCS Share Allow date-only expiration for public shares
   (#288, reva/#918) - OCS Share Remove array from OCS Share update response (#252, reva/#919) -
   OCS Share Implement GET request for single shares (#249, reva/#921)

   Https://github.com/owncloud/ocis-reva/issues/288
   https://github.com/owncloud/ocis-reva/issues/252
   https://github.com/owncloud/ocis-reva/issues/249
   https://github.com/owncloud/ocis-reva/pull/328
   https://github.com/cs3org/reva/pull/909 https://github.com/cs3org/reva/pull/915
   https://github.com/cs3org/reva/pull/916 https://github.com/cs3org/reva/pull/905
   https://github.com/cs3org/reva/pull/912 https://github.com/cs3org/reva/pull/913
   https://github.com/cs3org/reva/pull/926 https://github.com/cs3org/reva/pull/911
   https://github.com/cs3org/reva/pull/918 https://github.com/cs3org/reva/pull/919
   https://github.com/cs3org/reva/pull/921

  * Enhancement - Update reva to v0.1.1-0.20200629131207-04298ea1c088: [#309](https://github.com/owncloud/ocis-reva/pull/309)

   - Update reva to v0.1.1-0.20200629094927-e33d65230abc (#309) - Fix public link file share
   (#278, reva/#895, reva/#900) - Delete public share (reva/#899) - Updated reva to
   v0.1.1-0.20200629131207-04298ea1c088 (#313)

   Https://github.com/owncloud/ocis-reva/issues/278
   https://github.com/owncloud/ocis-reva/pull/309
   https://github.com/cs3org/reva/pull/895 https://github.com/cs3org/reva/pull/899
   https://github.com/cs3org/reva/pull/900
   https://github.com/owncloud/ocis-reva/pull/313

  * Enhancement - Update reva to v0.1.1-0.20200626111234-e21c32db9614: [#261](https://github.com/owncloud/ocis-reva/pull/261)

   - Updated reva to v0.1.1-0.20200626111234-e21c32db9614 (#304) - TUS upload support through
   datagateway (#261, reva/#878, reva/#888) - Added support for differing metrics path for
   Prometheus to Mentix (reva/#875) - More data exported by Mentix (reva/#881) - Implementation
   of file operations in public folder shares (#49, #293, reva/#877) - Make httpclient trust
   local certificates for now (reva/#880) - EOS homes are not configured with an enable-flag
   anymore, but with a dedicated storage driver. We're using it now and adapted default configs of
   storages (reva/#891, #304)

   Https://github.com/owncloud/ocis-reva/issues/49
   https://github.com/owncloud/ocis-reva/issues/293
   https://github.com/owncloud/ocis-reva/issues/261
   https://github.com/owncloud/ocis-reva/pull/261
   https://github.com/cs3org/reva/pull/875 https://github.com/cs3org/reva/pull/877
   https://github.com/cs3org/reva/pull/878 https://github.com/cs3org/reva/pull/881
   https://github.com/cs3org/reva/pull/880 https://github.com/cs3org/reva/pull/888
   https://github.com/owncloud/ocis-reva/pull/304
   https://github.com/cs3org/reva/pull/891

  * Enhancement - Update reva to v0.1.1-0.20200624063447-db5e6635d5f0: [#279](https://github.com/owncloud/ocis-reva/pull/279)

   - Updated reva to v0.1.1-0.20200624063447-db5e6635d5f0 (#279) - Local storage: URL-encode
   file ids to ease integration with other microservices like WOPI (reva/#799) - Mentix fixes
   (reva/#803, reva/#817) - OCDAV: fix returned timestamp format (#116, reva/#805) - OCM: add
   default prefix (#814) - add the content-length header to the responses (reva/#816) - Deps:
   clean (reva/#818) - Fix trashbin listing (#112, #253, #254, reva/#819) - Make the json
   publicshare driver configurable (reva/#820) - TUS: Return metadata headers after direct
   upload (ocis/#216, reva/#813) - Set mtime to storage after simple upload (#174, reva/#823,
   reva/#841) - Configure grpc client to allow for insecure conns and skip server certificate
   verification (reva/#825) - Deployment: simplify config with more default values
   (reva/#826, reva/#837, reva/#843, reva/#848, reva/#842) - Separate local fs into home and
   with home disabled (reva/#829) - Register reflection after other services (reva/#831) -
   Refactor EOS fs (reva/#830) - Add ocs-share-permissions to the propfind response (#47,
   reva/#836) - OCS: Properly read permissions when creating public link (reva/#852) - localfs:
   make normalize return associated error (reva/#850) - EOS grpc driver (reva/#664) - OCS: Add
   support for legacy public link arg publicUpload (reva/#853) - Add cache layer to user REST
   package (reva/#849) - Meshdirectory: pass query params to selected provider (reva/#863) -
   Pass etag in quotes from the fs layer (#269, reva/#866, reva/#867) - OCM: use refactored
   cs3apis provider definition (reva/#864)

   Https://github.com/owncloud/ocis-reva/issues/116
   https://github.com/owncloud/ocis-reva/issues/112
   https://github.com/owncloud/ocis-reva/issues/253
   https://github.com/owncloud/ocis-reva/issues/254
   https://github.com/owncloud/ocis/issues/216
   https://github.com/owncloud/ocis-reva/issues/174
   https://github.com/owncloud/ocis-reva/issues/47
   https://github.com/owncloud/ocis-reva/issues/269
   https://github.com/owncloud/ocis-reva/pull/279
   https://github.com/owncloud/cs3org/reva/pull/799
   https://github.com/owncloud/cs3org/reva/pull/803
   https://github.com/owncloud/cs3org/reva/pull/817
   https://github.com/owncloud/cs3org/reva/pull/805
   https://github.com/owncloud/cs3org/reva/pull/814
   https://github.com/owncloud/cs3org/reva/pull/816
   https://github.com/owncloud/cs3org/reva/pull/818
   https://github.com/owncloud/cs3org/reva/pull/819
   https://github.com/owncloud/cs3org/reva/pull/820
   https://github.com/owncloud/cs3org/reva/pull/823
   https://github.com/owncloud/cs3org/reva/pull/841
   https://github.com/owncloud/cs3org/reva/pull/813
   https://github.com/owncloud/cs3org/reva/pull/825
   https://github.com/owncloud/cs3org/reva/pull/826
   https://github.com/owncloud/cs3org/reva/pull/837
   https://github.com/owncloud/cs3org/reva/pull/843
   https://github.com/owncloud/cs3org/reva/pull/848
   https://github.com/owncloud/cs3org/reva/pull/842
   https://github.com/owncloud/cs3org/reva/pull/829
   https://github.com/owncloud/cs3org/reva/pull/831
   https://github.com/owncloud/cs3org/reva/pull/830
   https://github.com/owncloud/cs3org/reva/pull/836
   https://github.com/owncloud/cs3org/reva/pull/852
   https://github.com/owncloud/cs3org/reva/pull/850
   https://github.com/owncloud/cs3org/reva/pull/664
   https://github.com/owncloud/cs3org/reva/pull/853
   https://github.com/owncloud/cs3org/reva/pull/849
   https://github.com/owncloud/cs3org/reva/pull/863
   https://github.com/owncloud/cs3org/reva/pull/866
   https://github.com/owncloud/cs3org/reva/pull/867
   https://github.com/owncloud/cs3org/reva/pull/864

  * Enhancement - Add TUS global capability: [#177](https://github.com/owncloud/ocis-reva/issues/177)

   The TUS global capabilities from Reva are now exposed.

   The advertised max chunk size can be configured using the "--upload-max-chunk-size" CLI
   switch or "REVA_FRONTEND_UPLOAD_MAX_CHUNK_SIZE" environment variable. The advertised
   http method override can be configured using the "--upload-http-method-override" CLI
   switch or "REVA_FRONTEND_UPLOAD_HTTP_METHOD_OVERRIDE" environment variable.

   Https://github.com/owncloud/ocis-reva/issues/177
   https://github.com/owncloud/ocis-reva/pull/228

  * Enhancement - Update reva to v0.1.1-0.20200603071553-e05a87521618: [#244](https://github.com/owncloud/ocis-reva/issues/244)

   - Updated reva to v0.1.1-0.20200603071553-e05a87521618 (#244) - Add option to disable TUS on
   OC layer (#177, reva/#791) - Dataprovider now supports method override (#177, reva/#792) -
   OCS fixes for create public link (reva/#798)

   Https://github.com/owncloud/ocis-reva/issues/244
   https://github.com/owncloud/ocis-reva/issues/177
   https://github.com/cs3org/reva/pull/791 https://github.com/cs3org/reva/pull/792
   https://github.com/cs3org/reva/pull/798

  * Enhancement - Add public shares service: [#49](https://github.com/owncloud/ocis-reva/issues/49)

   Added Public Shares service with CRUD operations and File Public Shares Manager

   Https://github.com/owncloud/ocis-reva/issues/49
   https://github.com/owncloud/ocis-reva/pull/232

  * Enhancement - Update reva to v0.1.1-0.20200529120551-4f2d9c85d3c9: [#49](https://github.com/owncloud/ocis-reva/issues/49)

   - Updated reva to v0.1.1-0.20200529120551 (#232) - Public Shares CRUD, File Public Shares
   Manager (#49, #232, reva/#681, reva/#788) - Disable HTTP-KeepAlives to reduce fd count
   (ocis/#268, reva/#787) - Fix trashbin listing (#229, reva/#782) - Create PUT wrapper for TUS
   uploads (reva/#770) - Add security access headers for ocdav requests (#66, reva/#780) - Add
   option to revad cmd to specify logging level (reva/#772) - New metrics package (reva/#740) -
   Remove implicit data member from memory store (reva/#774) - Added TUS global capabilities
   (#177, reva/#775) - Fix PROPFIND with Depth 1 for cross-storage operations (reva/#779)

   Https://github.com/owncloud/ocis-reva/issues/49
   https://github.com/owncloud/ocis-reva/issues/229
   https://github.com/owncloud/ocis-reva/issues/66
   https://github.com/owncloud/ocis-reva/issues/177
   https://github.com/owncloud/ocis/issues/268
   https://github.com/owncloud/ocis-reva/pull/232
   https://github.com/cs3org/reva/pull/787 https://github.com/cs3org/reva/pull/681
   https://github.com/cs3org/reva/pull/788 https://github.com/cs3org/reva/pull/782
   https://github.com/cs3org/reva/pull/770 https://github.com/cs3org/reva/pull/780
   https://github.com/cs3org/reva/pull/772 https://github.com/cs3org/reva/pull/740
   https://github.com/cs3org/reva/pull/774 https://github.com/cs3org/reva/pull/775
   https://github.com/cs3org/reva/pull/779

  * Enhancement - Update reva to v0.1.1-0.20200520150229: [#161](https://github.com/owncloud/ocis-reva/pull/161)

   - Update reva to v0.1.1-0.20200520150229 (#161, #180, #192, #207, #221) - Return arbitrary
   metadata with stat, upload without TUS (reva/#766) - Stat file before returning datagateway
   URL when initiating download (reva/#765) - REST driver for user package (reva/#747) - Sharing
   behavior now consistent with the old backend (#20, #26, #43, #44, #46, #94 ,reva/#748) - Mentix
   service (reva/#755) - meshdirectory: add mentix driver for gocdb sites integration
   (reva/#754) - Add functionality to commit to storage for OCM shares (reva/#760) - Add option in
   config to disable tus (reva/#759) - ocdav: fix custom property XML parsing in PROPPATCH
   handler (#203, reva/#743) - ocdav: fix PROPPATCH response for removed properties (#186,
   reva/#742) - ocdav: implement PROPFIND infinity depth (#212, reva/#758) - Local fs: Allow
   setting of arbitrary metadata, minor bug fixes (reva/#764) - Local fs: metadata handling and
   share persistence (reva/#732) - Local fs: return file owner info in stat (reva/#750) - Fixed
   regression when uploading empty files to OCFS or EOS with PUT and TUS (#188, reva/#734) - On
   delete move the file versions to the trashbin (#94, reva/#731) - Fix OCFS move operation (#182,
   reva/#729) - Fix OCFS custom property / xattr removal (reva/#728) - Retry trashbin in case of
   timestamp collision (reva/#730) - Disable chunking v1 by default (reva/#678) - Implement ocs
   to http status code mapping (#26, reva/#696, reva/#707, reva/#711) - Handle the case if
   directory already exists (reva/#695) - Added TUS upload support (reva/#674, reva/#725,
   reva/#717) - Always return file sizes in Webdav PROPFIND (reva/#712) - Use default mime type
   when none was detected (reva/#713) - Fixed Webdav shallow COPY (reva/#714) - Fixed arbitrary
   namespace usage for custom properties in PROPFIND (#57, reva/#720) - Implement returning
   Webdav custom properties from xattr (#57, reva/#721) - Minor fix in OCM share pkg (reva/#718)

   Https://github.com/owncloud/ocis-reva/issues/20
   https://github.com/owncloud/ocis-reva/issues/26
   https://github.com/owncloud/ocis-reva/issues/43
   https://github.com/owncloud/ocis-reva/issues/44
   https://github.com/owncloud/ocis-reva/issues/46
   https://github.com/owncloud/ocis-reva/issues/94
   https://github.com/owncloud/ocis-reva/issues/26
   https://github.com/owncloud/ocis-reva/issues/67
   https://github.com/owncloud/ocis-reva/issues/57
   https://github.com/owncloud/ocis-reva/issues/94
   https://github.com/owncloud/ocis-reva/issues/188
   https://github.com/owncloud/ocis-reva/issues/182
   https://github.com/owncloud/ocis-reva/issues/212
   https://github.com/owncloud/ocis-reva/issues/186
   https://github.com/owncloud/ocis-reva/issues/203
   https://github.com/owncloud/ocis-reva/pull/161
   https://github.com/owncloud/ocis-reva/pull/180
   https://github.com/owncloud/ocis-reva/pull/192
   https://github.com/owncloud/ocis-reva/pull/207
   https://github.com/owncloud/ocis-reva/pull/221
   https://github.com/cs3org/reva/pull/766 https://github.com/cs3org/reva/pull/765
   https://github.com/cs3org/reva/pull/755 https://github.com/cs3org/reva/pull/754
   https://github.com/cs3org/reva/pull/747 https://github.com/cs3org/reva/pull/748
   https://github.com/cs3org/reva/pull/760 https://github.com/cs3org/reva/pull/759
   https://github.com/cs3org/reva/pull/678 https://github.com/cs3org/reva/pull/696
   https://github.com/cs3org/reva/pull/707 https://github.com/cs3org/reva/pull/711
   https://github.com/cs3org/reva/pull/695 https://github.com/cs3org/reva/pull/674
   https://github.com/cs3org/reva/pull/725 https://github.com/cs3org/reva/pull/717
   https://github.com/cs3org/reva/pull/712 https://github.com/cs3org/reva/pull/713
   https://github.com/cs3org/reva/pull/720 https://github.com/cs3org/reva/pull/718
   https://github.com/cs3org/reva/pull/731 https://github.com/cs3org/reva/pull/734
   https://github.com/cs3org/reva/pull/729 https://github.com/cs3org/reva/pull/728
   https://github.com/cs3org/reva/pull/730 https://github.com/cs3org/reva/pull/758
   https://github.com/cs3org/reva/pull/742 https://github.com/cs3org/reva/pull/764
   https://github.com/cs3org/reva/pull/743 https://github.com/cs3org/reva/pull/732
   https://github.com/cs3org/reva/pull/750

  * Bugfix - Stop advertising unsupported chunking v2: [#145](https://github.com/owncloud/ocis-reva/pull/145)

   Removed "chunking" attribute in the DAV capabilities. Please note that chunking v2 is
   advertised as "chunking 1.0" while chunking v1 is the attribute "bigfilechunking" which is
   already false.

   Https://github.com/owncloud/ocis-reva/pull/145

  * Enhancement - Allow configuring the gateway for dataproviders: [#136](https://github.com/owncloud/ocis-reva/pull/136)

   This allows using basic or bearer auth when directly talking to dataproviders.

   Https://github.com/owncloud/ocis-reva/pull/136

  * Enhancement - Use a configured logger on reva runtime: [#153](https://github.com/owncloud/ocis-reva/pull/153)

   For consistency reasons we need a configured logger that is inline with an ocis logger, so the
   log cascade can be easily parsed by a human.

   Https://github.com/owncloud/ocis-reva/pull/153

  * Bugfix - Fix eos user sharing config: [#127](https://github.com/owncloud/ocis-reva/pull/127)

   We have added missing config options for the user sharing manager and added a dedicated eos
   storage command with pre configured settings for the eos-docker container. It configures a
   `Shares` folder in a users home when using eos as the storage driver.

   Https://github.com/owncloud/ocis-reva/pull/127

  * Enhancement - Update reva to v1.1.0-20200414133413: [#127](https://github.com/owncloud/ocis-reva/pull/127)

   Adds initial public sharing and ocm implementation.

   Https://github.com/owncloud/ocis-reva/pull/127

  * Bugfix - Fix eos config: [#125](https://github.com/owncloud/ocis-reva/pull/125)

   We have added missing config options for the home layout to the config struct that is passed to
   eos.

   Https://github.com/owncloud/ocis-reva/pull/125

  * Bugfix - Set correct flag type in the flagsets: [#75](https://github.com/owncloud/ocis-reva/issues/75)

   While upgrading to the micro/cli version 2 there where two instances of `StringFlag` which had
   not been changed to `StringSliceFlag`. This caused `ocis-reva users` and `ocis-reva
   storage-root` to fail on startup.

   Https://github.com/owncloud/ocis-reva/issues/75
   https://github.com/owncloud/ocis-reva/pull/76

  * Bugfix - We fixed a typo in the `REVA_LDAP_SCHEMA_MAIL` environment variable: [#113](https://github.com/owncloud/ocis-reva/pull/113)

   It was misspelled as `REVA_LDAP_SCHEMA_Mail`.

   Https://github.com/owncloud/ocis-reva/pull/113

  * Bugfix - Allow different namespaces for /webdav and /dav/files: [#68](https://github.com/owncloud/ocis-reva/pull/68)

   After fbf131c the path for the "new" webdav path does not contain a username
   `/remote.php/dav/files/textfile0.txt`. It used to be
   `/remote.php/dav/files/oc/einstein/textfile0.txt` So it lost `oc/einstein`.

   This PR allows setting up different namespaces for `/webav` and `/dav/files`:

   `/webdav` is jailed into `/home` - which uses the home storage driver and uses the logged in user
   to construct the path `/dav/files` is jailed into `/oc` - which uses the owncloud storage
   driver and expects a username as the first path segment

   This mimics oc10

   The `WEBDAV_NAMESPACE_JAIL` environment variable is split into - `WEBDAV_NAMESPACE` and -
   `DAV_FILES_NAMESPACE` accordingly.

   Https://github.com/owncloud/ocis-reva/pull/68 related:

  * Change - Use /home as default namespace: [#68](https://github.com/owncloud/ocis-reva/pull/68)

   Currently, cross storage etag propagation is not yet implemented, which prevents the desktop
   client from detecting changes via the PROPFIND to /. / is managed by the root storage provider
   which is independent of the home and oc storage providers. If a file changes in /home/foo, the
   etag change will only be propagated to the root of the home storage provider.

   This change jails users into the `/home` namespace, and allows configuring the namespace to
   use for the two webdav endpoints using the new environment variable `WEBDAV_NAMESPACE_JAIL`
   which affects both endpoints `/dav/files` and `/webdav`.

   This will allow us to focus on getting a single storage driver like eos or owncloud tested and
   better resembles what owncloud 10 does.

   To get back the global namespace, which ultimately is the goal, just set the above environment
   variable to `/`.

   Https://github.com/owncloud/ocis-reva/pull/68

  * Change - Initial release of basic version: [#1](https://github.com/owncloud/ocis-reva/issues/1)

   Just prepared an initial basic version to start a reva server and start integrating with the
   go-micro base dextension framework of ownCloud Infinite Scale.

   Https://github.com/owncloud/ocis-reva/issues/1

  * Change - Start multiple services with dedicated commands: [#6](https://github.com/owncloud/ocis-reva/issues/6)

   The initial version would only allow us to use a set of reva configurations to start multiple
   services. We use a more opinionated set of commands to start dedicated services that allows us
   to configure them individually. It allows us to switch eg. the user backend to LDAP and fully use
   it on the cli.

   Https://github.com/owncloud/ocis-reva/issues/6

  * Change - Storage providers now default to exposing data servers: [#89](https://github.com/owncloud/ocis-reva/issues/89)

   The flags that let reva storage providers announce that they expose a data server now defaults
   to true:

   `REVA_STORAGE_HOME_EXPOSE_DATA_SERVER=1` `REVA_STORAGE_OC_EXPOSE_DATA_SERVER=1`

   Https://github.com/owncloud/ocis-reva/issues/89

  * Change - Default to running behind ocis-proxy: [#113](https://github.com/owncloud/ocis-reva/pull/113)

   We changed the default configuration to integrate better with ocis.

   - We use ocis-glauth as the default ldap server on port 9125 with base `dc=example,dc=org`. - We
   use a dedicated technical `reva` user to make ldap binds - Clients are supposed to use the
   ocis-proxy endpoint `https://localhost:9200` - We removed unneeded ocis configuration
   from the frontend which no longer serves an oidc provider. - We changed the default user
   OpaqueID attribute from `sub` to `preferred_username`. The latter is a claim populated by
   konnectd that can also be used by the reva ldap user manager to look up users by their OpaqueId

   Https://github.com/owncloud/ocis-reva/pull/113

  * Enhancement - Expose owncloud storage driver config in flagset: [#87](https://github.com/owncloud/ocis-reva/issues/87)

   Three new flags are now available:

   - scan files on startup to generate missing fileids default: `true` env var:
   `REVA_STORAGE_OWNCLOUD_SCAN` cli option: `--storage-owncloud-scan`

   - autocreate home path for new users default: `true` env var:
   `REVA_STORAGE_OWNCLOUD_AUTOCREATE` cli option: `--storage-owncloud-autocreate`

   - the address of the redis server default: `:6379` env var:
   `REVA_STORAGE_OWNCLOUD_REDIS_ADDR` cli option: `--storage-owncloud-redis`

   Https://github.com/owncloud/ocis-reva/issues/87

  * Enhancement - Update reva to v0.0.2-0.20200212114015-0dbce24f7e8b: [#91](https://github.com/owncloud/ocis-reva/pull/91)

   Reva has seen a lot of changes that allow us to - reduce the configuration overhead - use the
   autocreate home folder option - use the home folder path layout option - no longer start the root
   storage

   Https://github.com/owncloud/ocis-reva/pull/91 related:

  * Enhancement - Allow configuring user sharing driver: [#115](https://github.com/owncloud/ocis-reva/pull/115)

   We now default to `json` which persists shares in the sharing manager in a json file instead of an
   in memory db.

   Https://github.com/owncloud/ocis-reva/pull/115

   https://github.com/owncloud/product/issues/244

* Enhancement - Add the store service: [#244](https://github.com/owncloud/product/issues/244)

   Tags: store

  * Enhancement - Add version command: [#226](https://github.com/owncloud/product/issues/226)
  * Bugfix - Removed code from other service: [#7](https://github.com/owncloud/ocis-store/pull/7)
  * Bugfix - Build docker images with alpine:latest instead of alpine:edge: [#5](https://github.com/owncloud/ocis-store/pull/5)
  * Change - Initial release of basic version: [#1](https://github.com/owncloud/ocis-store/pull/1)

   https://github.com/owncloud/product/issues/244

* Enhancement - Add the thumbnails service: [#244](https://github.com/owncloud/product/issues/244)

   Tags: thumbnails

  * Enhancement - Add version command: [#226](https://github.com/owncloud/product/issues/226)
  * Bugfix - Build docker images with alpine:latest instead of alpine:edge: [#35](https://github.com/owncloud/ocis-thumbnails/pull/35)
  * Enhancement - Serve the metrics endpoint: [#37](https://github.com/owncloud/ocis-thumbnails/issues/37)
  * Change - Add more default resolutions: [#23](https://github.com/owncloud/ocis-thumbnails/issues/23)
  * Change - Refactor code to remove code smells: [#21](https://github.com/owncloud/ocis-thumbnails/issues/21)
  * Change - Use micro service error api: [#31](https://github.com/owncloud/ocis-thumbnails/issues/31)
  * Enhancement - Limit users to access own thumbnails: [#5](https://github.com/owncloud/ocis-thumbnails/issues/5)
  * Bugfix - Fix usage of context.Context: [#18](https://github.com/owncloud/ocis-thumbnails/issues/18)
  * Bugfix - Fix execution when passing program flags: [#15](https://github.com/owncloud/ocis-thumbnails/issues/15)
  * Change - Initial release of basic version: [#1](https://github.com/owncloud/ocis-thumbnails/issues/1)
  * Change - Use predefined resolutions for thumbnail generation: [#7](https://github.com/owncloud/ocis-thumbnails/issues/7)
  * Change - Implement the first working version: [#3](https://github.com/owncloud/ocis-thumbnails/pull/3)

   https://github.com/owncloud/product/issues/244

* Enhancement - Add a command to list the versions of running instances: [#226](https://github.com/owncloud/product/issues/226)

   Tags: accounts

   Added a micro command to list the versions of running accounts services.

   https://github.com/owncloud/product/issues/226

* Enhancement - Add the webdav service: [#244](https://github.com/owncloud/product/issues/244)

   Tags: webdav

  * Enhancement - Add version command: [#226](https://github.com/owncloud/product/issues/226)
  * Bugfix - Build docker images with alpine:latest instead of alpine:edge: [#22](https://github.com/owncloud/ocis-webdav/pull/22)
  * Change Change status not found on missing thumbnail: [#20](https://github.com/owncloud/ocis-webdav/issues/20)
  * Change - Initial release of basic version: [#1](https://github.com/owncloud/ocis-webdav/issues/1)
  * Change - Update ocis-pkg to version 2.2.0: [#16](https://github.com/owncloud/ocis-webdav/issues/16)
  * Enhancement - Configuration: [#14](https://github.com/owncloud/ocis-webdav/pull/14)
  * Enhancement - Implement preview API: [#13](https://github.com/owncloud/ocis-webdav/pull/13)

   https://github.com/owncloud/product/issues/244

* Enhancement - Better adopt Go-Micro: [#840](https://github.com/owncloud/ocis/pull/840)

   Tags: ocis

   There are a few building blocks that we were relying on default behavior, such as
   `micro.Registry` and the go-micro client. In order for oCIS to work in any environment and not
   relying in black magic configuration or running daemons we need to be able to:

   - Provide with a configurable go-micro registry. - Use our own go-micro client adjusted to our
   own needs (i.e: custom timeout, custom dial timeout, custom transport...)

   This PR is relying on 2 env variables from Micro: `MICRO_REGISTRY` and
   `MICRO_REGISTRY_ADDRESS`. The latter does not make sense to provide if the registry is not
   `etcd`.

   The current implementation only accounts for `mdns` and `etcd` registries, defaulting to
   `mdns` when not explicitly defined to use `etcd`.

   https://github.com/owncloud/ocis/pull/840

* Enhancement - Add permission check when assigning and removing roles: [#879](https://github.com/owncloud/ocis/issues/879)

   Everyone could add and remove roles from users. Added a new permission and a check so that only
   users with the role management permissions can assign and unassign roles.

   https://github.com/owncloud/ocis/issues/879

* Enhancement - Create OnlyOffice extension: [#857](https://github.com/owncloud/ocis/pull/857)

   Tags: OnlyOffice

   We've created an OnlyOffice extension which enables users to create and edit docx documents
   and open spreadsheets and presentations.

   https://github.com/owncloud/ocis/pull/857

* Enhancement - Show basic-auth warning only once: [#886](https://github.com/owncloud/ocis/pull/886)

   Show basic-auth warning only on startup instead on every request.

   https://github.com/owncloud/ocis/pull/886

* Enhancement - Add glauth fallback backend: [#649](https://github.com/owncloud/ocis/pull/649)

   We introduced the `fallback-datastore` config option and the corresponding options to allow
   configuring a simple chain of two handlers.

   Simple, because it is intended for bind and single result search queries. Merging large sets of
   results is currently out of scope. For now, the implementation will only search the fallback
   backend if the default backend returns an error or the number of results is 0. This is sufficient
   to allow an IdP to authenticate users from ocis as well as owncloud 10 as described in the [bridge
   scenario](https://owncloud.github.io/ocis/deployment/bridge/).

   https://github.com/owncloud/ocis-glauth/issues/18
   https://github.com/owncloud/ocis/pull/649

* Enhancement - Tidy dependencies: [#845](https://github.com/owncloud/ocis/pull/845)

   Methodology:

   ``` go-modules() { find . \( -name vendor -o -name '[._].*' -o -name node_modules \) -prune -o
   -name go.mod -print | sed 's:/go.mod$::' } ```

   ``` for m in $(go-modules); do (cd $m && go mod tidy); done ```

   https://github.com/owncloud/ocis/pull/845

* Enhancement - Launch a storage to store ocis-metadata: [#602](https://github.com/owncloud/ocis/pull/602)

   Tags: metadata, accounts, settings

   In the future accounts, settings etc. should be stored in a dedicated metadata storage. The
   services should talk to this storage directly, bypassing reva-gateway.

   https://github.com/owncloud/ocis/pull/602

* Enhancement - Add a version command to ocis: [#915](https://github.com/owncloud/ocis/pull/915)

   The version command was only implemented in the extensions. This adds the version command to
   ocis to list all services in the ocis namespace.

   https://github.com/owncloud/ocis/pull/915

* Enhancement - Create a proxy access-log: [#889](https://github.com/owncloud/ocis/pull/889)

   Logs client access at the proxy

   https://github.com/owncloud/ocis/pull/889

* Enhancement - Cache userinfo in proxy: [#877](https://github.com/owncloud/ocis/pull/877)

   Tags: proxy

   We introduced caching for the userinfo response. The token expiration is used for cache
   invalidation if available. Otherwise we fall back to a preconfigured TTL (default 10
   seconds).

   https://github.com/owncloud/ocis/pull/877

* Enhancement - Update reva to v1.4.1-0.20201125144025-57da0c27434c: [#1320](https://github.com/cs3org/reva/pull/1320)

   Mostly to bring fixes to pressing changes.

   https://github.com/cs3org/reva/pull/1320
   https://github.com/cs3org/reva/pull/1338

* Enhancement - Runtime Cleanup: [#1066](https://github.com/owncloud/ocis/pull/1066)

   Small runtime cleanup prior to Tech Preview release

   https://github.com/owncloud/ocis/pull/1066

* Enhancement - Update OCIS Runtime: [#1108](https://github.com/owncloud/ocis/pull/1108)

   - enhances the overall behavior of our runtime - runtime `db` file configurable - two new env
   variables to deal with the runtime - `RUNTIME_DB_FILE` and `RUNTIME_KEEP_ALIVE` -
   `RUNTIME_KEEP_ALIVE` defaults to `false` to provide backwards compatibility - if
   `RUNTIME_KEEP_ALIVE` is set to `true`, if a supervised process terminates the runtime will
   attempt to start with the same environment provided.

   https://github.com/owncloud/ocis/pull/1108

* Enhancement - Simplify tracing config: [#92](https://github.com/owncloud/product/issues/92)

   We now apply the oCIS tracing config to all services which have tracing. With this it is possible
   to set one tracing config for all services at the same time.

   https://github.com/owncloud/product/issues/92
   https://github.com/owncloud/ocis/pull/329
   https://github.com/owncloud/ocis/pull/409

* Enhancement - Update glauth to dev fd3ac7e4bbdc93578655d9a08d8e23f105aaa5b2: [#834](https://github.com/owncloud/ocis/pull/834)

   We updated glauth to dev commit fd3ac7e4bbdc93578655d9a08d8e23f105aaa5b2, which allows to
   skip certificate checks for the owncloud backend.

   https://github.com/owncloud/ocis/pull/834

* Enhancement - Update glauth to dev 4f029234b2308: [#786](https://github.com/owncloud/ocis/pull/786)

   Includes a bugfix, don't mix graph and provisioning api.

   https://github.com/owncloud/ocis/pull/786

* Enhancement - Update konnectd to v0.33.8: [#744](https://github.com/owncloud/ocis/pull/744)

   This update adds options which allow the configuration of oidc-token expiration parameters:
   KONNECTD_ACCESS_TOKEN_EXPIRATION, KONNECTD_ID_TOKEN_EXPIRATION and
   KONNECTD_REFRESH_TOKEN_EXPIRATION.

   Other changes from upstream:

   - Generate random endsession state for external authority - Update dependencies in
   Dockerfile - Set prompt=None to avoid loops with external authority - Update Jenkins
   reporting plugin from checkstyle to recordIssues - Remove extra kty key from JWKS top level
   document - Fix regression which encodes URL fragments twice - Avoid generating
   fragment/query URLs with wrong order - Return state for oidc endsession response redirects -
   Use server provided username to avoid case mismatch - Use signed-out-uri if set as fallback for
   goodbye redirect on saml slo - Add checks to ensure post_logout_redirect_uri is not empty - Fix
   SAML2 logout request parsing - Cure panic when no state is found in saml esr - Use SAML IdP Issuer
   value from meta data entityID - Allow configuration of expiration of oidc access, id and
   refresh tokens - Implement trampolin for external OIDC authority end session - Update
   ca-certificates version

   https://github.com/owncloud/ocis/pull/744

* Enhancement - Update reva to v1.4.1-0.20201123062044-b2c4af4e897d: [#823](https://github.com/owncloud/ocis/pull/823)

  * Refactor the uploading files workflow from various clients [cs3org/reva#1285](https://github.com/cs3org/reva/pull/1285), [cs3org/reva#1314](https://github.com/cs3org/reva/pull/1314)
  * [OCS] filter share with me requests [cs3org/reva#1302](https://github.com/cs3org/reva/pull/1302)
  * Fix listing shares for nonexistent path [cs3org/reva#1316](https://github.com/cs3org/reva/pull/1316)
  * prevent nil pointer when listing shares [cs3org/reva#1317](https://github.com/cs3org/reva/pull/1317)
  * Sharee retrieves the information about a share -but gets response containing all the shares [owncloud/ocis-reva#260](https://github.com/owncloud/ocis-reva/issues/260)
  * Deleting a public link after renaming a file [owncloud/ocis-reva#311](https://github.com/owncloud/ocis-reva/issues/311)
  * Avoid log spam [cs3org/reva#1323](https://github.com/cs3org/reva/pull/1323), [cs3org/reva#1324](https://github.com/cs3org/reva/pull/1324)
  * Fix trashbin [cs3org/reva#1326](https://github.com/cs3org/reva/pull/1326)

   https://github.com/owncloud/ocis-reva/issues/260
   https://github.com/owncloud/ocis-reva/issues/311
   https://github.com/owncloud/ocis/pull/823
   https://github.com/cs3org/reva/pull/1285
   https://github.com/cs3org/reva/pull/1302
   https://github.com/cs3org/reva/pull/1314
   https://github.com/cs3org/reva/pull/1316
   https://github.com/cs3org/reva/pull/1317
   https://github.com/cs3org/reva/pull/1323
   https://github.com/cs3org/reva/pull/1324
   https://github.com/cs3org/reva/pull/1326

* Enhancement - Update reva to v1.4.1-0.20201130061320-ac85e68e0600: [#980](https://github.com/owncloud/ocis/pull/980)

  * Fix move operation in ocis storage driver [csorg/reva#1343](https://github.com/cs3org/reva/pull/1343)

   https://github.com/owncloud/ocis/issues/975
   https://github.com/owncloud/ocis/pull/980
   https://github.com/cs3org/reva/pull/1343

* Enhancement - Update reva to cdb3d6688da5: [#748](https://github.com/owncloud/ocis/pull/748)

  * let the gateway filter invalid references

   https://github.com/owncloud/ocis/pull/748
   https://github.com/cs3org/reva/pull/1274

* Enhancement - Update reva to dd3a8c0f38: [#725](https://github.com/owncloud/ocis/pull/725)

  * fixes etag propagation in the ocis driver

   https://github.com/owncloud/ocis/pull/725
   https://github.com/cs3org/reva/pull/1264

* Enhancement - Update reva to v1.4.1-0.20201127111856-e6a6212c1b7b: [#971](https://github.com/owncloud/ocis/pull/971)

   Tags: reva

  * Fix capabilities response for multiple client versions #1331 [cs3org/reva#1331](https://github.com/cs3org/reva/pull/1331)
  * Fix home storage redirect for remote.php/dav/files [cs3org/reva#1342](https://github.com/cs3org/reva/pull/1342)

   https://github.com/owncloud/ocis/pull/971
   https://github.com/cs3org/reva/pull/1331
   https://github.com/cs3org/reva/pull/1342

* Enhancement - Update reva to 063b3db9162b: [#1091](https://github.com/owncloud/ocis/pull/1091)

   - bring public link removal changes to OCIS. - fix subcommand name collision from renaming
   phoenix -> web.

   https://github.com/owncloud/ocis/issues/1098
   https://github.com/owncloud/ocis/pull/1091

* Enhancement - Add www-authenticate based on user agent: [#1009](https://github.com/owncloud/ocis/pull/1009)

   Tags: reva, proxy

   We now comply with HTTP spec by adding Www-Authenticate headers on every `401` request.
   Furthermore, we not only take care of such a thing at the Proxy but also Reva will take care of it.
   In addition, we now are able to lock-in a set of User-Agent to specific challenges.

   Admins can use this feature by configuring oCIS + Reva following this approach:

   ``` STORAGE_FRONTEND_MIDDLEWARE_AUTH_CREDENTIALS_BY_USER_AGENT="mirall:basic,
   Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:83.0) Gecko/20100101
   Firefox/83.0:bearer" \
   PROXY_MIDDLEWARE_AUTH_CREDENTIALS_BY_USER_AGENT="mirall:basic, Mozilla/5.0
   (Macintosh; Intel Mac OS X 10.15; rv:83.0) Gecko/20100101 Firefox/83.0:bearer" \
   PROXY_ENABLE_BASIC_AUTH=true \ go run cmd/ocis/main.go server ```

   We introduced two new environment variables:

   `STORAGE_FRONTEND_MIDDLEWARE_AUTH_CREDENTIALS_BY_USER_AGENT` as well as
   `PROXY_MIDDLEWARE_AUTH_CREDENTIALS_BY_USER_AGENT`, The reason they have the same value
   is not to rely on the os env on a distributed environment, so in redundancy we trust. They both
   configure the same on the backend storage and oCIS Proxy.

   https://github.com/owncloud/ocis/pull/1009
