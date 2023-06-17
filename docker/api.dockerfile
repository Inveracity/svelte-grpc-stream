from python:3.11-alpine

RUN mkdir -p /app/api
COPY api /app/api
WORKDIR /app
RUN pip install -r api/requirements.txt

CMD ["uvicorn", "--host=0.0.0.0", "--port=8000", "api.main:app", "--reload"]