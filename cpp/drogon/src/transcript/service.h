#ifndef __TRANSCRIPT__SERVICE_H__
#define __TRANSCRIPT__SERVICE_H__

#include <drogon/orm/DbClient.h>

namespace transcript {

class TranscriptService {
public:
  TranscriptService(drogon::orm::DbClientPtr db) : db(db) {}
private:
  drogon::orm::DbClientPtr db;
};

};

#endif
