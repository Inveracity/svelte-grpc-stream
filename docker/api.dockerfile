from python:3.11-alpine

RUN mkdir /app
WORKDIR /app
COPY api .
RUN pip install -r requirements.txt

CMD ["uvicorn", "--host=0.0.0.0", "--port=8000", "main:app", "--reload"]