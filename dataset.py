import numpy as np
from torch.utils.data import Dataset

class Dataset(Dataset):
    def __init__(self, keys: np.ndarray, values: np.ndarray, gridNx):


        super(Dataset, self).__init__()

    def __len__(self):
        pass

    def __getitem__(self, item):
        pass

    def __add__(self, other):
        pass