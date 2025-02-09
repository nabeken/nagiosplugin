#!/bin/sh

fail() {
  local msg=$1
  echo "NOT OK: $msg" >&2
  exit 1
}

if [ ! -f "./_testmain" ]; then
  fail "_testmain does not exist"
fi

# UNKNOWN
RET_MSG=$(./_testmain)
RET_CODE=$?

if [ "${RET_MSG}" != "TEST CHECK UNKNOWN: no check result specified" ]; then
  fail "unexpected output; got '${RET_MSG}'"
fi

# CRITICAL
RET_MSG=$(./_testmain -critical)
RET_CODE=$?

if [ "${RET_MSG}" != "TEST CHECK CRITICAL: CRITICAL" ]; then
  fail "unexpected output; got '${RET_MSG}'"
fi

# WARNING
RET_MSG=$(./_testmain -warning)
RET_CODE=$?

if [ "${RET_MSG}" != "TEST CHECK WARNING: WARNING" ]; then
  fail "unexpected output; got '${RET_MSG}'"
fi

# OK
RET_MSG=$(./_testmain -ok)
RET_CODE=$?

if [ "${RET_MSG}" != "TEST CHECK OK: OK" ]; then
  fail "unexpected output; got '${RET_MSG}'"
fi

echo "All tests passed"
