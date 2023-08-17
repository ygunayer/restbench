#include "transcript.h"

void api::v1::transcript::TranscriptController::getV1(
  const HttpRequestPtr &req,
  std::function<void(const HttpResponsePtr &)> &&callback,
  const std::string &id) {
  auto response = HttpResponse::newHttpResponse();
  response->setBody("Hello, World!");
  callback(response);
}
