/*

	This package implements decision trees.

	ID3DecisionTree:
		Builds a decision tree using the ID3 algorithm
			by picking the Attribute which maximises
			Information Gain at each node.

		Attributes must be CategoricalAttributes at
			present, so discretise beforehand (see
			filters)

  CART (Classification and Regression Trees):
    Builds a binary decision tree using the CART algorithm
      using a greedy approach to find the best split at each node.

    Can be used for regression and classficiation.
      Attributes have to be FloatAttributes even for classification.
      Hence, convert to Integer Labels before hand for Classficiation.

	RandomTree:
		Builds a decision tree using the ID3 algorithm
			by picking the Attribute amongst those
			randomly selected that maximises Information
			Gain

		Attributes must be CategoricalAttributes at
			present, so discretise beforehand (see
			filters)

*/

package trees
