# Build clockwerk simple dashboard
FROM node:13.12.0-alpine

# Change working directory
WORKDIR /app

# Cerate env PATH
ENV PATH /app/node_modules/.bin:$PATH

# Copy package.json to working dir
COPY package.json ./

# Copy package-lock.json to working dir
COPY package-lock.json ./

# Install prerequisite for reactjs app
RUN npm install --silent

# Install react-scripts for reactjs app
RUN npm install react-scripts@3.4.1 -g --silent

# Copying all source
COPY . ./

# Start the dashboard
ENTRYPOINT ["npm", "start"]