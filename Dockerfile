FROM openjdk:8-jdk-alpine
RUN addgroup -S appuser && adduser -S appuser -G appuser
USER appuser:appuser
