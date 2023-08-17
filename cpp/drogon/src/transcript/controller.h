#ifndef __TRANSCRIPT__CONTROLLER_H__
#define __TRANSCRIPT__CONTROLLER_H__

#include <drogon/HttpController.h>
#include "service.h"

using namespace drogon;
using namespace transcript;

namespace api {
namespace v1 {
namespace transcript {

class TranscriptController : public HttpController<TranscriptController> {
public:
  TranscriptController(const std::shared_ptr<TranscriptService> service) : service(service) {}

  METHOD_LIST_BEGIN;
    ADD_METHOD_TO(TranscriptController::getV1, "/transcript/{id}", Get);
  METHOD_LIST_END;

  void getV1(const HttpRequestPtr &req, std::function<void(const HttpResponsePtr &)> &&callback, const std::string &id);
private:
  const std::shared_ptr<TranscriptService> service;
};

}
}
}

#endif
