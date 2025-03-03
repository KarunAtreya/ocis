package proxy

import future.keywords.if
import data.utils

default granted := true

granted = false if {
    utils.is_request_type_put
    not input.request.path == "/data"
    not utils.collection_contains(utils.ALLOWED_FILE_EXTENSIONS, input.request.path)
}

granted = false if {
    utils.is_request_type_post
    startswith(input.request.path, "/remote.php")
    not utils.collection_contains(utils.ALLOWED_FILE_EXTENSIONS, input.resource.name)
}
