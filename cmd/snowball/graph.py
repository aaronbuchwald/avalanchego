import pandas as pd
import matplotlib.pyplot as plt

# Read the CSV file into a pandas DataFrame
data = pd.read_csv('sim.csv')

# Create a scatter plot
plt.scatter(data['byz'], data['rounds'], alpha=0.5)

# Add colorbar to accentuate data concentration
plt.colorbar()

# Set labels and title
plt.xlabel('byz')
plt.ylabel('rounds')
plt.title('Scatter Plot')

# Show the plot
plt.show()
